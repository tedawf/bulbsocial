package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/db"
	mockdb "github.com/tedawf/bulbsocial/internal/db/mock"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestGetUser(t *testing.T) {
	user := db.User{
		ID:                db.RandomInt(1, 1000),
		Email:             db.RandomUsername() + "@email.com",
		Username:          db.RandomUsername(),
		CreatedAt:         time.Now().Round(time.Second),
		PasswordChangedAt: sql.NullTime{},
	}

	testCases := []struct {
		name          string
		userID        int64
		buildStubs    func(ctrl *gomock.Controller, store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Happy",
			userID: user.ID,
			buildStubs: func(ctrl *gomock.Controller, store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn func(db.Querier) error) error {
						mq := mockdb.NewMockQuerier(ctrl)

						mq.EXPECT().
							GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
							Times(1).
							Return(user, nil)

						return fn(mq)
					})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var gotResponse APIResponse[UserResponse]
				err = json.Unmarshal(data, &gotResponse)
				require.NoError(t, err)
				require.Equal(t, NewUserResponse(user), gotResponse.Data)
			},
		},
		{
			name:   "NotFound",
			userID: user.ID,
			buildStubs: func(ctrl *gomock.Controller, store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn func(db.Querier) error) error {
						mq := mockdb.NewMockQuerier(ctrl)

						mq.EXPECT().
							GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
							Times(1).
							Return(db.User{}, sql.ErrNoRows)

						return fn(mq)
					})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			userID: user.ID,
			buildStubs: func(ctrl *gomock.Controller, store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn func(db.Querier) error) error {
						mq := mockdb.NewMockQuerier(ctrl)

						mq.EXPECT().
							GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
							Times(1).
							Return(db.User{}, sql.ErrConnDone)

						return fn(mq)
					})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "BadRequest",
			userID: -1,
			buildStubs: func(ctrl *gomock.Controller, store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(ctrl, store)

			// start test server and send request
			server := NewServer(store, zap.NewNop().Sugar())
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/%d", tc.userID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
