package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "服务器错误",
	INVALID_PARAMS: "参数错误",
	EXIST:          "该记录已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
