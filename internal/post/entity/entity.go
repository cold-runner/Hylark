package entity

import (
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
)

type Post struct {
	row  *model.Post
	tags []string
	state
}
