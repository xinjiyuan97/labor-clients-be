namespace go auth

include "common.thrift"

// 用户注册请求
struct RegisterReq {
    1: string phone (api.query="phone", api.vd="len($)>0");
    2: string password (api.query="password", api.vd="len($)>0");
    3: string role (api.query="role", api.vd="len($)>0");
    4: string username (api.query="username", api.vd="len($)>0");
}

// 用户注册响应
struct RegisterResp {
    1: common.BaseResp base (api.body="base");
    2: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    3: string token (api.body="token");
    4: string expires_at (api.body="expires_at");
}

// 用户登录请求
struct LoginReq {
    1: string phone (api.query="phone", api.vd="len($)>0");
    2: string password (api.query="password", api.vd="len($)>0");
}

// 用户登录响应
struct LoginResp {
    1: common.BaseResp base (api.body="base");
    2: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    3: string token (api.body="token");
    4: string expires_at (api.body="expires_at");
}

// 用户登出请求
struct LogoutReq {
    1: string token (api.header="Authorization", api.vd="len($)>0");
}

// 用户登出响应
struct LogoutResp {
    1: common.BaseResp base (api.body="base");
}

// 刷新Token请求
struct RefreshTokenReq {
    1: string refresh_token (api.query="refresh_token", api.vd="len($)>0");
}

// 刷新Token响应
struct RefreshTokenResp {
    1: common.BaseResp base (api.body="base");
    2: string token (api.body="token");
    3: string expires_at (api.body="expires_at");
}

// 获取用户信息请求
struct GetUserProfileReq {
}

// 获取用户信息响应
struct GetUserProfileResp {
    1: common.BaseResp base (api.body="base");
    2: common.UserInfo user_info (api.body="user_info");
    3: common.WorkerInfo worker_info (api.body="worker_info");
}

// 修改密码请求
struct ChangePasswordReq {
    1: string old_password (api.body="old_password", api.vd="len($)>0");
    2: string new_password (api.body="new_password", api.vd="len($)>0");
}

// 修改密码响应
struct ChangePasswordResp {
    1: common.BaseResp base (api.body="base");
}

// 发送短信验证码请求
struct SendSMSCodeReq {
    1: string phone (api.body="phone", api.vd="len($)>0");
}

// 发送短信验证码响应
struct SendSMSCodeResp {
    1: common.BaseResp base (api.body="base");
    2: string code (api.body="code");
    3: i32 expires_in (api.body="expires_in");
}

// 短信验证码登录请求
struct LoginWithSMSCodeReq {
    1: string phone (api.body="phone", api.vd="len($)>0");
    2: string code (api.body="code", api.vd="len($)>0");
}

// 第三方登录绑定请求
struct ThirdPartyLoginBindReq {
    1: string platform (api.body="platform", api.vd="len($)>0");
    2: string openid (api.body="openid", api.vd="len($)>0");
    3: string unionid (api.body="unionid");
    4: string appid (api.body="appid", api.vd="len($)>0");
    5: string phone (api.body="phone", api.vd="len($)>0");
    6: string code (api.body="code", api.vd="len($)>0");
    7: string nickname (api.body="nickname");
    8: string avatar (api.body="avatar");
}

// 第三方登录绑定响应
struct ThirdPartyLoginBindResp {
    1: common.BaseResp base (api.body="base");
    2: bool is_new_user (api.body="is_new_user");
    3: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    4: string token (api.body="token");
    5: string expires_at (api.body="expires_at");
}

service AuthService {
    RegisterResp Register(1: RegisterReq request) (api.post="/api/v1/auth/register");
    LoginResp Login(1: LoginReq request) (api.post="/api/v1/auth/login");
    LogoutResp Logout(1: LogoutReq request) (api.post="/api/v1/auth/logout");
    RefreshTokenResp RefreshToken(1: RefreshTokenReq request) (api.post="/api/v1/auth/refresh");
    GetUserProfileResp GetUserProfile(1: GetUserProfileReq request) (api.get="/api/v1/auth/profile");
    ChangePasswordResp ChangePassword(1: ChangePasswordReq request) (api.post="/api/v1/auth/change-password");
    SendSMSCodeResp SendSMSCode(1: SendSMSCodeReq request) (api.post="/api/v1/auth/send-sms-code");
    LoginResp LoginWithSMSCode(1: LoginWithSMSCodeReq request) (api.post="/api/v1/auth/login-with-sms");
    ThirdPartyLoginBindResp ThirdPartyLoginBind(1: ThirdPartyLoginBindReq request) (api.post="/api/v1/auth/third-party-bind");
}
