package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tedawf/tradebulb/internal/auth"
	"github.com/tedawf/tradebulb/internal/mail"
	"github.com/tedawf/tradebulb/internal/store"
	"github.com/tedawf/tradebulb/internal/store/cache"
	"go.uber.org/zap"
)

type config struct {
	addr        string
	db          dbConfig
	env         string
	mail        mailConfig
	frontendURL string
	auth        authConfig
	redisCfg    redisConfig
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type mailConfig struct {
	sendGrid  sendGridConfig
	fromEmail string
	exp       time.Duration
}

type sendGridConfig struct {
	apiKey string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	mailer        mail.Client
	authenticator auth.Authenticator
	cacheStorage  cache.Storage
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.With(app.BasicAuthMiddleware()).Get("/health", app.healthCheck)

		r.Route("/posts", func(r chi.Router) {
			r.Use(app.TokenAuthMiddleware())
			r.Post("/", app.createPost)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postContextMiddleware)

				r.Get("/", app.getPost)
				r.Patch("/", app.checkPostOwnership("moderator", app.updatePost))
				r.Delete("/", app.checkPostOwnership("admin", app.deletePost))
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Put("/verify/{token}", app.verifyUser)

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.TokenAuthMiddleware())

				r.Get("/", app.getUser)
				r.Put("/follow", app.followUser)
				r.Put("/unfollow", app.unfollowUser)
			})

			r.Group(func(r chi.Router) {
				r.Use(app.TokenAuthMiddleware())
				r.Get("/feed", app.getUserFeed)
			})
		})

		// public
		r.Route("/auth", func(r chi.Router) {
			r.Post("/user", app.registerUser)
			r.Post("/token", app.createToken)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// graceful shutdown (after 5s)
	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
