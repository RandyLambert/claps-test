package util

import (
	"fmt"
)

// 定义错误
type Err struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
	Errord  error  // 保存内部错误信息
}

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (
	OK = &Err{Code: 0, Message: "OK"}

	// 系统错误, 前缀为 100
	InternalServerError = &Err{Code: 10001, Message: "内部服务器错误"}
	ErrBind             = &Err{Code: 10002, Message: "请求参数错误"}
	ErrTokenSign        = &Err{Code: 10003, Message: "签名 jwt 时发生错误"}
	ErrEncrypt          = &Err{Code: 10004, Message: "加密用户密码时发生错误"}

	// 数据库错误, 前缀为 201
	ErrDatabase = &Err{Code: 20100, Message: "数据库错误"}
	ErrFill     = &Err{Code: 20101, Message: "从数据库填充 struct 时发生错误"}

	// 认证错误, 前缀是 202
	ErrValidation   = &Err{Code: 20201, Message: "验证失败"}
	ErrTokenInvalid = &Err{Code: 20202, Message: "jwt 是无效的"}

	// 用户错误, 前缀为 203
	ErrUserNotFound      = &Err{Code: 20301, Message: "用户没找到"}
	ErrPasswordIncorrect = &Err{Code: 20302, Message: "密码错误"}
)

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Errord)
}

/*
// SendSuccessResp 返回成功请求
func SendSuccessResp(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: message,
		Data:    data,
	})
}


// SendFailResp 返回失败请求
func SendFailResp(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, Response{
		Code:    -1,
		Message: message,
	})
}

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		SendSuccessResp(c, "操作成功", data)
		return
	}

	// 针对不同的错误类型进行处理
	switch errors.Cause(err).(type) {
	case *myerr.ParameterError:
		// 如果只是参数错误 返回400 并将错误信息直接返回展示
		SendFailResp(c, http.StatusBadRequest, err.Error())
	default:
		// 服务端出现未定义的异常 返回500 并打印日志
		logStackInfo(err)
		SendFailResp(c, http.StatusInternalServerError, "服务端异常")
	}

	return
}
 */