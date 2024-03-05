package production

import (
	"github.com/cold-runner/Hylark/internal/post/repository"
	"github.com/cold-runner/Hylark/internal/post/store"
)

type repo struct {
	storeIns store.Store
}

func NewRepository(s store.Store) repository.Repository {
	return &repo{storeIns: s}
}

func (r repo) Post() repository.PostRepository {
	return newPostRepository(r.storeIns)
}

func (r repo) Category() repository.CategoryRepository {
	return newCategoryRepository(r.storeIns)
}

func (r repo) Tag() repository.TagRepository {
	return newTagRepository(r.storeIns)
}
