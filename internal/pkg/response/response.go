package response

import "github.com/cloudwego/kitex/pkg/kerrors"

type (
	BizErrCode  int32
	BizErrExtra map[string]string
)

const (
	ErrSuccess BizErrCode = iota + 100000
	ErrInternal
	ErrBadRequest
)

var bizErrMap = map[BizErrCode]string{
	ErrSuccess:    "ok",
	ErrInternal:   "internal error",
	ErrBadRequest: "bad request",
}

func BizErr(code BizErrCode) error {
	return kerrors.NewBizStatusError(int32(code), bizErrMap[code])
}

func BizErrWithExtra(code BizErrCode, extra map[string]string) error {
	return kerrors.NewBizStatusErrorWithExtra(int32(code), bizErrMap[code], extra)
}
