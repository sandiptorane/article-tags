package handler

import (
	"article-tags/database/dynamo"
	"article-tags/database/model"

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
