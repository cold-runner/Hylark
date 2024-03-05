package production

import (
	"github.com/cold-runner/Hylark/internal/post/factory"
	"github.com/cold-runner/Hylark/internal/post/store"
)

type fact struct {
	storeIns store.Store
}

func NewFactory(s store.Store) factory.Factory {
	return &fact{storeIns: s}
}

func (f fact) Post() factory.PostFactory {
	return newPostFactory(f.storeIns)
}
