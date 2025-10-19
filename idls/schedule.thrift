namespace go schedule

include "common.thrift"

// 创建日程请求
struct CreateScheduleReq {
    1: string title (api.body="title", api.vd="len($)>0");
    2: i64 job_id (api.body="job_id");
    3: string start_time (api.body="start_time", api.vd="len($)>0");
    4: string end_time (api.body="end_time", api.vd="len($)>0");
    5: string location (api.body="location");
    6: string notes (api.body="notes");
    7: i32 reminder_minutes (api.body="reminder_minutes");
}

// 创建日程响应
struct CreateScheduleResp {
    1: common.BaseResp base (api.body="base");
    2: i64 schedule_id (api.body="schedule_id");
    3: string title (api.body="title");
    4: string status (api.body="status");
}

// 获取日程列表请求
struct GetScheduleListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
    4: string status (api.query="status");
}

// 获取日程列表响应
struct GetScheduleListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.ScheduleInfo> schedules (api.body="schedules");
}

// 获取今日排班请求
struct GetTodayScheduleReq {
    1: string date (api.query="date");
}

// 获取今日排班响应
struct GetTodayScheduleResp {
    1: common.BaseResp base (api.body="base");
    2: list<common.ScheduleInfo> schedules (api.body="schedules");
}

// 获取日程详情请求
struct GetScheduleDetailReq {
    1: i64 schedule_id (api.path="schedule_id", api.vd="$>0");
}

// 获取日程详情响应
struct GetScheduleDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.ScheduleInfo schedule (api.body="schedule");
}

// 更新日程请求
struct UpdateScheduleReq {
    1: i64 schedule_id (api.path="schedule_id", api.vd="$>0");
    2: string title (api.body="title");
    3: string start_time (api.body="start_time");
    4: string end_time (api.body="end_time");
    5: string location (api.body="location");
    6: string notes (api.body="notes");
    7: i32 reminder_minutes (api.body="reminder_minutes");
}

// 更新日程响应
struct UpdateScheduleResp {
    1: common.BaseResp base (api.body="base");
    2: common.ScheduleInfo schedule (api.body="schedule");
}

// 更新日程状态请求
struct UpdateScheduleStatusReq {
    1: i64 schedule_id (api.path="schedule_id", api.vd="$>0");
    2: string status (api.body="status", api.vd="len($)>0");
}

// 更新日程状态响应
struct UpdateScheduleStatusResp {
    1: common.BaseResp base (api.body="base");
    2: string status (api.body="status");
}

// 删除日程请求
struct DeleteScheduleReq {
    1: i64 schedule_id (api.path="schedule_id", api.vd="$>0");
}

// 删除日程响应
struct DeleteScheduleResp {
    1: common.BaseResp base (api.body="base");
}

service ScheduleService {
    CreateScheduleResp CreateSchedule(1: CreateScheduleReq request) (api.post="/api/v1/schedules");
    GetScheduleListResp GetScheduleList(1: GetScheduleListReq request) (api.get="/api/v1/schedules");
    GetTodayScheduleResp GetTodaySchedule(1: GetTodayScheduleReq request) (api.get="/api/v1/schedules/today");
    GetScheduleDetailResp GetScheduleDetail(1: GetScheduleDetailReq request) (api.get="/api/v1/schedules/:schedule_id");
    UpdateScheduleResp UpdateSchedule(1: UpdateScheduleReq request) (api.put="/api/v1/schedules/:schedule_id");
    UpdateScheduleStatusResp UpdateScheduleStatus(1: UpdateScheduleStatusReq request) (api.put="/api/v1/schedules/:schedule_id/status");
    DeleteScheduleResp DeleteSchedule(1: DeleteScheduleReq request) (api.delete="/api/v1/schedules/:schedule_id");
}
