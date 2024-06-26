package result

import (
	"github.com/gin-gonic/gin"
)

// RespMsg : 响应数据结构
type RespMsg struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewRespMsg : 生成response对象
func NewRespMsg(code ResCode, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  code.StatusText(),
		Data: data,
	}
}

func CommonResp(c *gin.Context, httpCode int, statusCode ResCode, data interface{}) {
	c.JSON(httpCode, *NewRespMsg(statusCode, data))
	c.Abort() //此路由后的 gin.HandlerFunc 将不再被调用
}

func DownloadFileResp(c *gin.Context, httpCode int, filename string, data []byte) {
	c.Status(httpCode)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename) // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(data)
}

type ListResp struct {
	Code  ResCode     `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

func RespList(c *gin.Context, httpCode int, statusCode ResCode, data interface{}, total int) {
	c.JSON(httpCode, ListResp{
		Code:  statusCode,
		Msg:   statusCode.StatusText(),
		Data:  data,
		Total: total,
	})
	c.Abort()
}
