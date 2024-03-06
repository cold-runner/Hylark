package store

import (
	"context"

	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

type Store interface {
	Lark() LarkStore
	Social() SocialStore
}

type LarkStore interface {
	Create(c context.Context, user *model.Lark) error
	Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error
	Query(c context.Context, conds ...gen.Condition) (*model.Lark, error)
}

type SocialStore interface {
	Create(c context.Context, s *model.UserInteraction) error
	Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.UserInteraction) error
	Query(c context.Context, conds ...gen.Condition) (*model.UserInteraction, error)
}
