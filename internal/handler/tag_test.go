package handler

import (
	"article-tags/internal/model"
	mocks "article-tags/mocks/internal_/model"
	svctypes "article-tags/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Resp struct {
	Status  int                       `json:"status,omitempty"`
	Message string                    `json:"message,omitempty"`
	Data    *svctypes.GetUserTagsResp `json:"data,omitempty"`
}

func TestApplication_AddTag(t *testing.T) {

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		mockDB     func(t *testing.T) model.UserTagStore
		args       args
		body       []byte
		wantStatus int
		wantResp   Resp
	}{
		{
			name: "Success",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Save(mock.Anything, mock.Anything).Return(nil)
				db.EXPECT().GetByPublicationTag(mock.Anything, mock.Anything).Return(nil, nil)
				db.EXPECT().DescribeTable(mock.Anything).Return(errors.New("table not exists"))
				db.EXPECT().CreateTable(mock.Anything).Return(nil)

				return db
			},
			body: []byte(`{
               "username": "Sandip",
               "publication" : "this is a content",
               "tags": ["tech","science"]
              }`),
			wantStatus: http.StatusCreated,
			wantResp: Resp{
				Status:  http.StatusCreated,
				Message: "tags added successfully",
				Data:    nil,
			},
		},
		{
			name: "Should fail when create table fail",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().DescribeTable(mock.Anything).Return(errors.New("table not exists"))
				db.EXPECT().CreateTable(mock.Anything).Return(errors.New("create table failed"))

				return db
			},
			body: []byte(`{
               "username": "Sandip",
               "publication" : "this is a content",
               "tags": ["tech","science"]
              }`),
			wantStatus: http.StatusInternalServerError,
			wantResp: Resp{
				Status:  http.StatusInternalServerError,
				Message: "db operation failed",
				Data:    nil,
			},
		},
		{
			name: "Should fail when validation failed",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				return db
			},
			body: []byte(`{
               "username": "",
               "publication" : "this is a content",
               "tags": ["tech","science"]
              }`),
			wantStatus: http.StatusBadRequest,
			wantResp: Resp{
				Status:  http.StatusBadRequest,
				Message: "bad request",
				Data:    nil,
			},
		},
		{
			name: "should return error when save failed",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Save(mock.Anything, mock.Anything).Return(errors.New("dynamo:error"))
				db.EXPECT().GetByPublicationTag(mock.Anything, mock.Anything).Return(nil, nil)
				db.EXPECT().DescribeTable(mock.Anything).Return(nil)

				return db
			},
			body: []byte(`{
               "username": "Sandip",
               "publication" : "ST",
               "tags": ["tech","science"]
              }`),
			wantStatus: http.StatusInternalServerError,
			wantResp: Resp{
				Status:  http.StatusInternalServerError,
				Message: "failed to save tag",
				Data:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				ArticleStore: tt.mockDB(t),
			}

			r := gin.Default()
			r.POST("/tags", app.AddTag)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/tags", bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())

			var actualResp Resp
			err = json.Unmarshal(w.Body.Bytes(), &actualResp)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			assert.Equal(t, tt.wantResp.Message, actualResp.Message)
			assert.Equal(t, tt.wantResp.Status, actualResp.Status)
		})
	}
}

func TestApplication_GetPopularTags(t *testing.T) {
	type args struct {
		username    string
		publication string
	}

	tests := []struct {
		name       string
		mockDB     func(t *testing.T) model.UserTagStore
		args       args
		body       []byte
		wantStatus int
		wantResp   Resp
	}{
		{
			name: "Success",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().GetPopularTags(mock.Anything, mock.Anything, mock.Anything).Return(
					[]*model.UserTag{{
						PK:          "Sandip#ST",
						SK:          "tech",
						Publication: "ST",
					}}, nil)

				return db
			},
			args: args{
				username:    "Sandip",
				publication: "ST",
			},
			wantStatus: http.StatusOK,
			wantResp: Resp{
				Status:  http.StatusOK,
				Message: "success",
				Data: &svctypes.GetUserTagsResp{
					Publication: "ST",
					Tags:        []string{"tech"},
				},
			},
		},
		{
			name: "Should fail when get call fail",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().GetPopularTags(mock.Anything, mock.Anything, mock.Anything).Return(
					nil, errors.New("dynamo:error"))

				return db
			},
			args: args{
				username:    "Sandip",
				publication: "ST",
			},
			wantStatus: http.StatusInternalServerError,
			wantResp: Resp{
				Status:  http.StatusInternalServerError,
				Message: "db operation failed",
				Data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				ArticleStore: tt.mockDB(t),
			}

			r := gin.Default()
			r.GET("/tags/:publication/popular", app.GetPopularTags)

			w := httptest.NewRecorder()
			url := fmt.Sprintf("/tags/%s/popular?username=%s", tt.args.publication, tt.args.username)
			req, err := http.NewRequest("GET", url, bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())

			var actualResp Resp
			err = json.Unmarshal(w.Body.Bytes(), &actualResp)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			assert.Equal(t, tt.wantResp.Message, actualResp.Message)
			assert.Equal(t, tt.wantResp.Status, actualResp.Status)
			assert.Equal(t, tt.wantResp.Data, actualResp.Data)
		})
	}
}

