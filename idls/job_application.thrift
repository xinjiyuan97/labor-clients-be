namespace go job_application

include "common.thrift"

// 申请岗位请求
struct ApplyJobReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
}

// 申请岗位响应
struct ApplyJobResp {
    1: common.BaseResp base (api.body="base");
    2: i64 application_id (api.body="application_id" go.tag="json:\"application_id,string\"");
    3: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    4: string status (api.body="status");
    5: string applied_at (api.body="applied_at");
}

// 获取申请列表请求
struct GetApplicationListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string status (api.query="status");
}

// 获取申请列表响应
struct GetApplicationListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.JobApplicationInfo> applications (api.body="applications");
}

// 获取我的申请请求
struct GetMyApplicationsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string status (api.query="status");
}

// 获取我的申请响应
struct GetMyApplicationsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.JobApplicationInfo> applications (api.body="applications");
}

// 获取申请详情请求
struct GetApplicationDetailReq {
    1: i64 application_id (api.path="application_id", api.vd="$>0" go.tag="json:\"application_id,string\"");
}

// 获取申请详情响应
struct GetApplicationDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.JobApplicationInfo application (api.body="application");
}

// 取消申请请求
struct CancelApplicationReq {
    1: i64 application_id (api.path="application_id", api.vd="$>0" go.tag="json:\"application_id,string\"");
    2: string cancel_reason (api.body="cancel_reason", api.vd="len($)>0");
}

// 取消申请响应
struct CancelApplicationResp {
    1: common.BaseResp base (api.body="base");
}

// 确认申请请求
struct ConfirmApplicationReq {
    1: i64 application_id (api.path="application_id", api.vd="$>0" go.tag="json:\"application_id,string\"");
}

// 确认申请响应
struct ConfirmApplicationResp {
    1: common.BaseResp base (api.body="base");
}

service JobApplicationService {
    ApplyJobResp ApplyJob(1: ApplyJobReq request) (api.post="/api/v1/job-applications");
    GetApplicationListResp GetApplicationList(1: GetApplicationListReq request) (api.get="/api/v1/job-applications");
    GetMyApplicationsResp GetMyApplications(1: GetMyApplicationsReq request) (api.get="/api/v1/job-applications/my");
    GetApplicationDetailResp GetApplicationDetail(1: GetApplicationDetailReq request) (api.get="/api/v1/job-applications/:application_id");
    CancelApplicationResp CancelApplication(1: CancelApplicationReq request) (api.put="/api/v1/job-applications/:application_id/cancel");
    ConfirmApplicationResp ConfirmApplication(1: ConfirmApplicationReq request) (api.put="/api/v1/job-applications/:application_id/confirm");
}
