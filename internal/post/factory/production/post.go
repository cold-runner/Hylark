package production

import (
	"github.com/cold-runner/Hylark/internal/post/entity"
	"github.com/cold-runner/Hylark/internal/post/factory"
	"github.com/cold-runner/Hylark/internal/post/store"
)

type postFactory struct {
	storeIns store.Store
}

func (p postFactory) Produce() *entity.Post {
	//TODO implement me
	panic("implement me")
}

func newPostFactory(s store.Store) factory.PostFactory {
	return &postFactory{storeIns: s}
}
