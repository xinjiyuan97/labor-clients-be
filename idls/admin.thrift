namespace go admin

include "common.thrift"

// ==================== 品牌方管理 ====================

// 品牌方信息
struct BrandDetail {
    1: i64 brand_id (api.body="brand_id");
    2: string company_name (api.body="company_name");
    3: string company_short_name (api.body="company_short_name");
    4: string logo (api.body="logo");
    5: string description (api.body="description");
    6: string website (api.body="website");
    7: string industry (api.body="industry");
    8: string company_size (api.body="company_size");
    9: string credit_code (api.body="credit_code");
    10: string business_license (api.body="business_license");
    11: string company_address (api.body="company_address");
    12: string business_scope (api.body="business_scope");
    13: string established_date (api.body="established_date");
    14: double registered_capital (api.body="registered_capital");
    15: string contact_person (api.body="contact_person");
    16: string contact_position (api.body="contact_position");
    17: string contact_phone (api.body="contact_phone");
    18: string contact_email (api.body="contact_email");
    19: string id_card_number (api.body="id_card_number");
    20: string id_card_front (api.body="id_card_front");
    21: string id_card_back (api.body="id_card_back");
    22: string tax_certificate (api.body="tax_certificate");
    23: string org_code_certificate (api.body="org_code_certificate");
    24: string bank_license (api.body="bank_license");
    25: string other_certificates (api.body="other_certificates");
    26: string bank_account (api.body="bank_account");
    27: string settlement_cycle (api.body="settlement_cycle");
    28: double deposit_amount (api.body="deposit_amount");
    29: string auth_status (api.body="auth_status");
    30: string account_status (api.body="account_status");
    31: i32 job_count (api.body="job_count");
    32: string activity_level (api.body="activity_level");
    33: string created_at (api.body="created_at");
    34: string updated_at (api.body="updated_at");
}

// 获取品牌方列表请求
struct GetBrandListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: string auth_status (api.query="auth_status");
    4: string account_status (api.query="account_status");
    5: string start_date (api.query="start_date");
    6: string end_date (api.query="end_date");
    7: string company_name (api.query="company_name");
    8: string sort_by (api.query="sort_by");
    9: string sort_order (api.query="sort_order");
}

// 获取品牌方列表响应
struct GetBrandListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<BrandDetail> brands (api.body="brands");
}

// 获取品牌方详情请求
struct GetBrandDetailReq {
    1: i64 brand_id (api.path="brand_id", api.vd="$>0");
}

// 获取品牌方详情响应
struct GetBrandDetailResp {
    1: common.BaseResp base (api.body="base");
    2: BrandDetail brand_info (api.body="brand_info");
}

// 创建品牌方请求
struct CreateBrandReq {
    1: string company_name (api.body="company_name", api.vd="len($)>0");
    2: string company_short_name (api.body="company_short_name");
    3: string logo (api.body="logo");
    4: string description (api.body="description");
    5: string website (api.body="website");
    6: string industry (api.body="industry");
    7: string company_size (api.body="company_size");
    8: string credit_code (api.body="credit_code");
    9: string company_address (api.body="company_address");
    10: string business_scope (api.body="business_scope");
    11: string contact_person (api.body="contact_person", api.vd="len($)>0");
    12: string contact_phone (api.body="contact_phone", api.vd="len($)>0");
    13: string contact_email (api.body="contact_email", api.vd="len($)>0");
    14: string business_license (api.body="business_license");
    15: string bank_account (api.body="bank_account");
    16: string settlement_cycle (api.body="settlement_cycle");
    17: double deposit_amount (api.body="deposit_amount");
    18: string remarks (api.body="remarks");
}

// 创建品牌方响应
struct CreateBrandResp {
    1: common.BaseResp base (api.body="base");
    2: i64 brand_id (api.body="brand_id");
}

// 更新品牌方请求
struct UpdateBrandReq {
    1: i64 brand_id (api.path="brand_id", api.vd="$>0");
    2: string company_name (api.body="company_name");
    3: string company_short_name (api.body="company_short_name");
    4: string logo (api.body="logo");
    5: string description (api.body="description");
    6: string website (api.body="website");
    7: string industry (api.body="industry");
    8: string company_size (api.body="company_size");
    9: string credit_code (api.body="credit_code");
    10: string company_address (api.body="company_address");
    11: string business_scope (api.body="business_scope");
    12: string contact_person (api.body="contact_person");
    13: string contact_phone (api.body="contact_phone");
    14: string contact_email (api.body="contact_email");
    15: string bank_account (api.body="bank_account");
    16: string settlement_cycle (api.body="settlement_cycle");
    17: double deposit_amount (api.body="deposit_amount");
    18: string account_status (api.body="account_status");
}

