package store

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type Store interface {
	Post() PostStore
	Tag() TagStore
	Category() CategoryStore
}

type PostStore interface {
	Create(c context.Context, post *model.Post) error
	Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, post *model.Post) error
	Query(c context.Context, conds ...gen.Condition) (*model.Post, error)
}

type CategoryStore interface {
	Create(c context.Context, category *model.Category) error
	Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, category *model.Category) error
	Query(c context.Context, conds ...gen.Condition) (*model.Category, error)
}

type TagStore interface {
	Create(c context.Context, tag *model.Tag) error
	Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, post *model.Tag) error
	Query(c context.Context, conds ...gen.Condition) (*model.Tag, error)
}
