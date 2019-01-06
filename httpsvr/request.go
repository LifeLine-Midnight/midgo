package httpsvr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Req http.Request 的别名
type Req = http.Request

// JsonBodyDecode 把 r.Body 中的  json 数据扔到 v 中，v 需要是一个 Ptr
func JsonBodyDecode(r *Req, v interface{}) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(reqBody, v)
	if err != nil {
		return err
	}
	return nil
}
