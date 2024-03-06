package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/user/repository"
	"github.com/cold-runner/Hylark/internal/user/store"
	"github.com/pkg/errors"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type larkRepository struct {
	storeIns store.Store
}

func newLarkRepository(s store.Store) repository.LarkRepository {
	return &larkRepository{storeIns: s}
}

func (l larkRepository) Persist(c context.Context, lark *model.Lark) error {
	if err := l.storeIns.Lark().Create(c, lark); err != nil {
		return err
	}
	return nil
}

func (l larkRepository) Update(c context.Context, lark *model.Lark) error {
	if err := l.storeIns.Lark().Update(c,
		[]field.Expr{user_srv.Lark.ALL},
		[]gen.Condition{user_srv.Q.Lark.ID.Eq(lark.ID)},
		lark); err != nil {

		return err
	}
	return nil
}

func (l larkRepository) IsExist(c context.Context, conds ...gen.Condition) (bool, error) {
	_, err := l.storeIns.Lark().Query(c, conds...)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return false, nil
	case err != nil:
		return false, err
	}
	return true, nil
}