func TestApplication_GetFollowedTags(t *testing.T) {
	type args struct {
		username    string
		publication string
	}

	tests := []struct {
		name       string
		mockDB     func(t *testing.T) model.UserTagStore
		args       args
		body       []byte
		wantStatus int
		wantResp   Resp
	}{
		{
			name: "Success",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Get(mock.Anything, mock.Anything, mock.Anything).Return(
					[]*model.UserTag{{
						PK:          "Sandip#ST",
						SK:          "tech",
						Publication: "ST",
					}}, nil)

				return db
			},
			args: args{
				username:    "Sandip",
				publication: "ST",
			},
			wantStatus: http.StatusOK,
			wantResp: Resp{
				Status:  http.StatusOK,
				Message: "success",
				Data: &svctypes.GetUserTagsResp{
					Publication: "ST",
					Tags:        []string{"tech"},
				},
			},
		},
		{
			name: "Should fail when get call fail",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Get(mock.Anything, mock.Anything, mock.Anything).Return(
					nil, errors.New("dynamo:error"))

				return db
			},
			args: args{
				username:    "Sandip",
				publication: "ST",
			},
			wantStatus: http.StatusInternalServerError,
			wantResp: Resp{
				Status:  http.StatusInternalServerError,
				Message: "db operation failed",
				Data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				ArticleStore: tt.mockDB(t),
			}

			r := gin.Default()
			r.GET("/tags/:publication", app.GetFollowedTags)

			w := httptest.NewRecorder()
			url := fmt.Sprintf("/tags/%s?username=%s", tt.args.publication, tt.args.username)
			req, err := http.NewRequest("GET", url, bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())

			var actualResp Resp
			err = json.Unmarshal(w.Body.Bytes(), &actualResp)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			assert.Equal(t, tt.wantResp.Message, actualResp.Message)
			assert.Equal(t, tt.wantResp.Status, actualResp.Status)
			assert.Equal(t, tt.wantResp.Data, actualResp.Data)
		})
	}
}

func TestApplication_DeleteTag(t *testing.T) {

	type args struct {
		publication string
	}
	tests := []struct {
		name       string
		mockDB     func(t *testing.T) model.UserTagStore
		args       args
		body       []byte
		wantStatus int
		wantResp   Resp
	}{
		{
			name: "Success",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Delete(mock.Anything, mock.Anything).Return(nil)
				db.EXPECT().GetByPublicationTag(mock.Anything, mock.Anything).Return(&model.UserTag{
					PK:          "Sandip#ST",
					SK:          "tech",
					Publication: "ST",
				}, nil)

				return db
			},
			args: args{publication: "ST"},
			body: []byte(`{
               "username": "Sandip",
               "tag": "tech"
              }`),
			wantStatus: http.StatusOK,
			wantResp: Resp{
				Status:  http.StatusOK,
				Message: "tag deleted successfully",
				Data:    nil,
			},
		},
		{
			name: "Should fail when validation failed",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				return db
			},
			args: args{
				publication: "ST",
			},
			body: []byte(`{
               "username": "",
               "tag": ""
              }`),
			wantStatus: http.StatusBadRequest,
			wantResp: Resp{
				Status:  http.StatusBadRequest,
				Message: "bad request",
				Data:    nil,
			},
		},
		{
			name: "should return error when delete fail",
			mockDB: func(t *testing.T) model.UserTagStore {
				db := mocks.NewUserTagStore(t)
				db.EXPECT().Delete(mock.Anything, mock.Anything).Return(errors.New("dynamo:error"))
				db.EXPECT().GetByPublicationTag(mock.Anything, mock.Anything).Return(&model.UserTag{
					PK:          "Sandip#ST",
					SK:          "tech",
					Publication: "ST",
				}, nil)

				return db
			},
			args: args{publication: "ST"},
			body: []byte(`{
               "username": "Sandip",
               "tag": "tech"
              }`),
			wantStatus: http.StatusInternalServerError,
			wantResp: Resp{
				Status:  http.StatusInternalServerError,
				Message: "failed to delete tag",
				Data:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				ArticleStore: tt.mockDB(t),
			}

			r := gin.Default()
			r.DELETE("/tags/:publication", app.DeleteTag)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", "/tags/"+tt.args.publication, bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())

			var actualResp Resp
			err = json.Unmarshal(w.Body.Bytes(), &actualResp)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			assert.Equal(t, tt.wantResp.Message, actualResp.Message)
			assert.Equal(t, tt.wantResp.Status, actualResp.Status)
		})
	}
}
