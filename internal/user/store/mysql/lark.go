package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

type larkStore struct {
	query *user_srv.Query
}

func newLarkStore(q *user_srv.Query) *larkStore {
	return &larkStore{query: q}
}

func (l larkStore) Create(c context.Context, ov *model.Lark) error {
	if err := l.query.Lark.Create(ov); err != nil {
		return err
	}
	return nil
}

func (l larkStore) Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error {
	_, err := l.query.WithContext(c).Lark.Select(selectScopes...).Where(whereScopes...).Updates(lark)
	if err != nil {
		return err
	}
	return nil
}

func (l larkStore) Query(c context.Context, conds ...gen.Condition) (*model.Lark, error) {
	tar, err := l.query.Lark.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
