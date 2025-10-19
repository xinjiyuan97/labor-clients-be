namespace go user

include "common.thrift"

// 更新个人信息请求
struct UpdateProfileReq {
    1: string username (api.body="username");
    2: string avatar (api.body="avatar");
    3: common.WorkerInfo worker_info (api.body="worker_info");
}

// 更新个人信息响应
struct UpdateProfileResp {
    1: common.BaseResp base (api.body="base");
    2: common.UserInfo user_info (api.body="user_info");
    3: common.WorkerInfo worker_info (api.body="worker_info");
}

// 上传头像请求
struct UploadAvatarReq {
    1: string avatar_file (api.form="avatar_file", api.vd="len($)>0");
}

// 上传头像响应
struct UploadAvatarResp {
    1: common.BaseResp base (api.body="base");
    2: string avatar_url (api.body="avatar_url");
}

// 上传认证文件请求
struct UploadCertReq {
    1: string cert_file (api.form="cert_file", api.vd="len($)>0");
    2: string cert_type (api.form="cert_type", api.vd="len($)>0");
}

// 上传认证文件响应
struct UploadCertResp {
    1: common.BaseResp base (api.body="base");
    2: string cert_type (api.body="cert_type");
    3: string file_url (api.body="file_url");
    4: i64 cert_id (api.body="cert_id");
}

// 获取我的收藏请求
struct GetMyFavoritesReq {
    1: common.PageReq page_req (api.body="page_req");
}

// 获取我的收藏响应
struct GetMyFavoritesResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.UserFavoriteJobInfo> favorites (api.body="favorites");
}

// 收藏岗位请求
struct FavoriteJobReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0");
}

// 收藏岗位响应
struct FavoriteJobResp {
    1: common.BaseResp base (api.body="base");
    2: i64 favorite_id (api.body="favorite_id");
}

// 取消收藏请求
struct UnfavoriteJobReq {
    1: i64 job_id (api.path="job_id", api.vd="$>0");
}

// 取消收藏响应
struct UnfavoriteJobResp {
    1: common.BaseResp base (api.body="base");
}

// 获取收入统计请求
struct GetIncomeReq {
    1: string period (api.query="period");
    2: i32 year (api.query="year");
    3: i32 month (api.query="month");
}

// 获取收入统计响应
struct GetIncomeResp {
    1: common.BaseResp base (api.body="base");
    2: double total_income (api.body="total_income");
    3: double pending_income (api.body="pending_income");
    4: double paid_income (api.body="paid_income");
    5: list<common.PaymentInfo> payments (api.body="payments");
}

// 获取收入详情请求
struct GetIncomeDetailReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
    4: string status (api.query="status");
}

// 获取收入详情响应
struct GetIncomeDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.PaymentInfo> payments (api.body="payments");
}

// 获取个人表现请求
struct GetPerformanceReq {
    1: string period (api.query="period");
}

// 获取个人表现响应
struct GetPerformanceResp {
    1: common.BaseResp base (api.body="base");
    2: i32 completed_jobs (api.body="completed_jobs");
    3: i32 total_applications (api.body="total_applications");
    4: double success_rate (api.body="success_rate");
    5: double average_rating (api.body="average_rating");
    6: i32 total_rating_count (api.body="total_rating_count");
}

service UserService {
    UpdateProfileResp UpdateProfile(1: UpdateProfileReq request) (api.put="/api/v1/user/profile");
    UploadAvatarResp UploadAvatar(1: UploadAvatarReq request) (api.post="/api/v1/user/upload-avatar");
    UploadCertResp UploadCert(1: UploadCertReq request) (api.post="/api/v1/user/upload-cert");
    GetMyFavoritesResp GetMyFavorites(1: GetMyFavoritesReq request) (api.get="/api/v1/user/favorites");
    FavoriteJobResp FavoriteJob(1: FavoriteJobReq request) (api.post="/api/v1/user/favorites");
    UnfavoriteJobResp UnfavoriteJob(1: UnfavoriteJobReq request) (api.delete="/api/v1/user/favorites/:job_id");
    GetIncomeResp GetIncome(1: GetIncomeReq request) (api.get="/api/v1/user/income");
    GetIncomeDetailResp GetIncomeDetail(1: GetIncomeDetailReq request) (api.get="/api/v1/user/income/detail");
    GetPerformanceResp GetPerformance(1: GetPerformanceReq request) (api.get="/api/v1/user/performance");
}
