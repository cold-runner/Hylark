package response

import "github.com/cloudwego/kitex/pkg/kerrors"

type (
	BizErrCode  int32
	BizErrExtra map[string]string
)

const (
	ErrInternal BizErrCode = iota + 100001
	ErrTokenInvalid
	ErrBadRequest
	ErrPasswordIncorrect
	ErrUserAlreadyExist
	ErrUserNotExist
	ErrSmsCodeIncorrect
	ErrCategoryNotExist
	ErrTagNotExist
	ErrAlreadyFollow
)

var bizErrMap = map[BizErrCode]string{
	ErrInternal:          "internal error",
	ErrTokenInvalid:      "token is invalid",
	ErrBadRequest:        "bad request",
	ErrPasswordIncorrect: "password is incorrect",
	ErrUserAlreadyExist:  "user already exist",
	ErrUserNotExist:      "user not exist",
	ErrSmsCodeIncorrect:  "sms code is incorrect",
	ErrCategoryNotExist:  "category not exist",
	ErrTagNotExist:       "tag not exist",
	ErrAlreadyFollow:     "already follow",
}

func BizErr(code BizErrCode) error {
	return kerrors.NewBizStatusError(int32(code), bizErrMap[code])
}

func BizErrWithExtra(code BizErrCode, extra map[string]string) error {
	return kerrors.NewBizStatusErrorWithExtra(int32(code), bizErrMap[code], extra)
}
