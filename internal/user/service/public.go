package service

import (
	"context"
	"github.com/cold-runner/Hylark/internal/pkg/response"
	"github.com/golang-jwt/jwt/v5"

	"github.com/pkg/errors"
)

func (s *Srv) validateSmsCode(ctx context.Context, phone, smsCode string) (bool, error) {
	cached, err := s.Cache.Get(ctx, phone)
	if err != nil {
		return false, err
	}
	if cached == nil {
		return false, errors.New("the code has not been sent.")
	}
	cachedCode := cached.(string)
	if cachedCode != smsCode {
		return false, errors.Errorf("sms code is incorrect. phone: %v, cached: %v, recv: %v", phone, cachedCode, smsCode)
	}
	return true, nil
}

type HylarkTokenClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func (s *Srv) validateToken(req hasToken) (*HylarkTokenClaims, error) {
	t, err := jwt.ParseWithClaims(req.GetToken(), &HylarkTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.JwtConfig.Key), nil
	})
	if err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}
	if !t.Valid {
		return nil, response.BizErr(response.ErrTokenInvalid)
	}
	claims, ok := t.Claims.(*HylarkTokenClaims)
	if !ok {
		return nil, response.BizErr(response.ErrTokenInvalid)
	}

	return claims, nil
}

type hasToken interface {
	GetToken() string
}
