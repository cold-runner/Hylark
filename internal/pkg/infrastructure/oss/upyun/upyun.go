package upyun

import (
	"context"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"io"

	"github.com/upyun/go-sdk/v3/upyun"
)

type upy struct {
	upyunClient *upyun.UpYun
}

func NewOss(opt *instance.UpyunConfig) oss.Oss {
	client := &upy{
		upyunClient: upyun.NewUpYun(&upyun.UpYunConfig{
			Bucket:   opt.Bucket,
			Operator: opt.Operator,
			Password: opt.Password,
		}),
	}
	return client
}

func (u *upy) Upload(c context.Context, file io.Reader, fileName string, uploadPath string) (fileUrl string, err error) {
	obj := &upyun.PutObjectConfig{
		Path:              uploadPath,
		Reader:            file,
		Headers:           nil,
		UseMD5:            false,
		UseResumeUpload:   false,
		ResumePartSize:    0,
		MaxResumePutTries: 5,
	}

	if err := u.upyunClient.Put(obj); err != nil {
		// TODO 返回正确的文件URL
		return u.upyunClient.Hosts["domain"], err
	}
	return "", nil
}

func (u *upy) Delete(c context.Context, filaPath string) error {
	obj := &upyun.DeleteObjectConfig{
		Path:   filaPath,
		Async:  false,
		Folder: false,
	}
	if err := u.upyunClient.Delete(obj); err != nil {
		return err
	}
	return nil
}
