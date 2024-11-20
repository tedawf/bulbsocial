package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tedawf/tradebulb/internal/auth"
	"github.com/tedawf/tradebulb/internal/ratelimiter"
	"github.com/tedawf/tradebulb/internal/store"
	"github.com/tedawf/tradebulb/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.Must(zap.NewProduction()).Sugar()

	mockStore := store.NewMockStore()
	mockCacheStorage := cache.NewMockStore()

	testAuthenticator := &auth.TestAuthenticator{}

	// rate limiter
	ratelimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.ratelimiter.RequestsPerTimeFrame,
		cfg.ratelimiter.TimeFrame,
	)

	return &application{
		config:        cfg,
		logger:        logger,
		store:         mockStore,
		cacheStorage:  mockCacheStorage,
		authenticator: testAuthenticator,
		ratelimiter:   ratelimiter,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected the response code to be %d and we got %d", expected, actual)
	}
}
