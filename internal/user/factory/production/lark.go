package production

import (
	"context"
	"github.com/cold-runner/Hylark/internal/user/entity"
	"github.com/cold-runner/Hylark/internal/user/store"
	"gorm.io/gen"
)

type larkFactory struct {
	storeIns store.Store
}

func newLarkFactory(s store.Store) *larkFactory {
	return &larkFactory{storeIns: s}
}

func (l larkFactory) Produce(c context.Context, cond ...gen.Condition) (*entity.Lark, error) {
	panic("impl me")
}
