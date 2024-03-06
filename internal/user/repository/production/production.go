package production

import (
	"github.com/cold-runner/Hylark/internal/user/repository"
	"github.com/cold-runner/Hylark/internal/user/store"
)

type repo struct {
	storeIns store.Store
}

func (r repo) Social() repository.SocialRepository {
	return newSocial(r.storeIns)
}

func (r repo) Lark() repository.LarkRepository {
	return newLarkRepository(r.storeIns)
}

func NewRepository(s store.Store) repository.Repository {
	return &repo{storeIns: s}
}
