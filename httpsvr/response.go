package httpsvr

// Resp 返回的结构体
// 会转为 json 返回
type Resp struct {
	Rtn  int         `json:"rtn"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SetRtn 设置应用层返回码，非 http 状态码
func (ret *Resp) SetRtn(i int) {
	ret.Rtn = i
}

// SetMsg 设置提示消息
func (ret *Resp) SetMsg(msg string) {
	ret.Msg = msg
}

// SetData 设置返回数据
func (ret *Resp) SetData(data interface{}) {
	ret.Data = data
}
