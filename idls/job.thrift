namespace go job

include "common.thrift"

// 获取岗位列表请求
struct GetJobListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: i32 category_id (api.query="category_id");
    3: string job_type (api.query="job_type");
    4: double distance (api.query="distance");
    5: double salary_min (api.query="salary_min");
    6: double salary_max (api.query="salary_max");
    7: string sort (api.query="sort");
    8: string order (api.query="order");
    9: double latitude (api.query="latitude");
    10: double longitude (api.query="longitude");
}

// 获取岗位列表响应
struct GetJobListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.JobInfo> jobs (api.body="jobs");
}

// 获取岗位详情请求
struct GetJobDetailReq {
    1: i64 job_id (api.path="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
}

// 获取岗位详情响应
struct GetJobDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.JobInfo job (api.body="job");
}

// 获取推荐岗位请求
struct GetRecommendJobsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: double latitude (api.query="latitude");
    3: double longitude (api.query="longitude");
}

// 获取推荐岗位响应
struct GetRecommendJobsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.JobInfo> jobs (api.body="jobs");
}

// 搜索岗位请求
struct SearchJobsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string keyword (api.query="keyword", api.vd="len($)>0");
    3: string location (api.query="location");
    4: double salary_min (api.query="salary_min");
    5: double salary_max (api.query="salary_max");
    6: string job_type (api.query="job_type");
    7: i32 category_id (api.query="category_id");
}

// 搜索岗位响应
struct SearchJobsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.JobInfo> jobs (api.body="jobs");
}

// 获取岗位分类请求
struct GetJobCategoriesReq {
    1: i32 parent_id (api.query="parent_id");
}

// 获取岗位分类响应
struct GetJobCategoriesResp {
    1: common.BaseResp base (api.body="base");
    2: list<common.JobCategoryInfo> categories (api.body="categories");
}

// 获取品牌列表请求
struct GetBrandListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string name (api.query="name");
    3: string auth_status (api.query="auth_status");
}

// 获取品牌列表响应
struct GetBrandListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.BrandInfo> brands (api.body="brands");
}

service JobService {
    GetJobListResp GetJobList(1: GetJobListReq request) (api.get="/api/v1/jobs");
    GetJobDetailResp GetJobDetail(1: GetJobDetailReq request) (api.get="/api/v1/jobs/:job_id");
    GetRecommendJobsResp GetRecommendJobs(1: GetRecommendJobsReq request) (api.get="/api/v1/jobs/recommend");
    SearchJobsResp SearchJobs(1: SearchJobsReq request) (api.get="/api/v1/jobs/search");
    GetJobCategoriesResp GetJobCategories(1: GetJobCategoriesReq request) (api.get="/api/v1/job-categories");
    GetBrandListResp GetBrandList(1: GetBrandListReq request) (api.get="/api/v1/brands");
}
