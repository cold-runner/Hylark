package production

import (
	"github.com/cold-runner/Hylark/internal/user/factory"
	"github.com/cold-runner/Hylark/internal/user/store"
)

type prodFactory struct {
	storeIns store.Store
}

func NewFactory(s store.Store) factory.Factory {
	return &prodFactory{storeIns: s}
}

func (f prodFactory) Lark() factory.LarkFactory {
	return newLarkFactory(f.storeIns)
}
