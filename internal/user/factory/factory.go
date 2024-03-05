package factory

import (
	"context"
	"github.com/cold-runner/Hylark/internal/user/entity"
	"github.com/cold-runner/Hylark/internal/user/store"
	"gorm.io/gen"
)

type Factory interface {
	Lark() LarkFactory
	GetStore() store.Store
}

type LarkFactory interface {
	Produce(c context.Context, cond ...gen.Condition) (*entity.Lark, error)
}
