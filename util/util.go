package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"spider-golang-web/global"
	"strings"
)

type Response struct {
	Code     ResCode     `json:"code"`
	Msg      interface{} `json:"msg"`
	Data     interface{} `json:"data"`
	Language string      `json:"language"`
}

func RspData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:     CodeSuccess,
		Msg:      Msg(CodeSuccess),
		Data:     data,
		Language: "golang",
	})
}

func RspError(c *gin.Context, code ResCode, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code:     code,
			Msg:      Msg(code),
			Data:     nil,
			Language: "golang",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:     code,
			Msg:      err.Error(),
			Data:     nil,
			Language: "golang",
		})
	}
}

func RspBindingError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		errMap := errs.Translate(global.Trans)
		graceMap := make(map[string]string)
		for k, v := range errMap {
			graceMap[k[strings.Index(k, ".")+1:]] = v
		}
		c.JSON(http.StatusOK, Response{
			Code: CodeInvalidParam,
			Msg:  graceMap,
			Data: nil,
		})
	} else {
		RspError(c, CodeServeBusy, err)
	}
}
