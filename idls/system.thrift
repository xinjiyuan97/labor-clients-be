namespace go system

include "common.thrift"

// 获取系统配置请求
struct GetSystemConfigReq {
    1: string config_key (api.query="config_key");
}

// 获取系统配置响应
struct GetSystemConfigResp {
    1: common.BaseResp base (api.body="base");
    2: map<string, string> configs (api.body="configs");
}

// 获取版本信息请求
struct GetVersionReq {
}

// 获取版本信息响应
struct GetVersionResp {
    1: common.BaseResp base (api.body="base");
    2: string version (api.body="version");
    3: string build_time (api.body="build_time");
    4: string git_commit (api.body="git_commit");
}

// 提交反馈请求
struct SubmitFeedbackReq {
    1: string feedback_type (api.body="feedback_type", api.vd="len($)>0");
    2: string content (api.body="content", api.vd="len($)>0");
    3: string contact_info (api.body="contact_info");
    4: list<string> attachments (api.body="attachments");
}

// 提交反馈响应
struct SubmitFeedbackResp {
    1: common.BaseResp base (api.body="base");
    2: i64 feedback_id (api.body="feedback_id" go.tag="json:\"feedback_id,string\"");
}

// 获取公告列表请求
struct GetNoticeListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string notice_type (api.query="notice_type");
}

// 获取公告列表响应
struct GetNoticeListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.CommunityPostInfo> notices (api.body="notices");
}

// 获取公告详情请求
struct GetNoticeDetailReq {
    1: i64 notice_id (api.path="notice_id", api.vd="$>0" go.tag="json:\"notice_id,string\"");
}

// 获取公告详情响应
struct GetNoticeDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.CommunityPostInfo notice (api.body="notice");
}

// 获取微信用户信息请求
struct GetWeChatUserInfoReq {
}

// 获取微信用户信息响应
struct GetWeChatUserInfoResp {
    1: common.BaseResp base (api.body="base");
    2: string openid (api.body="openid");
    3: string unionid (api.body="unionid");
    4: string appid (api.body="appid");
    5: string env (api.body="env");
    6: string cloudbase_access_token (api.body="cloudbase_access_token");
}

service SystemService {
    GetSystemConfigResp GetSystemConfig(1: GetSystemConfigReq request) (api.get="/api/v1/system/config");
    GetVersionResp GetVersion(1: GetVersionReq request) (api.get="/api/v1/system/version");
    SubmitFeedbackResp SubmitFeedback(1: SubmitFeedbackReq request) (api.post="/api/v1/system/feedback");
    GetNoticeListResp GetNoticeList(1: GetNoticeListReq request) (api.get="/api/v1/system/notices");
    GetNoticeDetailResp GetNoticeDetail(1: GetNoticeDetailReq request) (api.get="/api/v1/system/notices/:notice_id");
    GetWeChatUserInfoResp GetWeChatUserInfo(1: GetWeChatUserInfoReq request) (api.get="/api/v1/system/wechat-user-info");
}
