package service

import (
	"context"

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