// 更新品牌方响应
struct UpdateBrandResp {
    1: common.BaseResp base (api.body="base");
}

// 品牌方审核请求
struct ReviewBrandReq {
    1: i64 brand_id (api.path="brand_id", api.vd="$>0");
    2: string action (api.body="action", api.vd="len($)>0"); // pass, reject, freeze, request_more
    3: string auth_level (api.body="auth_level"); // A, B, C
    4: string reason (api.body="reason");
    5: string remarks (api.body="remarks");
}

// 品牌方审核响应
struct ReviewBrandResp {
    1: common.BaseResp base (api.body="base");
}

// 批量导入品牌方请求
struct BatchImportBrandsReq {
    1: string file_url (api.body="file_url", api.vd="len($)>0");
}

// 批量导入品牌方响应
struct BatchImportBrandsResp {
    1: common.BaseResp base (api.body="base");
    2: i32 success_count (api.body="success_count");
    3: i32 fail_count (api.body="fail_count");
    4: list<string> fail_reasons (api.body="fail_reasons");
}

// ==================== 用户管理 ====================

// 品牌方用户信息
struct BrandUserInfo {
    1: i64 user_id (api.body="user_id");
    2: string username (api.body="username");
    3: string real_name (api.body="real_name");
    4: i64 brand_id (api.body="brand_id");
    5: string brand_name (api.body="brand_name");
    6: string role (api.body="role");
    7: string permissions (api.body="permissions");
    8: string created_at (api.body="created_at");
    9: string last_login_at (api.body="last_login_at");
    10: string account_status (api.body="account_status");
}

// 获取用户列表请求
struct GetUserListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: i64 brand_id (api.query="brand_id");
    4: string role (api.query="role");
    5: string account_status (api.query="account_status");
    6: string username (api.query="username");
    7: string real_name (api.query="real_name");
}

// 获取用户列表响应
struct GetUserListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<BrandUserInfo> users (api.body="users");
}

// 创建用户请求
struct CreateUserReq {
    1: string username (api.body="username", api.vd="len($)>0");
    2: string real_name (api.body="real_name", api.vd="len($)>0");
    3: i64 brand_id (api.body="brand_id", api.vd="$>0");
    4: string role (api.body="role", api.vd="len($)>0");
    5: string password (api.body="password");
    6: string email (api.body="email");
    7: string phone (api.body="phone");
    8: string permissions (api.body="permissions");
}

// 创建用户响应
struct CreateUserResp {
    1: common.BaseResp base (api.body="base");
    2: i64 user_id (api.body="user_id");
}

// 更新用户请求
struct UpdateUserReq {
    1: i64 user_id (api.path="user_id", api.vd="$>0");
    2: string real_name (api.body="real_name");
    3: string role (api.body="role");
    4: string permissions (api.body="permissions");
    5: string account_status (api.body="account_status");
    6: string email (api.body="email");
    7: string phone (api.body="phone");
}

// 更新用户响应
struct UpdateUserResp {
    1: common.BaseResp base (api.body="base");
}

// 重置密码请求
struct ResetPasswordReq {
    1: i64 user_id (api.path="user_id", api.vd="$>0");
    2: string new_password (api.body="new_password", api.vd="len($)>0");
}

// 重置密码响应
struct ResetPasswordResp {
    1: common.BaseResp base (api.body="base");
}

// ==================== 岗位管理 ====================

// 获取岗位列表请求
struct GetJobListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: i64 brand_id (api.query="brand_id");
    4: i32 category_id (api.query="category_id");
    5: double min_salary (api.query="min_salary");
    6: double max_salary (api.query="max_salary");
    7: string start_date (api.query="start_date");
    8: string end_date (api.query="end_date");
    9: string status (api.query="status");
    10: string title (api.query="title");
}

// 获取岗位列表响应
struct GetJobListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<common.JobInfo> jobs (api.body="jobs");
}

// 审核岗位请求
struct ReviewJobReq {
    1: i64 job_id (api.path="job_id", api.vd="$>0");
    2: string action (api.body="action", api.vd="len($)>0"); // pass, reject, modify, offline
    3: string reason (api.body="reason");
    4: string remarks (api.body="remarks");
}

