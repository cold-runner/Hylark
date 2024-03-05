package mysql

import (
	"github.com/cold-runner/Hylark/gorm_gen/post_srv"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/cold-runner/Hylark/internal/post/store"

	"github.com/pkg/errors"
)

type mysql struct {
	q *post_srv.Query
}

func NewStore(config *instance.MysqlConfig) store.Store {
	dbIns, err := instance.NewMySQL(config)
	if err != nil {
		panic(errors.Errorf("cannot establish store connection: %v", err))
	}
	post_srv.SetDefault(dbIns)
	return &mysql{post_srv.Q}
}

func (m mysql) Post() store.PostStore {
	return newPost(m)
}

func (m mysql) Tag() store.TagStore {
	return newTag(m)
}

func (m mysql) Category() store.CategoryStore {
	return newCategory(m)
}
