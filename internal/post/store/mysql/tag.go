package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/internal/post/store"

	"github.com/cold-runner/Hylark/gorm_gen/post_srv"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type tag struct {
	q *post_srv.Query
}

func newTag(m mysql) store.TagStore {
	return &tag{m.q}
}

func (t tag) Create(c context.Context, tag *model.Tag) error {
	if err := t.q.Tag.Create(tag); err != nil {
		return err
	}
	return nil
}

func (t tag) Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, tag *model.Tag) error {
	_, err := t.q.WithContext(c).Tag.Select(selectScopes...).Where(whereScopes...).Updates(tag)
	if err != nil {
		return err
	}
	return nil
}

func (t tag) Query(c context.Context, conds ...gen.Condition) (*model.Tag, error) {
	tar, err := t.q.Tag.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
