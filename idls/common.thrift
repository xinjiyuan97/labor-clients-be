namespace go common

// 基础响应结构
struct BaseResp {
    1: i32 code (api.body="code");
    2: string message (api.body="message");
    3: string timestamp (api.body="timestamp");
}

// 分页请求
struct PageReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
}

// 分页响应
struct PageResp {
    1: i32 total (api.body="total");
    2: i32 page (api.body="page");
    3: i32 limit (api.body="limit");
}

// 用户基础信息
struct UserInfo {
    1: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    2: string username (api.body="username");
    3: string phone (api.body="phone");
    4: string avatar (api.body="avatar");
    5: string role (api.body="role");
}

// 零工详细信息
struct WorkerInfo {
    1: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    2: string real_name (api.body="real_name");
    3: string gender (api.body="gender");
    4: i32 age (api.body="age");
    5: string education (api.body="education");
    6: double height (api.body="height");
    7: string introduction (api.body="introduction");
    8: string work_experience (api.body="work_experience");
    9: double expected_salary (api.body="expected_salary");
}

// 品牌信息
struct BrandInfo {
    1: i64 brand_id (api.body="brand_id" go.tag="json:\"brand_id,string\"");
    2: string name (api.body="name");
    3: string logo (api.body="logo");
    4: string description (api.body="description");
    5: string auth_status (api.body="auth_status");
}

// 雇主信息
struct EmployerInfo {
    1: i64 employer_id (api.body="employer_id" go.tag="json:\"employer_id,string\"");
    2: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    3: i64 brand_id (api.body="brand_id" go.tag="json:\"brand_id,string\"");
    4: string company_name (api.body="company_name");
    5: string contact_person (api.body="contact_person");
    6: string contact_phone (api.body="contact_phone");
    7: string auth_status (api.body="auth_status");
}

// 岗位分类信息
struct JobCategoryInfo {
    1: i32 category_id (api.body="category_id");
    2: string name (api.body="name");
    3: string description (api.body="description");
    4: i32 parent_id (api.body="parent_id");
    5: i32 sort_order (api.body="sort_order");
}

// 岗位标签信息
struct JobTagInfo {
    1: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    2: string tag_name (api.body="tag_name");
    3: string tag_type (api.body="tag_type");
}

// 岗位信息
struct JobInfo {
    1: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    2: string title (api.body="title");
    3: string job_type (api.body="job_type");
    4: string description (api.body="description");
    5: double salary (api.body="salary");
    6: string salary_unit (api.body="salary_unit");
    7: string location (api.body="location");
    8: double latitude (api.body="latitude");
    9: double longitude (api.body="longitude");
    10: string requirements (api.body="requirements");
    11: string benefits (api.body="benefits");
    12: string start_time (api.body="start_time");
    13: string end_time (api.body="end_time");
    14: string status (api.body="status");
    15: i32 applicant_count (api.body="applicant_count");
    16: i32 max_applicants (api.body="max_applicants");
    17: i64 employer_id (api.body="employer_id" go.tag="json:\"employer_id,string\"");
    18: i64 brand_id (api.body="brand_id" go.tag="json:\"brand_id,string\"");
    19: i64 category_id (api.body="category_id" go.tag="json:\"category_id,string\"");
    20: double distance (api.body="distance");
    21: BrandInfo brand_info (api.body="brand_info");
    22: EmployerInfo employer_info (api.body="employer_info");
    23: JobCategoryInfo category_info (api.body="category_info");
    24: list<JobTagInfo> tags (api.body="tags");
}

// 日程信息
struct ScheduleInfo {
    1: i64 schedule_id (api.body="schedule_id" go.tag="json:\"schedule_id,string\"");
    2: i64 worker_id (api.body="worker_id" go.tag="json:\"worker_id,string\"");
    3: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    4: string title (api.body="title");
    5: string start_time (api.body="start_time");
    6: string end_time (api.body="end_time");
    7: string location (api.body="location");
    8: string notes (api.body="notes");
    9: string status (api.body="status");
    10: i32 reminder_minutes (api.body="reminder_minutes");
    11: JobInfo job_info (api.body="job_info");
}

