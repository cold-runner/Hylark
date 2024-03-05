package factory

import (
	"context"
	"github.com/cold-runner/Hylark/internal/user/entity"
	"gorm.io/gen"
)

type Factory interface {
	Lark() LarkFactory
}

type LarkFactory interface {
	Produce(c context.Context, cond ...gen.Condition) (*entity.Lark, error)
}
