package errno

import (
	"fmt"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func GetBaseResp(err error) *common.BaseResp {
	if err == nil {
		return &common.BaseResp{
			Code:    200,
			Message: "success",
		}
	}

	if e, ok := err.(*Error); ok {
		return &common.BaseResp{
			Code:    int32(e.Code),
			Message: e.Message,
		}
	} else {
		return &common.BaseResp{
			Code:    500,
			Message: "系统错误",
		}
	}
}
