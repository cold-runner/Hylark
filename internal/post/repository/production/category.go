package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"github.com/cold-runner/Hylark/internal/post/store"
	"gorm.io/gen"
)

type categoryRepository struct {
	s store.Store
}

func (c categoryRepository) Persist(context.Context, *model.Post) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (c categoryRepository) Get(ctx context.Context, cond ...gen.Condition) (*model.Category, error) {
	row, err := c.s.Category().Query(ctx, cond...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (c categoryRepository) GetList() {
	//TODO implement me
	panic("implement me")
}

func (c categoryRepository) Delete() {
	//TODO implement me
	panic("implement me")
}

func newCategoryRepository(s store.Store) *categoryRepository {
	return &categoryRepository{s: s}
}
