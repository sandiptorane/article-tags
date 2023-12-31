package model

import "context"

type UserTagStore interface {
	CreateTable(ctx context.Context) error
	DescribeTable(ctx context.Context) error
	Save(ctx context.Context, data *UserTag) error
	Get(ctx context.Context, publication, username string) ([]*UserTag, error)
	GetByPublicationTag(ctx context.Context, request *UserTagRequest) (*UserTag, error)
	GetPopularTags(ctx context.Context, username, publication string) ([]*UserTag, error)
	Delete(ctx context.Context, request *UserTagRequest) error
}