// 审核岗位响应
struct ReviewJobResp {
    1: common.BaseResp base (api.body="base");
}

// ==================== 数据统计 ====================

// 统计数据
struct StatisticsData {
    1: i32 total_count (api.body="total_count");
    2: i32 new_count (api.body="new_count");
    3: double growth_rate (api.body="growth_rate");
    4: i32 active_count (api.body="active_count");
    5: double pass_rate (api.body="pass_rate");
}

// 趋势数据
struct TrendData {
    1: string date (api.body="date");
    2: i32 count (api.body="count");
    3: double value (api.body="value");
}

// 获取品牌方统计请求
struct GetBrandStatisticsReq {
    1: string start_date (api.query="start_date");
    2: string end_date (api.query="end_date");
    3: string period (api.query="period"); // day, week, month
}

// 获取品牌方统计响应
struct GetBrandStatisticsResp {
    1: common.BaseResp base (api.body="base");
    2: StatisticsData brand_stats (api.body="brand_stats");
    3: list<TrendData> growth_trend (api.body="growth_trend");
    4: list<TrendData> auth_trend (api.body="auth_trend");
    5: list<TrendData> activity_trend (api.body="activity_trend");
}

// 获取岗位统计请求
struct GetJobStatisticsReq {
    1: string start_date (api.query="start_date");
    2: string end_date (api.query="end_date");
    3: string period (api.query="period");
}

// 获取岗位统计响应
struct GetJobStatisticsResp {
    1: common.BaseResp base (api.body="base");
    2: StatisticsData job_stats (api.body="job_stats");
    3: list<TrendData> job_trend (api.body="job_trend");
    4: list<TrendData> application_trend (api.body="application_trend");
    5: map<string, i32> category_distribution (api.body="category_distribution");
    6: map<string, i32> salary_distribution (api.body="salary_distribution");
    7: map<string, i32> location_distribution (api.body="location_distribution");
}

// 获取用户统计请求
struct GetUserStatisticsReq {
    1: string start_date (api.query="start_date");
    2: string end_date (api.query="end_date");
    3: string period (api.query="period");
}

// 获取用户统计响应
struct GetUserStatisticsResp {
    1: common.BaseResp base (api.body="base");
    2: StatisticsData worker_stats (api.body="worker_stats");
    3: StatisticsData brand_user_stats (api.body="brand_user_stats");
    4: list<TrendData> worker_trend (api.body="worker_trend");
    5: list<TrendData> brand_user_trend (api.body="brand_user_trend");
    6: map<string, i32> location_distribution (api.body="location_distribution");
    7: map<string, i32> age_distribution (api.body="age_distribution");
    8: map<string, i32> role_distribution (api.body="role_distribution");
}

// ==================== 消息管理 ====================

// 系统通知信息
struct SystemNoticeInfo {
    1: i64 notice_id (api.body="notice_id");
    2: string title (api.body="title");
    3: string content (api.body="content");
    4: string notice_type (api.body="notice_type");
    5: string send_method (api.body="send_method");
    6: string target_users (api.body="target_users");
    7: string status (api.body="status");
    8: string created_at (api.body="created_at");
    9: string sent_at (api.body="sent_at");
    10: i32 send_count (api.body="send_count");
    11: i32 success_count (api.body="success_count");
}

// 发送系统通知请求
struct SendSystemNoticeReq {
    1: string title (api.body="title", api.vd="len($)>0");
    2: string content (api.body="content", api.vd="len($)>0");
    3: string notice_type (api.body="notice_type", api.vd="len($)>0");
    4: string send_method (api.body="send_method", api.vd="len($)>0");
    5: string target_users (api.body="target_users", api.vd="len($)>0");
    6: string send_time (api.body="send_time");
}

// 发送系统通知响应
struct SendSystemNoticeResp {
    1: common.BaseResp base (api.body="base");
    2: i64 notice_id (api.body="notice_id");
}

// 获取通知列表请求
struct GetNoticeListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: string notice_type (api.query="notice_type");
    4: string status (api.query="status");
    5: string start_date (api.query="start_date");
    6: string end_date (api.query="end_date");
}

// 获取通知列表响应
struct GetNoticeListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<SystemNoticeInfo> notices (api.body="notices");
}

