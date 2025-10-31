namespace go message

include "common.thrift"

// 获取消息列表请求
struct GetMessageListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string msg_category (api.query="msg_category");
    3: bool is_read (api.query="is_read");
}

// 获取消息列表响应
struct GetMessageListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.MessageInfo> messages (api.body="messages");
}

// 获取消息详情请求
struct GetMessageDetailReq {
    1: i64 message_id (api.path="message_id", api.vd="$>0" go.tag="json:\"message_id,string\"");
}

// 获取消息详情响应
struct GetMessageDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.MessageInfo message (api.body="message");
}

// 发送消息请求
struct SendMessageReq {
    1: i64 to_user (api.body="to_user", api.vd="$>0" go.tag="json:\"to_user,string\"");
    2: string message_type (api.body="message_type", api.vd="len($)>0");
    3: string content (api.body="content", api.vd="len($)>0");
    4: string msg_category (api.body="msg_category");
}

// 发送消息响应
struct SendMessageResp {
    1: common.BaseResp base (api.body="base");
    2: i64 message_id (api.body="message_id" go.tag="json:\"message_id,string\"");
}

// 标记消息已读请求
struct MarkMessageReadReq {
    1: i64 message_id (api.path="message_id", api.vd="$>0" go.tag="json:\"message_id,string\"");
}

// 标记消息已读响应
struct MarkMessageReadResp {
    1: common.BaseResp base (api.body="base");
}

// 批量标记已读请求
struct BatchMarkReadReq {
    1: list<i64> message_ids (api.body="message_ids", api.vd="len($)>0" go.tag="json:\"message_ids\"");
}

// 批量标记已读响应
struct BatchMarkReadResp {
    1: common.BaseResp base (api.body="base");
    2: i32 count (api.body="count");
}

// 获取未读消息数请求
struct GetUnreadCountReq {
    1: string msg_category (api.query="msg_category");
}

// 获取未读消息数响应
struct GetUnreadCountResp {
    1: common.BaseResp base (api.body="base");
    2: i32 unread_count (api.body="unread_count");
}

service MessageService {
    GetMessageListResp GetMessageList(1: GetMessageListReq request) (api.get="/api/v1/messages");
    GetMessageDetailResp GetMessageDetail(1: GetMessageDetailReq request) (api.get="/api/v1/messages/:message_id");
    SendMessageResp SendMessage(1: SendMessageReq request) (api.post="/api/v1/messages");
    MarkMessageReadResp MarkMessageRead(1: MarkMessageReadReq request) (api.put="/api/v1/messages/:message_id/read");
    BatchMarkReadResp BatchMarkRead(1: BatchMarkReadReq request) (api.put="/api/v1/messages/batch-read");
    GetUnreadCountResp GetUnreadCount(1: GetUnreadCountReq request) (api.get="/api/v1/messages/unread-count");
}
