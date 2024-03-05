package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"github.com/cold-runner/Hylark/internal/post/store"
)

type postRepository struct {
	s store.Store
}

func newPostRepository(s store.Store) *postRepository {
	return &postRepository{s: s}
}

func (p postRepository) Persist(c context.Context, post *model.Post) error {
	if err := p.s.Post().Create(c, post); err != nil {
		return err
	}
	return nil
}

func (p postRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) GetList() {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