// 岗位申请信息
struct JobApplicationInfo {
    1: i64 application_id (api.body="application_id" go.tag="json:\"application_id,string\"");
    2: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    3: i64 worker_id (api.body="worker_id" go.tag="json:\"worker_id,string\"");
    4: string status (api.body="status");
    5: string applied_at (api.body="applied_at");
    6: string confirmed_at (api.body="confirmed_at");
    7: string cancel_reason (api.body="cancel_reason");
    8: i32 worker_rating (api.body="worker_rating");
    9: i32 employer_rating (api.body="employer_rating");
    10: string review (api.body="review");
    11: JobInfo job_info (api.body="job_info");
}

// 消息信息
struct MessageInfo {
    1: i64 message_id (api.body="message_id" go.tag="json:\"message_id,string\"");
    2: i64 from_user (api.body="from_user" go.tag="json:\"from_user,string\"");
    3: i64 to_user (api.body="to_user" go.tag="json:\"to_user,string\"");
    4: string message_type (api.body="message_type");
    5: string content (api.body="content");
    6: string msg_category (api.body="msg_category");
    7: bool is_read (api.body="is_read");
    8: string created_at (api.body="created_at");
}

// 考勤记录信息
struct AttendanceRecordInfo {
    1: i32 record_id (api.body="record_id");
    2: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    3: i64 worker_id (api.body="worker_id" go.tag="json:\"worker_id,string\"");
    4: string check_in (api.body="check_in");
    5: string check_out (api.body="check_out");
    6: double work_hours (api.body="work_hours");
    7: string check_in_location (api.body="check_in_location");
    8: string check_out_location (api.body="check_out_location");
    9: string status (api.body="status");
    10: JobInfo job_info (api.body="job_info");
}

// 评价信息
struct ReviewInfo {
    1: i64 review_id (api.body="review_id" go.tag="json:\"review_id,string\"");
    2: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    3: i64 employer_id (api.body="employer_id" go.tag="json:\"employer_id,string\"");
    4: i64 worker_id (api.body="worker_id" go.tag="json:\"worker_id,string\"");
    5: i32 rating (api.body="rating");
    6: string content (api.body="content");
    7: string review_type (api.body="review_type");
    8: string created_at (api.body="created_at");
    9: JobInfo job_info (api.body="job_info");
    10: EmployerInfo employer_info (api.body="employer_info");
}

// 支付信息
struct PaymentInfo {
    1: i64 payment_id (api.body="payment_id" go.tag="json:\"payment_id,string\"");
    2: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    3: i64 worker_id (api.body="worker_id" go.tag="json:\"worker_id,string\"");
    4: i64 employer_id (api.body="employer_id" go.tag="json:\"employer_id,string\"");
    5: double amount (api.body="amount");
    6: string payment_method (api.body="payment_method");
    7: string status (api.body="status");
    8: string paid_at (api.body="paid_at");
    9: double platform_fee (api.body="platform_fee");
    10: JobInfo job_info (api.body="job_info");
}

// 社区帖子信息
struct CommunityPostInfo {
    1: i64 post_id (api.body="post_id" go.tag="json:\"post_id,string\"");
    2: i64 author_id (api.body="author_id" go.tag="json:\"author_id,string\"");
    3: string title (api.body="title");
    4: string content (api.body="content");
    5: string post_type (api.body="post_type");
    6: i32 view_count (api.body="view_count");
    7: i32 like_count (api.body="like_count");
    8: string status (api.body="status");
    9: string created_at (api.body="created_at");
    10: UserInfo author_info (api.body="author_info");
}

// 用户收藏信息
struct UserFavoriteJobInfo {
    1: i64 favorite_id (api.body="favorite_id" go.tag="json:\"favorite_id,string\"");
    2: i64 user_id (api.body="user_id" go.tag="json:\"user_id,string\"");
    3: i64 job_id (api.body="job_id" go.tag="json:\"job_id,string\"");
    4: string created_at (api.body="created_at");
    5: JobInfo job_info (api.body="job_info");
}
