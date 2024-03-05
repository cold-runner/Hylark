package factory

import "github.com/cold-runner/Hylark/internal/post/entity"

type Factory interface {
	Post() PostFactory
}

type PostFactory interface {
	Produce() *entity.Post
}
