package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/user/repository"
	"github.com/cold-runner/Hylark/internal/user/store"
)

type larkRepository struct {
	storeIns store.Store
}

func newLarkRepository(s store.Store) repository.LarkRepository {
	return &larkRepository{storeIns: s}
}

func (l larkRepository) Persist(c context.Context, lark *model.Lark) error {
	if err := l.storeIns.Create(c, lark); err != nil {
		return err
	}
	return nil
}
