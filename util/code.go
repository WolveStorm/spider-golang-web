package util

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeServeBusy
)

var codeMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeInvalidParam: "请求参数异常",
	CodeServeBusy:    "服务器繁忙",
}

func Msg(code ResCode) string {
	s, ok := codeMap[code]
	if !ok {
		return codeMap[CodeServeBusy]
	} else {
		return s
	}
}
