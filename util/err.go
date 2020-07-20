package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//自定义错误类型　Err
type Err struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
	Errord  error  // 保存内部错误信息
}


//所有都是指针类型
var (
	OK = &Err{Code: 0, Message: "成功"}

	//100表示服务端错误
	ErrDataBase = 10001 //数据库出错
	ErrInternalServer = 10002 //内部错误
	ErrThirdParty = 10003 //第三方请求错误

	//200客户端错误
	ErrUnauthorized = 20001 //未登录认证
	ErrBadRequest = 20002 // 客户端没有参数
)

func NewErr(err error,code int,message string)*Err{
	return &Err{Code:code,Message:message,Errord:err}
}

//实现Error接口,把三个字段拼接起来
func (err *Err) Error() string {
	//return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Errord)
	return fmt.Sprintf("Err - code: %d, message: %s. ", err.Code, err.Message )
}

// SendSuccessResp 返回成功请求
func sendSuccessResp(c *gin.Context,data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    OK.Code,
		"message": OK.Message
		"data":    data,
	})
}


// SendFailResp 返回失败请求,失败就不返回data了
func sendFailResp(c *gin.Context, httpCode int, err *Err) {
	c.JSON(httpCode, gin.H{
		"code":    err.Code,
		"message": err.Error(),
	})
}

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
//　所有Err要求在service层封装好,返还给前端
// 400 401 501
/*
func test(ctx *gin.Context) {
	//正确请求
	HandleResponse(ctx,OK,model.User{})
	//出错请求
	HandleResponse(ctx,NewErr(errors.New("test"),ErrDataBase,"查询项目时出错"),nil)
}
 */
func HandleResponse(c *gin.Context, err *Err, data interface{} ) {
	// 如果没有错误，也就是没有Errod字段,就是正常请求
	if err.Code == 0{
		sendSuccessResp(c, data)
		return
	}

	//根据错误的不同
	switch err.Code {
	case ErrDataBase:
		sendFailResp(c,http.StatusInternalServerError,err)
	case ErrInternalServer:
		sendFailResp(c,http.StatusInternalServerError,err)
	case ErrThirdParty:
		sendFailResp(c,http.StatusInternalServerError,err)
	case ErrBadRequest:
		sendFailResp(c,http.StatusBadRequest,err)
	case ErrUnauthorized:
		sendFailResp(c,http.StatusUnauthorized,err)

	}
	//结构化输出错误
	log.WithFields(log.Fields{
		"code": err.Code,
		"err": err.Errord,
	}).Error(err.Message)

	return
}