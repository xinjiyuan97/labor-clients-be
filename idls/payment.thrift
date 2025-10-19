namespace go payment

include "common.thrift"

// 获取支付记录请求
struct GetPaymentRecordsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
    4: string status (api.query="status");
}

// 获取支付记录响应
struct GetPaymentRecordsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.PaymentInfo> payments (api.body="payments");
}

// 获取支付详情请求
struct GetPaymentDetailReq {
    1: i64 payment_id (api.path="payment_id", api.vd="$>0");
}

// 获取支付详情响应
struct GetPaymentDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.PaymentInfo payment (api.body="payment");
}

// 申请提现请求
struct ApplyWithdrawReq {
    1: double amount (api.body="amount", api.vd="$>0");
    2: string payment_method (api.body="payment_method", api.vd="len($)>0");
    3: string account_info (api.body="account_info", api.vd="len($)>0");
}

// 申请提现响应
struct ApplyWithdrawResp {
    1: common.BaseResp base (api.body="base");
    2: i64 withdraw_id (api.body="withdraw_id");
}

// 获取提现记录请求
struct GetWithdrawRecordsReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string start_date (api.query="start_date");
    3: string end_date (api.query="end_date");
    4: string status (api.query="status");
}

// 获取提现记录响应
struct GetWithdrawRecordsResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.PaymentInfo> withdraws (api.body="withdraws");
}

service PaymentService {
    GetPaymentRecordsResp GetPaymentRecords(1: GetPaymentRecordsReq request) (api.get="/api/v1/payments");
    GetPaymentDetailResp GetPaymentDetail(1: GetPaymentDetailReq request) (api.get="/api/v1/payments/:payment_id");
    ApplyWithdrawResp ApplyWithdraw(1: ApplyWithdrawReq request) (api.post="/api/v1/payments/withdraw");
    GetWithdrawRecordsResp GetWithdrawRecords(1: GetWithdrawRecordsReq request) (api.get="/api/v1/payments/withdraw-records");
}
