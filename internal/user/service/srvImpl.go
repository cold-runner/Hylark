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
		return nil, response.BizErr(response.ErrUserNotExist)
	case err != nil:
		klog.Errorf("unknow err. err: %v", err)
		return nil, response.BizErr(response.ErrInternal)
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.GetRow().Password), []byte(req.GetPassword()))
	if err != nil {
		klog.Infof("password is incorrect. phone: %v", req.GetPhone())
		return nil, response.BizErr(response.ErrPasswordIncorrect)
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod(s.Config.JwtConfig.Algorithm), HylarkTokenClaims{
		UUID: row.GetRow().ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Config.JwtConfig.Issuer,
			Subject:   s.Config.JwtConfig.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.Config.JwtConfig.ExpireTime * time.Minute)),
		},
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

func (s *Srv) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (r *user.UpdateUserInfoResponse, err error) {
	var userId string
	if err = func(request *user.UpdateUserInfoRequest) error {
		t, err := jwt.ParseWithClaims(req.GetToken(), &HylarkTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config.JwtConfig.Key), nil
		})
		if err != nil {
			return response.BizErr(response.ErrInternal)
		}
		if !t.Valid {
			return response.BizErr(response.ErrTokenInvalid)
		}
		if claims, ok := t.Claims.(*HylarkTokenClaims); ok {
			userId = claims.UUID
		} else {
			return response.BizErr(response.ErrTokenInvalid)
		}
		gender := req.GetGender()
		if gender != "男" && gender != "女" && gender != "保密" {
			klog.Infof("gender is illegal")
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "gender is illegal"})
		}
		// TODO 学院校验
		// TODO 专业校验
		reg, _ := regexp2.Compile(`^(大一|大二|大三|大四|研究生|毕业生)$`, 0)
		if match, _ := reg.MatchString(req.GetGrade()); !match {
			klog.Infof("grade is illegal")
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "grade is illegal"})
		}
		// TODO 省份校验
		// TODO 年龄校验
		// TODO 个人介绍校验
		return nil
	}(req); err != nil {
		return nil, err
	}
	id, err := uuid.Parse(userId)
	row := &model.Lark{
		ID:           id,
		UpdatedAt:    time.Now(),
		Gender:       nil,
		College:      req.GetCollege(),
		Major:        req.GetMajor(),
		Grade:        req.GetGrade(),
		Province:     req.GetProvince(),
		Age:          int32(req.GetAge()),
		Introduction: req.GetIntroduction(),
		Avatar:       pkg.Convert(req.GetAvatar()),
	}
	err = s.Repository.Lark().Update(ctx, row)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, response.BizErr(response.ErrUserNotExist)
	case err != nil:
		return nil, response.BizErr(response.ErrInternal)
	}
	return user.NewUpdateUserInfoResponse(), nil
}
