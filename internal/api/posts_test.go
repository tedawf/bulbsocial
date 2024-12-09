package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/db"
	mockdb "github.com/tedawf/bulbsocial/internal/db/mock"
	"github.com/tedawf/bulbsocial/internal/util"
	"go.uber.org/mock/gomock"
)

func TestCreatePostAPI(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(user.ID)

	testCases := []struct {
		name          string
		body          map[string]interface{}
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: map[string]interface{}{
				"user_id": user.ID,
				"title":   post.Title,
				"content": post.Content,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreatePostParams{
					UserID:  user.ID,
					Title:   post.Title,
					Content: post.Content,
				}
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchPost(t, recorder.Body, post)
			},
		},
		{
			name: "InternalError",
			body: map[string]interface{}{
				"user_id": user.ID,
				"title":   post.Title,
				"content": post.Content,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Post{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: map[string]interface{}{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NoSuchUser",
			body: map[string]interface{}{
				"user_id": user.ID,
				"title":   post.Title,
				"content": post.Content,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreatePostParams{
					UserID:  user.ID,
					Title:   post.Title,
					Content: post.Content,
				}
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Post{}, &pq.Error{Code: "23503"}) // foreign key violation
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/posts"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetPostAPI(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(user.ID)

	testCases := []struct {
		name          string
		postID        string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			postID: fmt.Sprintf("%d", post.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostByID(gomock.Any(), gomock.Eq(post.ID)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPost(t, recorder.Body, post)
			},
		},
		{
			name:   "InvalidID",
			postID: "abc",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostByID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "NotFound",
			postID: fmt.Sprintf("%d", post.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostByID(gomock.Any(), gomock.Eq(post.ID)).
					Times(1).
					Return(db.Post{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/posts/%s", tc.postID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomPost(userID int64) db.Post {
	return db.Post{
		ID:      util.RandomInt(1, 1000),
		UserID:  userID,
		Title:   util.RandomTitle(),
		Content: util.RandomContent(),
	}
}

func requireBodyMatchPost(t *testing.T, body *bytes.Buffer, post db.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPost APIResponse[db.Post]
	err = json.Unmarshal(data, &gotPost)

	require.NoError(t, err)
	require.Equal(t, post.ID, gotPost.Data.ID)
	require.Equal(t, post.UserID, gotPost.Data.UserID)
	require.Equal(t, post.Title, gotPost.Data.Title)
	require.Equal(t, post.Content, gotPost.Data.Content)
	require.Equal(t, post.CreatedAt, gotPost.Data.CreatedAt)
	require.Equal(t, post.UpdatedAt, gotPost.Data.UpdatedAt)
}
