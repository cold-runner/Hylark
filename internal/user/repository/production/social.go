package production

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/user/store"
)

type socialRepository struct {
	storeIns store.Store
}

func newSocial(s store.Store) *socialRepository {
	return &socialRepository{storeIns: s}
}

func (s socialRepository) Persist(c context.Context, row *model.UserInteraction) error {
	err := s.storeIns.Social().Create(c, row)
	if err != nil {
		return err
	}
	return nil
}
