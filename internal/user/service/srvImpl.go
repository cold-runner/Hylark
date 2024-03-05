package service

import (
	"context"
	"github.com/cold-runner/Hylark/gorm_gen/user_srv"
	"github.com/golang-jwt/jwt/v5"
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
		// 手机号校验
		reg, _ := regexp2.Compile(`^1[3456789]\d{9}$`, 0)
		if match, _ := reg.MatchString(req.GetPhone()); !match {
			klog.Infof("phone nubner is illegal")
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "phone number is illegal"})
		}
		// 校验验证码
		correct, err := s.validateSmsCode(ctx, req.GetPhone(), req.GetSmsCode())
		switch {
		case err != nil:
			klog.Infof("validate sms code failed! err: %v", err)
			return response.BizErr(response.ErrInternal)
		case !correct:
			klog.Infof("sms code is incorrect! err: %v", err)
			return response.BizErr(response.ErrSmsCodeIncorrect)
		}

		// 删除缓存
		if err = s.Cache.Del(ctx, req.GetPhone()); err != nil {
			klog.Warnf("delete cached sms code failed! err: %v", err)
		}

		// 校验密码强度：至少包括字母、数字,8~20位
		reg, _ = regexp2.Compile(`^(?=.*[a-zA-Z])(?=.*\d).{8,20}$`, 0)
		if match, _ := reg.MatchString(*req.Password); !match {
			klog.Infof("password strength too low")
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "password strength too low"})
		}
		return nil
	}(ctx, req); err != nil {
		return nil, err
	}

	// 加密密码
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		klog.Errorf("encrypt password failed! err: %v", err)
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
		klog.Info("user already exist")
		return nil, response.BizErr(response.ErrUserAlreadyExist)
	case err != nil:
		klog.Errorf("store register info failed! err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	return user.NewRegisterResponse(), nil
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
	return
}

func (s *Srv) PasswordLogin(ctx context.Context, req *user.PasswordLoginRequest) (r *user.PasswordLoginResponse, err error) {
	row, err := s.Factory.Lark().Produce(ctx, user_srv.Q.Lark.Phone.Eq(req.GetPhone()))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		klog.Infof("user not exist. phone: %v", req.GetPhone())
		return nil, response.BizErr(response.ErrUserAlreadyExist)
	case err != nil:
		klog.Errorf("unknow err. err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.GetRow().Password), []byte(req.GetPassword()))
	if err != nil {
		klog.Infof("password is incorrect. phone: %v", req.GetPhone())
		return nil, response.BizErr(response.ErrPasswordIncorrect)
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod(s.Config.JwtConfig.Algorithm), jwt.MapClaims{
		"iss": s.Config.JwtConfig.Issuer,
		"sub": s.Config.JwtConfig.Subject,
		"exp": s.Config.JwtConfig.ExpireTime * time.Minute,
		"uid": row.GetRow().ID.String(),
	})
	str, err := t.SignedString([]byte(s.Config.JwtConfig.Key))
	if err != nil {
		klog.Errorf("sign and issue token failed. err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	klog.Infof("login success. phone: %v", req.GetPhone())
	r = user.NewPasswordLoginResponse()
	r.SetToken(pkg.Convert(str))
	return
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
	entity, err := s.Factory.Lark().Produce(ctx, user_srv.Q.Lark.Phone.Eq(req.GetPhone()))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, response.BizErr(response.ErrUserNotExist)
	case err != nil:
		return nil, response.BizErr(response.ErrInternal)
	}

	// 上传照片
	err = entity.UploadStuCard(ctx, s.Oss, req.GetStuCardPhoto())
	if err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}

	// 持久化实体对象
	if err = s.Repository.Lark().Persist(ctx, entity.GetRow()); err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}

	// TODO 通知管理员进行认证
	return user.NewCertificateResponse(), nil
}
