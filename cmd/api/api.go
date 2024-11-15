package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tedawf/tradebulb/internal/auth"
	"github.com/tedawf/tradebulb/internal/mail"
	"github.com/tedawf/tradebulb/internal/store"
	"go.uber.org/zap"
)

type config struct {
	addr        string
	db          dbConfig
	env         string
	mail        mailConfig
	frontendURL string
	auth        authConfig
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
			r.Post("/", app.createPost)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postContextMiddleware)

				r.Get("/", app.getPost)
				r.Delete("/", app.deletePost)
				r.Patch("/", app.updatePost)
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Put("/verify/{token}", app.verifyUser)

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)

				r.Get("/", app.getUser)
				r.Put("/follow", app.followUser)
				r.Put("/unfollow", app.unfollowUser)
			})

			r.Group(func(r chi.Router) {
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

	app.logger.Infow("Server has started", "addr", app.config.addr, "env", app.config.env)

	return srv.ListenAndServe()
}
