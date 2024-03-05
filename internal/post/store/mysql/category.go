package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/internal/post/store"

	"github.com/cold-runner/Hylark/gorm_gen/post_srv"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type category struct {
	q *post_srv.Query
}

func newCategory(m mysql) store.CategoryStore {
	return &category{m.q}
}

func (c category) Create(ctx context.Context, category *model.Category) error {
	if err := c.q.Category.Create(category); err != nil {
		return err
	}
	return nil
}

func (c category) Update(ctx context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, category *model.Category) error {
	_, err := c.q.WithContext(ctx).Category.Select(selectScopes...).Where(whereScopes...).Updates(category)
	if err != nil {
		return err
	}
	return nil
}

func (c category) Query(ctx context.Context, conds ...gen.Condition) (*model.Category, error) {
	tar, err := c.q.Category.WithContext(ctx).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
