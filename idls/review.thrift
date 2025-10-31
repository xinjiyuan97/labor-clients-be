namespace go review

include "common.thrift"

// 获取收到的评价请求
struct GetReceivedReviewsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
}

// 获取收到的评价响应
struct GetReceivedReviewsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.ReviewInfo> reviews (api.body="reviews");
}

// 获取给出的评价请求
struct GetGivenReviewsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
}

// 获取给出的评价响应
struct GetGivenReviewsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.ReviewInfo> reviews (api.body="reviews");
}

// 获取评价详情请求
struct GetReviewDetailReq {
    1: i64 review_id (api.path="review_id", api.vd="$>0" go.tag="json:\"review_id,string\"");
}

// 获取评价详情响应
struct GetReviewDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.ReviewInfo review (api.body="review");
}

// 发布评价请求
struct CreateReviewReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
    2: i64 employer_id (api.body="employer_id", api.vd="$>0" go.tag="json:\"employer_id,string\"");
    3: i32 rating (api.body="rating", api.vd="$>=1&&$<=5");
    4: string content (api.body="content", api.vd="len($)>0");
    5: string review_type (api.body="review_type", api.vd="len($)>0");
}

// 发布评价响应
struct CreateReviewResp {
    1: common.BaseResp base (api.body="base");
    2: i64 review_id (api.body="review_id" go.tag="json:\"review_id,string\"");
}

// 更新评价请求
struct UpdateReviewReq {
    1: i64 review_id (api.path="review_id", api.vd="$>0" go.tag="json:\"review_id,string\"");
    2: i32 rating (api.body="rating", api.vd="$>=1&&$<=5");
    3: string content (api.body="content", api.vd="len($)>0");
}

// 更新评价响应
struct UpdateReviewResp {
    1: common.BaseResp base (api.body="base");
    2: common.ReviewInfo review (api.body="review");
}

service ReviewService {
    GetReceivedReviewsResp GetReceivedReviews(1: GetReceivedReviewsReq request) (api.get="/api/v1/reviews/received");
    GetGivenReviewsResp GetGivenReviews(1: GetGivenReviewsReq request) (api.get="/api/v1/reviews/given");
    GetReviewDetailResp GetReviewDetail(1: GetReviewDetailReq request) (api.get="/api/v1/reviews/:review_id");
    CreateReviewResp CreateReview(1: CreateReviewReq request) (api.post="/api/v1/reviews");
    UpdateReviewResp UpdateReview(1: UpdateReviewReq request) (api.put="/api/v1/reviews/:review_id");
}
