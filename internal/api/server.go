package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/config"
	"github.com/tedawf/bulbsocial/internal/db"
	"github.com/tedawf/bulbsocial/internal/service"
	"go.uber.org/zap"
)

// Server serves HTTP requests for our backend service
type Server struct {
	router     *chi.Mux
	logger     *zap.SugaredLogger
	config     config.Config
	tokenMaker auth.TokenMaker

	userService *service.UserService
	postService *service.PostService
}

// setupRoutes initializes all HTTP routes
func (s *Server) setupRoutes() {
	s.router.Get("/health", s.handleHealthCheck)

	s.router.Route("/users", func(r chi.Router) {
		r.Post("/", s.handleCreateUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", s.handleGetUser)
		})

		r.Post("/login", s.handleLoginUser)
	})

	s.router.Route("/posts", func(r chi.Router) {
		r.Post("/", s.handleCreatePost)

		r.Route("/{postID}", func(r chi.Router) {
			r.Get("/", s.handleGetPost)
		})
	})

	s.router.Get("/feed", s.handleGetFeed)
}

// setupMiddlewares adds middlewares to the router
func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Use(s.tokenAuthMiddleware())
}

// NewServer creates a new HTTP server with routes and dependencies
func NewServer(store db.Store, logger *zap.SugaredLogger, config config.Config) (*Server, error) {
	tokenMaker, err := auth.NewJWTMaker(config.AuthTokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		router:     chi.NewRouter(),
		logger:     logger,
		config:     config,
		tokenMaker: tokenMaker,

		userService: service.NewUserService(store, tokenMaker),
		postService: service.NewPostService(store),
	}

	server.setupMiddlewares()
	server.setupRoutes()

	return server, nil
}

// Start starts the HTTP server on a specified address
func (s *Server) Start(address string) error {
	return http.ListenAndServe(address, s.router)
}