// 消息模板信息
struct MessageTemplateInfo {
    1: i64 template_id (api.body="template_id");
    2: string template_name (api.body="template_name");
    3: string template_type (api.body="template_type");
    4: string subject (api.body="subject");
    5: string content (api.body="content");
    6: string variables (api.body="variables");
    7: string status (api.body="status");
    8: string created_at (api.body="created_at");
    9: string updated_at (api.body="updated_at");
}

// 创建消息模板请求
struct CreateMessageTemplateReq {
    1: string template_name (api.body="template_name", api.vd="len($)>0");
    2: string template_type (api.body="template_type", api.vd="len($)>0");
    3: string subject (api.body="subject");
    4: string content (api.body="content", api.vd="len($)>0");
    5: string variables (api.body="variables");
}

// 创建消息模板响应
struct CreateMessageTemplateResp {
    1: common.BaseResp base (api.body="base");
    2: i64 template_id (api.body="template_id");
}

// ==================== 财务管理 ====================

// 收入统计信息
struct IncomeStatsInfo {
    1: string period (api.body="period");
    2: double service_fee_income (api.body="service_fee_income");
    3: double ad_fee_income (api.body="ad_fee_income");
    4: double other_income (api.body="other_income");
    5: double total_income (api.body="total_income");
    6: double growth_rate (api.body="growth_rate");
}

// 获取收入统计请求
struct GetIncomeStatisticsReq {
    1: string start_date (api.query="start_date");
    2: string end_date (api.query="end_date");
    3: string period (api.query="period");
}

// 获取收入统计响应
struct GetIncomeStatisticsResp {
    1: common.BaseResp base (api.body="base");
    2: IncomeStatsInfo income_stats (api.body="income_stats");
    3: list<TrendData> income_trend (api.body="income_trend");
    4: map<string, double> source_analysis (api.body="source_analysis");
}

// 品牌方结算信息
struct BrandSettlementInfo {
    1: i64 settlement_id (api.body="settlement_id");
    2: i64 brand_id (api.body="brand_id");
    3: string brand_name (api.body="brand_name");
    4: double amount (api.body="amount");
    5: string settlement_cycle (api.body="settlement_cycle");
    6: string status (api.body="status");
    7: string created_at (api.body="created_at");
    8: string settled_at (api.body="settled_at");
    9: string remarks (api.body="remarks");
}

// 获取结算列表请求
struct GetSettlementListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: i64 brand_id (api.query="brand_id");
    4: string status (api.query="status");
    5: string start_date (api.query="start_date");
    6: string end_date (api.query="end_date");
}

// 获取结算列表响应
struct GetSettlementListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<BrandSettlementInfo> settlements (api.body="settlements");
}

// 处理结算请求
struct ProcessSettlementReq {
    1: i64 settlement_id (api.path="settlement_id", api.vd="$>0");
    2: string action (api.body="action", api.vd="len($)>0"); // approve, reject, complete
    3: string remarks (api.body="remarks");
}

// 处理结算响应
struct ProcessSettlementResp {
    1: common.BaseResp base (api.body="base");
}

// ==================== 系统设置 ====================

// 系统配置信息
struct SystemConfigInfo {
    1: string config_key (api.body="config_key");
    2: string config_value (api.body="config_value");
    3: string config_type (api.body="config_type");
    4: string description (api.body="description");
    5: string updated_at (api.body="updated_at");
}

// 获取系统配置请求
struct GetSystemConfigReq {
    1: string config_key (api.query="config_key");
}

// 获取系统配置响应
struct GetSystemConfigResp {
    1: common.BaseResp base (api.body="base");
    2: SystemConfigInfo config_info (api.body="config_info");
}

// 更新系统配置请求
struct UpdateSystemConfigReq {
    1: string config_key (api.path="config_key", api.vd="len($)>0");
    2: string config_value (api.body="config_value", api.vd="len($)>0");
    3: string description (api.body="description");
}

// 更新系统配置响应
struct UpdateSystemConfigResp {
    1: common.BaseResp base (api.body="base");
}

// 管理员信息
struct AdminInfo {
    1: i64 admin_id (api.body="admin_id");
    2: string username (api.body="username");
    3: string real_name (api.body="real_name");
    4: string role (api.body="role");
    5: string permissions (api.body="permissions");
    6: string created_at (api.body="created_at");
    7: string last_login_at (api.body="last_login_at");
    8: string account_status (api.body="account_status");
}

