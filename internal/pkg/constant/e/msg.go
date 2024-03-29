package e

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "操作成功",   // 200
	CodeInvalidParams: "请求参数错误", // 400
	CodeNotFound:      "记录找不到",  // 404
	CodeInternalError: "服务器繁忙",  // 500

	CodeConvDataErr:       "数据转换错误",   // 10000
	CodeValidateParamsErr: "参数校验错误",   // 10001
	CodeInvalidToken:      "无效的token", // 10002
	CodeNeedLogin:         "请先登陆",     // 10003
	CodeInvalidID:         "无效的ID",    // 10004

	CodeUserHasExist:   "该用户已存在",        // 20002
	CodeEmailExist:     "该邮箱已存在",        // 20003
	CodeWrongPassword:  "密码错误",          // 20004
	CodeUserHasDeleted: "用户已经删除,请勿重复操作", // 20004
}

// Msg .
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeInternalError]
	}
	return msg
}
