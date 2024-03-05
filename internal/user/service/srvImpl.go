package service

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/cold-runner/Hylark/gorm_gen/user_srv/model"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/internal/pkg/response"
	"github.com/cold-runner/Hylark/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Srv) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	if err = func(context.Context, *user.RegisterRequest) error {
		// 校验手机号
		correct, err := s.validateSmsCode(ctx, req.GetPhone(), req.GetSmsCode())
		switch {
		case err != nil:
			klog.Errorf("validate sms code failed! err: %v", err)
			return response.BizErr(response.ErrInternal)
		case !correct:
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "sms code is incorrect"})
		}

		// 删除缓存
		_ = s.Cache.Del(ctx, req.GetPhone()) // TODO 添加监控

		// 校验密码强度：至少包括字母、数字,8~20位
		reg, _ := regexp2.Compile(`^(?=.*[a-zA-Z])(?=.*\d).{8,20}$`, 0)
		if match, _ := reg.MatchString(*req.Password); !match {
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "password strength too low"})
		}
		return nil
	}(ctx, req); err != nil {
		return nil, err
	}

	// 加密密码
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}
	// 存储注册用户
	u := &model.Lark{
		ID:       uuid.New(),
		Phone:    req.GetPhone(),
		Password: string(encryptPassword),
		State:    1,
	}
	err = s.Repository.Lark().Persist(ctx, u)
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return nil, response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "user already exist"})
	case err != nil:
		klog.Errorf("store register info failed! err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	return user.NewRegisterResponse(), response.BizErr(response.ErrSuccess)
}

func (s *Srv) SendSmsCode(ctx context.Context, req *user.SendSmsCodeRequest) (r *user.SendSmsCodeResponse, err error) {
	// TODO 可配置的验证码位数和过期时间
	rcode := strconv.Itoa(rand.IntN(999) + 1000)

	// 缓存验证码
	err = s.Cache.SetExpiration(ctx, req.GetPhone(), rcode, time.Minute*10)
	if err != nil {
		klog.Errorf("cache sms code failed! err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	// 发送验证码
	if err = s.Sms.SendToSingle(ctx, req.GetPhone(), []string{rcode, "10"}); err != nil {
		klog.Errorf("send sms code failed! err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	r = user.NewSendSmsCodeResponse()
	r.SetCode(pkg.Convert(rcode))
	return r, response.BizErr(response.ErrSuccess)
}

func (s *Srv) Certificate(ctx context.Context, req *user.CertificateRequest) (r *user.CertificateResponse, err error) {
	if err = func(context.Context, *user.CertificateRequest) error {
		// 校验文件类型
		ft, err := pkg.FileTypeFromBinary(req.GetStuCardPhoto())
		switch {
		case err != nil:
			return response.BizErr(response.ErrInternal)
		case ft != "jpg" && ft != "png":
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "file type is illegal"})
		}
		return nil
	}(ctx, req); err != nil {
		return nil, err
	}

	// 工厂生产业务实体
	entity, err := s.Factory.Lark().Produce(ctx, user_srv.Lark.Phone.Eq(req.GetPhone()))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "user not exist"})
	case err != nil:
		return nil, response.BizErr(response.ErrInternal)
	}

	// 上传照片
	err = entity.UploadStuCard(ctx, s.Oss, req.GetStuCardPhoto())
	if err != nil {
		return nil, err
	}

	// 持久化实体对象
	if err = s.Repository.Lark().Persist(ctx, entity.GetRow()); err != nil {
		return nil, err
	}

	// TODO 通知管理员进行认证
	return user.NewCertificateResponse(), response.BizErr(response.ErrSuccess)
}
