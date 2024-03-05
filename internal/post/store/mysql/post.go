package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv"
	"github.com/cold-runner/Hylark/internal/post/store"

	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

type post struct {
	q *post_srv.Query
}

func newPost(m mysql) store.PostStore {
	return &post{q: m.q}
}

func (m post) Create(c context.Context, post *model.Post) error {
	if err := m.q.Post.Create(post); err != nil {
		return err
	}
	return nil
}

func (m post) Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.Post) error {
	_, err := m.q.WithContext(c).Post.Select(selectScopes...).Where(whereScopes...).Updates(lark)
	if err != nil {
		return err
	}
	return nil
}

func (m post) Query(c context.Context, conds ...gen.Condition) (*model.Post, error) {
	tar, err := m.q.Post.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
