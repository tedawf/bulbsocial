package main

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tedawf/tradebulb/internal/auth"
	"github.com/tedawf/tradebulb/internal/db"
	"github.com/tedawf/tradebulb/internal/env"
	"github.com/tedawf/tradebulb/internal/mail"
	"github.com/tedawf/tradebulb/internal/store"
	"github.com/tedawf/tradebulb/internal/store/cache"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		frontendURL: env.GetString("FRONTEND_URL", "localhost:4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/tradebulb_local?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "dev"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "secret"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "bulbsocial",
			},
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("Database connection pool established")

	store := store.NewStorage(db)

	// cache
	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("redis cache connection pool established")

		defer rdb.Close()
	}

	cacheStorage := cache.NewRedisStorage(rdb)

	mailer := mail.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)

	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuthenticator,
		cacheStorage:  cacheStorage,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
