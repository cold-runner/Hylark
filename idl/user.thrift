namespace go user

include "common.thrift"

struct RegisterRequest {
    1: optional string phone
    2: optional string password
    3: optional string sms_code
}

struct RegisterResponse {
}

struct SendSmsCodeRequest {
    1: optional string phone
}

struct SendSmsCodeResponse {
    1: optional string code
}

struct CertificateRequest {
    1: optional string phone
    2: optional binary stu_card_photo
}

struct CertificateResponse {
}

