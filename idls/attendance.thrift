namespace go attendance

include "common.thrift"

// 获取考勤记录请求
struct GetAttendanceRecordsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: i64 job_id (api.query="job_id" go.tag="json:\"job_id,string\"");
    3: string start_date (api.query="start_date");
    4: string end_date (api.query="end_date");
}

// 获取考勤记录响应
struct GetAttendanceRecordsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.AttendanceRecordInfo> records (api.body="records");
}

// 获取考勤详情请求
struct GetAttendanceDetailReq {
    1: i32 record_id (api.path="record_id", api.vd="$>0");
}

// 获取考勤详情响应
struct GetAttendanceDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.AttendanceRecordInfo record (api.body="record");
}

// 打卡请求
struct CheckInReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
    2: string check_in_location (api.body="check_in_location", api.vd="len($)>0");
    3: double latitude (api.body="latitude");
    4: double longitude (api.body="longitude");
}

// 打卡响应
struct CheckInResp {
    1: common.BaseResp base (api.body="base");
    2: i32 record_id (api.body="record_id");
    3: string check_in_time (api.body="check_in_time");
}

// 签退请求
struct CheckOutReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
    2: string check_out_location (api.body="check_out_location", api.vd="len($)>0");
    3: double latitude (api.body="latitude");
    4: double longitude (api.body="longitude");
}

// 签退响应
struct CheckOutResp {
    1: common.BaseResp base (api.body="base");
    2: i32 record_id (api.body="record_id");
    3: string check_out_time (api.body="check_out_time");
    4: double work_hours (api.body="work_hours");
}

// 申请请假请求
struct ApplyLeaveReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
    2: string leave_date (api.body="leave_date", api.vd="len($)>0");
    3: string leave_reason (api.body="leave_reason", api.vd="len($)>0");
}

// 申请请假响应
struct ApplyLeaveResp {
    1: common.BaseResp base (api.body="base");
    2: i64 leave_id (api.body="leave_id" go.tag="json:\"leave_id,string\"");
}

// 申请补卡请求
struct ApplyMakeupReq {
    1: i64 job_id (api.body="job_id", api.vd="$>0" go.tag="json:\"job_id,string\"");
    2: string makeup_date (api.body="makeup_date", api.vd="len($)>0");
    3: string makeup_reason (api.body="makeup_reason", api.vd="len($)>0");
    4: string makeup_time (api.body="makeup_time", api.vd="len($)>0");
}

// 申请补卡响应
struct ApplyMakeupResp {
    1: common.BaseResp base (api.body="base");
    2: i64 makeup_id (api.body="makeup_id" go.tag="json:\"makeup_id,string\"");
}

service AttendanceService {
    GetAttendanceRecordsResp GetAttendanceRecords(1: GetAttendanceRecordsReq request) (api.get="/api/v1/attendance/records");
    GetAttendanceDetailResp GetAttendanceDetail(1: GetAttendanceDetailReq request) (api.get="/api/v1/attendance/records/:record_id");
    CheckInResp CheckIn(1: CheckInReq request) (api.post="/api/v1/attendance/checkin");
    CheckOutResp CheckOut(1: CheckOutReq request) (api.post="/api/v1/attendance/checkout");
    ApplyLeaveResp ApplyLeave(1: ApplyLeaveReq request) (api.post="/api/v1/attendance/apply-leave");
    ApplyMakeupResp ApplyMakeup(1: ApplyMakeupReq request) (api.post="/api/v1/attendance/apply-makeup");
}
