package response

var (
	RespApiAlreadyExist = Resp{
		Code: "xxxxxx",
		Msg:  "api already exist",
	}
	RespApiNotExist = Resp{
		Code: "xxxxxx",
		Msg:  "api not exist",
	}
	RespSvcNotExist = Resp{
		Code: "xxxxxx",
		Msg:  "svc not exist",
	}
	RespBadParam = Resp{
		Code: "xxxxxx",
		Msg:  "validate failed",
	}
	RespServerError = Resp{
		Code: "xxxxxx",
		Msg:  "server internal error",
	}
	RespUnauthorized = Resp{
		Code: "xxxxxx",
		Msg:  "unauthorized request",
	}
)

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