// 创建管理员请求
struct CreateAdminReq {
    1: string username (api.body="username", api.vd="len($)>0");
    2: string real_name (api.body="real_name", api.vd="len($)>0");
    3: string password (api.body="password", api.vd="len($)>0");
    4: string role (api.body="role", api.vd="len($)>0");
    5: string permissions (api.body="permissions");
}

// 创建管理员响应
struct CreateAdminResp {
    1: common.BaseResp base (api.body="base");
    2: i64 admin_id (api.body="admin_id");
}

// 获取管理员列表请求
struct GetAdminListReq {
    1: i32 page (api.query="page", api.vd="$>=1");
    2: i32 limit (api.query="limit", api.vd="$>=1&&$<=100");
    3: string role (api.query="role");
    4: string account_status (api.query="account_status");
}

// 获取管理员列表响应
struct GetAdminListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_info (api.body="page_info");
    3: list<AdminInfo> admins (api.body="admins");
}

// ==================== 服务定义 ====================

service AdminService {
    // 品牌方管理
    GetBrandListResp GetBrandList(1: GetBrandListReq request) (api.get="/api/v1/admin/brands");
    GetBrandDetailResp GetBrandDetail(1: GetBrandDetailReq request) (api.get="/api/v1/admin/brands/:brand_id");
    CreateBrandResp CreateBrand(1: CreateBrandReq request) (api.post="/api/v1/admin/brands");
    UpdateBrandResp UpdateBrand(1: UpdateBrandReq request) (api.put="/api/v1/admin/brands/:brand_id");
    ReviewBrandResp ReviewBrand(1: ReviewBrandReq request) (api.post="/api/v1/admin/brands/:brand_id/review");
    BatchImportBrandsResp BatchImportBrands(1: BatchImportBrandsReq request) (api.post="/api/v1/admin/brands/batch-import");

    // 用户管理
    GetUserListResp GetUserList(1: GetUserListReq request) (api.get="/api/v1/admin/users");
    CreateUserResp CreateUser(1: CreateUserReq request) (api.post="/api/v1/admin/users");
    UpdateUserResp UpdateUser(1: UpdateUserReq request) (api.put="/api/v1/admin/users/:user_id");
    ResetPasswordResp ResetPassword(1: ResetPasswordReq request) (api.post="/api/v1/admin/users/:user_id/reset-password");

    // 岗位管理
    GetJobListResp GetJobList(1: GetJobListReq request) (api.get="/api/v1/admin/jobs");
    ReviewJobResp ReviewJob(1: ReviewJobReq request) (api.post="/api/v1/admin/jobs/:job_id/review");

    // 数据统计
    GetBrandStatisticsResp GetBrandStatistics(1: GetBrandStatisticsReq request) (api.get="/api/v1/admin/statistics/brands");
    GetJobStatisticsResp GetJobStatistics(1: GetJobStatisticsReq request) (api.get="/api/v1/admin/statistics/jobs");
    GetUserStatisticsResp GetUserStatistics(1: GetUserStatisticsReq request) (api.get="/api/v1/admin/statistics/users");

    // 消息管理
    SendSystemNoticeResp SendSystemNotice(1: SendSystemNoticeReq request) (api.post="/api/v1/admin/notices");
    GetNoticeListResp GetNoticeList(1: GetNoticeListReq request) (api.get="/api/v1/admin/notices");
    CreateMessageTemplateResp CreateMessageTemplate(1: CreateMessageTemplateReq request) (api.post="/api/v1/admin/message-templates");

    // 财务管理
    GetIncomeStatisticsResp GetIncomeStatistics(1: GetIncomeStatisticsReq request) (api.get="/api/v1/admin/finance/income");
    GetSettlementListResp GetSettlementList(1: GetSettlementListReq request) (api.get="/api/v1/admin/finance/settlements");
    ProcessSettlementResp ProcessSettlement(1: ProcessSettlementReq request) (api.post="/api/v1/admin/finance/settlements/:settlement_id/process");

    // 系统设置
    GetSystemConfigResp GetSystemConfig(1: GetSystemConfigReq request) (api.get="/api/v1/admin/config/:config_key");
    UpdateSystemConfigResp UpdateSystemConfig(1: UpdateSystemConfigReq request) (api.put="/api/v1/admin/config/:config_key");
    CreateAdminResp CreateAdmin(1: CreateAdminReq request) (api.post="/api/v1/admin/admins");
    GetAdminListResp GetAdminList(1: GetAdminListReq request) (api.get="/api/v1/admin/admins");
}
