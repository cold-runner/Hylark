package repository

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"gorm.io/gen"
)

type Repository interface {
	Lark() LarkRepository
	Social() SocialRepository
}

type LarkRepository interface {
	Persist(c context.Context, lark *model.Lark) error
	Update(c context.Context, lark *model.Lark) error
	IsExist(c context.Context, conds ...gen.Condition) (bool, error)
}

type SocialRepository interface {
	Persist(c context.Context, row *model.UserInteraction) error
}

type SocialRepository interface {
	Persist(c context.Context, row *model.UserInteraction) error
}
