package dynamo

import (
	"article-tags/database/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamodbAPI this interface used by application code. lists dynamodb functions
type DynamodbAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

type ArticleTag struct {
	db DynamodbAPI
}

// verify interface compliance in compile time
var _ DynamodbAPI = (*dynamodb.Client)(nil)

// GetInstance returns instance of articleStore
func GetInstance(db *dynamodb.Client) *ArticleTag {
	return &ArticleTag{db: db}
}

// Save user tag
func (a *ArticleTag) Save(ctx context.Context, data *model.UserTag) error {
	items, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	log.Println("input item", items)

	input := &dynamodb.PutItemInput{
		Item:      items,
		TableName: aws.String("article-tag"),
	}

	_, err = a.db.PutItem(ctx, input)

	counterInput := dynamodb.UpdateItemInput{
		TableName: aws.String("article-tag"),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("PUB#%s", data.Publication)},
			"SK": &types.AttributeValueMemberS{Value: data.SK},
		},
		UpdateExpression: aws.String("SET total_count = if_not_exists(total_count, :v1) + :incr, tag = :v2"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v1":   &types.AttributeValueMemberN{Value: "0"},
			":incr": &types.AttributeValueMemberN{Value: "1"},
			":v2":   &types.AttributeValueMemberS{Value: data.SK},
		},
	}

	_, err = a.db.UpdateItem(ctx, &counterInput)
	if err != nil {
		fmt.Println("error storing new item 2222: ", err)
		return err
	}
	return err
}

// Get UserTag
func (a *ArticleTag) Get(ctx context.Context, publication, username string) ([]*model.UserTag, error) {
	input := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK := :pub_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pub_id": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s#%s", username, publication)},
		},
		TableName: aws.String("article-tag"),
	}

	res, err := a.db.Query(ctx, input)
	if err != nil {
		log.Println("error fetching article", err)
		return nil, err
	}

	var articles []*model.UserTag

	for _, a := range res.Items {
		var article model.UserTag

		err = attributevalue.UnmarshalMap(a, &article)
		if err != nil {
			log.Println("unmarshall error", err)
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

// GetByPublicationTag fetch user followed tag for particular tag
func (a *ArticleTag) GetByPublicationTag(ctx context.Context, request model.UserTagRequest) (*model.UserTag, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s#%s", request.Username, request.Publication)},
			"SK": &types.AttributeValueMemberS{Value: request.Tag},
		},

		TableName: aws.String("article-tag"),
	}

	res, err := a.db.GetItem(ctx, input)
	if err != nil {
		log.Println("error fetching article", err)
		return nil, err
	}

	if res.Item == nil {
		return nil, nil
	}

	var article model.UserTag

	err = attributevalue.UnmarshalMap(res.Item, &article)
	if err != nil {
		log.Println("unmarshall error", err)
		return nil, err
	}

	return &article, nil
}

func (a *ArticleTag) GetPopularTags(ctx context.Context, username, publication string) ([]*model.UserTag, error) {
	// fetch already added tag for user
	existingTags, err := a.Get(ctx, publication, username)
	if err != nil {
		log.Println("error fetching users followed tags:", err)
		return nil, err
	}

	input := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK:= :pub_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pub_id": &types.AttributeValueMemberS{Value: fmt.Sprintf("PUB#%s", publication)},
		},
		TableName: aws.String("article-tag"),
	}

	if len(existingTags) > 0 {
		prepareFilterExpression(input, existingTags)
	}

	res, err := a.db.Query(ctx, input)
	if err != nil {
		log.Println("error fetching article", err)
		return nil, err
	}

	var articles []*model.UserTag

	for _, a := range res.Items {
		var article model.UserTag

		err = attributevalue.UnmarshalMap(a, &article)
		if err != nil {
			log.Println("unmarshall error", err)
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

func prepareFilterExpression(queryInput *dynamodb.QueryInput, existingTags []*model.UserTag) {
	filterExpression := "NOT (Tag IN ("

	var filterAttr []string

	for k, val := range existingTags {
		key := fmt.Sprintf(":exclude%v", k)

		filterAttr = append(filterAttr, key)

		queryInput.ExpressionAttributeValues[key] = &types.AttributeValueMemberS{Value: val.SK}
	}

	// join the filter expression placeholder
	filterExpression += strings.Join(filterAttr, ", ")

	// end of bracket
	filterExpression += "))"

	// update the filter expression
	queryInput.FilterExpression = aws.String(filterExpression)
}

func (a *ArticleTag) Delete(ctx context.Context, request *model.UserTagRequest) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s#%s", request.Username, request.Publication)},
			"SK": &types.AttributeValueMemberS{Value: request.Tag},
		},

		TableName: aws.String("article-tag"),
	}

	_, err := a.db.DeleteItem(ctx, input)

	counterInput := dynamodb.UpdateItemInput{
		TableName: aws.String("article-tag"),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("PUB#%s", request.Publication)},
			"SK": &types.AttributeValueMemberS{Value: request.Tag},
		},
		UpdateExpression: aws.String("SET total_count = if_not_exists(total_count, :v1) - :decr, tag = :v2"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v1":   &types.AttributeValueMemberN{Value: "0"},
			":decr": &types.AttributeValueMemberN{Value: "1"},
			":v2":   &types.AttributeValueMemberS{Value: request.Tag},
		},
	}

	_, err = a.db.UpdateItem(ctx, &counterInput)
	if err != nil {
		fmt.Println("error storing new item 2222: ", err)
		return err
	}

	return err
}
