package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/config"
	"github.com/tedawf/bulbsocial/internal/db"
	"github.com/tedawf/bulbsocial/internal/util"
	"go.uber.org/zap"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		AuthTokenKey:        util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(store, zap.NewNop().Sugar(), config)
	require.NoError(t, err)

	return server
}
