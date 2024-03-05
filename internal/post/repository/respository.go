package repository

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"gorm.io/gen"
)

type Repository interface {
	Post() PostRepository
	Category() CategoryRepository
	Tag() TagRepository
}

type PostRepository interface {
	Persist(c context.Context, row *model.Post) (err error)
	Get()
	GetList()
	Delete()
}

type CategoryRepository interface {
	Persist(c context.Context, p *model.Post) error
	Get(c context.Context, cond ...gen.Condition) (*model.Category, error)
	GetList()
	Delete()
}

type TagRepository interface {
	Persist(context.Context, *model.Post) error
	Get(c context.Context, cond ...gen.Condition) (*model.Tag, error)
	GetList()
	Delete()
}
