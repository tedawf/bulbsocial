package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tedawf/bulb-core/internal/store/cache"
)

func TestGetUser(t *testing.T) {
	withRedis := config{
		redisCfg: redisConfig{
			enabled: true,
		},
	}

	app := newTestApplication(t, withRedis)
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		mockCacheStorage := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStorage.On("Get", int64(1)).Return(nil, nil).Twice()
		mockCacheStorage.On("Set", mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStorage.Calls = nil // reset
	})

	t.Run("should hit cache first and if not exists, set user in cache", func(t *testing.T) {
		mockCacheStorage := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStorage.On("Get", int64(42)).Return(nil, nil)
		mockCacheStorage.On("Get", int64(1)).Return(nil, nil)
		mockCacheStorage.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStorage.AssertNumberOfCalls(t, "Get", 2)

		mockCacheStorage.Calls = nil // reset
	})

	t.Run("should not hit cache if not enabled", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisConfig{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()

		mockCacheStorage := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStorage.AssertNotCalled(t, "Get")

		mockCacheStorage.Calls = nil // reset
	})
}
