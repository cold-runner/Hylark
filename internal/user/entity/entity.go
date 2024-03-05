package entity

import (
	"bytes"
	"context"
	"github.com/cold-runner/Hylark/internal/pkg"

	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss"
)

type Lark struct {
	row *model.Lark
	// 其他必须字段
}

func (l *Lark) Instantiate(row *model.Lark) {
	l.row = row
}

func (l *Lark) UploadStuCard(c context.Context, oss oss.Oss, data []byte) error {
	url, err := oss.Upload(c, bytes.NewReader(data), "student_card."+pkg.MustFileType(data), "/lark"+l.row.ID.String())
	if err != nil {
		return err
	}

	l.row.StuCardURL = url
	return nil
}

func (l *Lark) GetRow() *model.Lark {
	return l.row
}

func (l *Lark) IsCertificate() bool {
	return l.row.State == 1
}

func (l *Lark) IsLegal() bool {
	return l.row.State == 0
}
