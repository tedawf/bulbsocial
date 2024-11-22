package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tedawf/bulb-core/internal/ratelimiter"
)

func TestRateLimiterMiddleware(t *testing.T) {
	cfg := config{
		ratelimiter: ratelimiter.Config{
			RequestsPerTimeFrame: 20,
			TimeFrame:            time.Second * 5,
			Enabled:              true,
		},
		addr: ":8080",
	}

	app := newTestApplication(t, cfg)
	ts := httptest.NewServer(app.mount())
	defer ts.Close()

	client := &http.Client{}
	mockIP := "192.168.1.1"
	marginOfError := 2

	for i := 0; i < cfg.ratelimiter.RequestsPerTimeFrame+marginOfError; i++ {
		req, err := http.NewRequest("GET", ts.URL+"/v1/health", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		req.Header.Set("X-Forwarded-For", mockIP)

		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		defer res.Body.Close()

		if i < cfg.ratelimiter.RequestsPerTimeFrame {
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got %v", res.Status)
			}
		} else {
			if res.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status Too Many Requests, got %v", res.Status)
			}
		}
	}
}
