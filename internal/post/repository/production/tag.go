package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"github.com/cold-runner/Hylark/internal/post/store"
	"gorm.io/gen"
)

type tagRepository struct {
	s store.Store
}

func newTagRepository(s store.Store) *tagRepository {
	return &tagRepository{s: s}
}

func (t tagRepository) Persist(context.Context, *model.Post) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (t tagRepository) Get(c context.Context, cond ...gen.Condition) (*model.Tag, error) {
	row, err := t.s.Tag().Query(c, cond...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (t tagRepository) GetList() {
	//TODO implement me
	panic("implement me")
}

func (t tagRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
