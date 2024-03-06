package mysql

import (
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/cold-runner/Hylark/internal/user/store"

	"github.com/pkg/errors"
)

type mysql struct {
	q *user_srv.Query
}

func NewStore(config *instance.MysqlConfig) store.Store {
	dbIns, err := instance.NewMySQL(config)
	if err != nil {
		panic(errors.Errorf("cannot establish store connection: %v", err))
	}
	user_srv.SetDefault(dbIns)
	return &mysql{user_srv.Q}
}

func (m mysql) Lark() store.LarkStore {
	return newLarkStore(m.q)
}

func (m mysql) Social() store.SocialStore {
	return newSocialStore(m.q)
}
