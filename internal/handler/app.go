package handler

import (
	"article-tags/internal/database/dynamo"
	"article-tags/internal/model"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Application struct {
	ArticleStore model.UserTagStore
}

func NewApplication(db *dynamodb.Client) *Application {
	return &Application{
		ArticleStore: dynamo.GetInstance(db),
	}
}
