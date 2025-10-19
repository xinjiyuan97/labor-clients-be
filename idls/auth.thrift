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
    2: i64 user_id (api.body="user_id");
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
    2: i64 user_id (api.body="user_id");
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

service AuthService {
    RegisterResp Register(1: RegisterReq request) (api.post="/api/v1/auth/register");
    LoginResp Login(1: LoginReq request) (api.post="/api/v1/auth/login");
    LogoutResp Logout(1: LogoutReq request) (api.post="/api/v1/auth/logout");
    RefreshTokenResp RefreshToken(1: RefreshTokenReq request) (api.post="/api/v1/auth/refresh");
    GetUserProfileResp GetUserProfile(1: GetUserProfileReq request) (api.get="/api/v1/auth/profile");
}
