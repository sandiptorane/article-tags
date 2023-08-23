package dynamo

import (
	"article-tags/internal/model"
	mocks "article-tags/mocks/internal_/database/dynamo"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestArticleTag_CreateTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
		mockDB  func(t *testing.T) DynamodbAPI
	}{
		{
			name:    "Success",
			wantErr: nil,
			mockDB: func(t *testing.T) DynamodbAPI {
				api := mocks.NewDynamodbAPI(t)
				api.EXPECT().CreateTable(mock.Anything, mock.Anything).Return(&dynamodb.CreateTableOutput{}, nil)

				return api
			},
		},
		{
			name:    "create table fail",
			wantErr: errors.New("dynamo:error"),
			mockDB: func(t *testing.T) DynamodbAPI {
				api := mocks.NewDynamodbAPI(t)
				api.EXPECT().CreateTable(mock.Anything, mock.Anything).Return(&dynamodb.CreateTableOutput{}, errors.New("dynamo:error"))

				return api
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleTag{
				db: tt.mockDB(t),
			}
			err := a.CreateTable(context.TODO())
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestArticleTag_Delete(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *model.UserTagRequest
	}
	tests := []struct {
		name    string
		wantErr error
		mockDB  func(t *testing.T) DynamodbAPI
		args    args
	}{
		{
			name:    "Success",
			wantErr: nil,
			args: args{
				ctx: context.TODO(),
				request: &model.UserTagRequest{
					Username:    "Sandip",
					Publication: "ST",
					Tag:         "tech",
				},
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)
				db.EXPECT().DeleteItem(mock.Anything, mock.Anything).Return(&dynamodb.DeleteItemOutput{
					Attributes: map[string]types.AttributeValue{
						"PK": &types.AttributeValueMemberS{Value: "test"},
					},
				}, nil)

				db.EXPECT().UpdateItem(mock.Anything, mock.Anything).Return(&dynamodb.UpdateItemOutput{}, nil)
				return db
			},
		},
		{
			name:    "Should fail when delete fail",
			wantErr: errors.New("dynamo:error"),
			args: args{
				ctx: context.TODO(),
				request: &model.UserTagRequest{
					Username:    "Sandip",
					Publication: "ST",
					Tag:         "tech",
				},
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)
				db.EXPECT().DeleteItem(mock.Anything, mock.Anything).Return(nil, errors.New("dynamo:error"))
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleTag{
				db: tt.mockDB(t),
			}
			err := a.Delete(tt.args.ctx, tt.args.request)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestArticleTag_DescribeTable(t *testing.T) {
	type fields struct {
		db DynamodbAPI
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockDB  func(t *testing.T) DynamodbAPI
		wantErr error
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				context.TODO(),
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)
				db.EXPECT().DescribeTable(mock.Anything, mock.Anything).Return(nil, nil)

				return db
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleTag{
				db: tt.mockDB(t),
			}
			err := a.DescribeTable(tt.args.ctx)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestArticleTag_GetPopularTags(t *testing.T) {
	type args struct {
		ctx         context.Context
		username    string
		publication string
	}
	tests := []struct {
		name    string
		wantErr error
		args    args
		mockDB  func(t *testing.T) DynamodbAPI
	}{
		{
			name:    "Success",
			wantErr: nil,
			args: args{
				ctx:         context.TODO(),
				username:    "test",
				publication: "test",
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)

				db.EXPECT().Query(mock.Anything, mock.Anything).Return(&dynamodb.QueryOutput{
					Items: func() []map[string]types.AttributeValue {
						var data []map[string]types.AttributeValue

						data = append(data, map[string]types.AttributeValue{
							"PK":          &types.AttributeValueMemberS{Value: "test"},
							"SK":          &types.AttributeValueMemberS{Value: "test"},
							"publication": &types.AttributeValueMemberS{Value: "ST"},
						})

						return data
					}(),
				}, nil)

				return db
			},
		},
		{
			name:    "Should fail when query fail",
			wantErr: errors.New("dynamo:error"),
			args: args{
				ctx:         context.TODO(),
				username:    "test",
				publication: "test",
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)

				db.EXPECT().Query(mock.Anything, mock.Anything).Return(nil, errors.New("dynamo:error"))

				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleTag{
				db: tt.mockDB(t),
			}
			_, err := a.GetPopularTags(tt.args.ctx, tt.args.username, tt.args.publication)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestArticleTag_Save(t *testing.T) {
	type args struct {
		ctx context.Context
		req *model.UserTagRequest
	}
	tests := []struct {
		name    string
		wantErr error
		args    args
		mockDB  func(t *testing.T) DynamodbAPI
	}{
		{
			name:    "Success",
			wantErr: nil,
			args: args{
				ctx: context.TODO(),
				req: &model.UserTagRequest{
					Username:    "Sandip",
					Publication: "ST",
					Tag:         "tech",
				},
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)
				db.EXPECT().PutItem(mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{
					Attributes: nil,
				}, nil)

				db.EXPECT().UpdateItem(mock.Anything, mock.Anything).Return(&dynamodb.UpdateItemOutput{}, nil)
				return db
			},
		},
		{
			name:    "Should fail when save fail",
			wantErr: errors.New("dynamo:error"),
			args: args{
				ctx: context.TODO(),
				req: &model.UserTagRequest{
					Username:    "Sandip",
					Publication: "ST",
					Tag:         "tech",
				},
			},
			mockDB: func(t *testing.T) DynamodbAPI {
				db := mocks.NewDynamodbAPI(t)
				db.EXPECT().PutItem(mock.Anything, mock.Anything).Return(nil, errors.New("dynamo:error"))

				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleTag{
				db: tt.mockDB(t),
			}
			err := a.Save(tt.args.ctx, tt.args.req)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}
