package mysql

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type socialStore struct {
	query *user_srv.Query
}

func newSocialStore(q *user_srv.Query) *socialStore {
	return &socialStore{query: q}
}

func (s socialStore) Create(c context.Context, social *model.UserInteraction) error {
	if err := s.query.UserInteraction.WithContext(c).Create(social); err != nil {
		return err
	}
	return nil
}

func (s socialStore) Update(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, social *model.UserInteraction) error {
	_, err := s.query.WithContext(c).UserInteraction.Select(selectScopes...).Where(whereScopes...).Updates(social)
	if err != nil {
		return err
	}
	return nil
}

func (s socialStore) Query(c context.Context, conds ...gen.Condition) (*model.UserInteraction, error) {
	tar, err := s.query.UserInteraction.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return tar, nil
}
