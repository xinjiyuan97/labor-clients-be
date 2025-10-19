namespace go main

include "common.thrift"
include "auth.thrift"
include "schedule.thrift"
include "job.thrift"
include "job_application.thrift"
include "message.thrift"
include "user.thrift"
include "community.thrift"
include "attendance.thrift"
include "review.thrift"
include "payment.thrift"
include "system.thrift"
include "upload.thrift"

// 主业务的 ping 接口
struct PingReq {
}

struct PingResp {
    1: common.BaseResp base (api.body="base");
    2: string message (api.body="message");
}

service MainService {
    PingResp Ping(1: PingReq request) (api.get="/ping");
}
