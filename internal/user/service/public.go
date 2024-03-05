package service

import "context"

func (s *Srv) validateSmsCode(ctx context.Context, phone, smsCode string) (bool, error) {
	cached, err := s.Cache.Get(ctx, phone)
	if err != nil {
		return false, err
	}
	if cached != smsCode {
		return false, nil
	}
	return true, nil
}
