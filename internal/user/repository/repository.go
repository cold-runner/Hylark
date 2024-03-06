package repository

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
)

type Repository interface {
	Lark() LarkRepository
}

type LarkRepository interface {
	Persist(c context.Context, lark *model.Lark) error
	Update(c context.Context, lark *model.Lark) error
}
