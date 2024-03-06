package entity

import (
	"bytes"
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss"
	"github.com/cold-runner/Hylark/internal/user/repository"
	"github.com/google/uuid"
)

type Lark struct {
	row *model.Lark
	// 其他必须字段
}

func (l *Lark) Instantiate(row *model.Lark) *Lark {
	l.row = row
	return l
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

func (l *Lark) Follow(c context.Context, repo repository.Repository, subjectId, objectId string) error {
	social := &model.UserInteraction{
		ID:         uuid.New(),
		UserID:     subjectId,
		FollowedID: objectId,
	}
	if err := repo.Social().Persist(c, social); err != nil {
		return err
	}
	return nil
}
