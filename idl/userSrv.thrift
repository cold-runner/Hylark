namespace go user

include "user.thrift"

service srv {
    user.RegisterResponse Register(user.RegisterRequest req)
    user.SendSmsCodeResponse SendSmsCode(user.SendSmsCodeRequest req)
    user.PasswordLoginResponse PasswordLogin(user.PasswordLoginRequest req)
    user.CertificateResponse Certificate(user.CertificateRequest req)
}
