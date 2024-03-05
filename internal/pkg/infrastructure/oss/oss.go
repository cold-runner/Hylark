package oss

import (
	"context"
	"io"
)

type Oss interface {
	Upload(c context.Context, file io.Reader, fileName string, uploadPath string) (fileUrl string, err error)
	Delete(c context.Context, filaPath string) error
}
