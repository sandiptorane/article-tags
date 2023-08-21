package model

import "context"

type UserTagStore interface {
	Save(ctx context.Context, data *UserTag) error
}
