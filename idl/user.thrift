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

struct PasswordLoginRequest {
    1: optional string phone
    2: optional string password
}

struct PasswordLoginResponse {
    1: optional string token
}

struct CertificateRequest {
    1: optional string phone
    2: optional binary stu_card_photo
}

struct CertificateResponse {
}

struct UpdateUserInfoRequest {
    1: optional string token
    2: optional string gender
    3: optional string college
    4: optional string major
    5: optional string grade
    6: optional string province
    7: optional i8 age
    8: optional string introduction
    9: optional string avatar
}

struct UpdateUserInfoResponse {
}
