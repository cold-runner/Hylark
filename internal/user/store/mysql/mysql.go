package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/cold-runner/Hylark/internal/user/store"

	"github.com/pkg/errors"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type mysql struct {
	q *user_srv.Query
}

func NewStore(config *instance.MysqlConfig) store.Store {
	dbIns, err := instance.NewMySQL(config)
	if err != nil {
		panic(errors.Errorf("cannot establish store connection: %v", err))
	}
	return &mysql{user_srv.Use(dbIns)}
}

func (m mysql) Create(c context.Context, ov *model.Lark) error {
	if err := m.q.Lark.Create(ov); err != nil {
		return err
	}
	return nil
}

func (m mysql) Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error {
	_, err := m.q.WithContext(c).Lark.Select(selectScopes...).Where(whereScopes...).Updates(lark)
	if err != nil {
		return err
	}
	return nil
}

func (m mysql) Query(c context.Context, conds ...gen.Condition) (*model.Lark, error) {
	tar, err := m.q.Lark.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
