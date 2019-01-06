package httpsvr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type ctrlMethodInfo struct {
	ctrlStructType reflect.Type
	ctrlPtrType    reflect.Type
	ctrlName       string
	methodName     string
	methodIdx      int
}

// AddAPIController 给 svr 实例添加 controller
// 利用反射自动获取其中的函数
// 例如 TestController 的 GetData 方法
// 会被映射为 {preURL}/test/getdata
func (svr *MidgoSvr) AddController(preURL string, ctrl interface{}) {
	// 标准化 preURL 到 /xxx/xxx 这样
	preURL = fmtPreURL(preURL)
	t := reflect.TypeOf(ctrl)
	var ctrlStructType reflect.Type
	var ctrlPtrType reflect.Type

	if t.Kind() == reflect.Ptr {
		ctrlPtrType = t
		ctrlStructType = reflect.ValueOf(ctrl).Elem().Type()
	} else {
		ctrlStructType = t
		ctrlPtrType = reflect.PtrTo(t)
	}

	ctrlName := strings.ToLower(ctrlStructType.Name())
	if strings.HasSuffix(ctrlName, "controller") {
		ctrlName = ctrlName[0 : len(ctrlName)-len("controller")]
	}

	reqType := reflect.TypeOf((*Req)(nil))
	resType := reflect.TypeOf((*Resp)(nil))

	for i := 0; i < ctrlPtrType.NumMethod(); i++ {
		method := ctrlPtrType.Method(i)
		methodName := strings.ToLower(method.Name)

		url := fmt.Sprintf("%s/%s/%s", preURL, ctrlName, methodName)
		fmt.Println(url)

		if method.Type.NumIn() == 2 && method.Type.NumOut() == 1 &&
			method.Type.In(1) == reqType && method.Type.Out(0) == resType {
			svr.urlCtrlMap[url] = &ctrlMethodInfo{
				ctrlStructType: ctrlStructType,
				ctrlPtrType:    ctrlPtrType,
				ctrlName:       ctrlName,
				methodName:     methodName,
				methodIdx:      i,
			}
		}
	}
}

func (svr *MidgoSvr) apiHandler(w http.ResponseWriter, r *http.Request) {
	ctrlMethodInfo, ok := svr.urlCtrlMap[r.URL.Path]
	fmt.Println(r.URL.Path)

	if !ok {
		w.Write([]byte(NotFoundMsg))
	} else {
		ctrlObj := reflect.New(ctrlMethodInfo.ctrlStructType)
		inputs := []reflect.Value{reflect.ValueOf(r)}

		rets := ctrlObj.Method(ctrlMethodInfo.methodIdx).Call(inputs)
		ret := rets[0].Interface().(*Resp)

		jsonData, err := json.Marshal(ret)
		if err != nil {
			w.Write([]byte(UnknowErrMsg))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(jsonData)
	}
}

// "/"" -> ""
// "url" -> "/url"
// "url/" -> "/url"
// "/url" -> "/url"
func fmtPreURL(preURL string) string {
	if len(preURL) == 0 || (len(preURL) == 1 && preURL[0] == '/') {
		return ""
	}

	var fmtURL string
	if preURL[0] != '/' {
		fmtURL = fmt.Sprintf("/%s", preURL)
	} else {
		fmtURL = preURL
	}

	if fmtURL[len(fmtURL)-1] == '/' {
		fmtURL = fmtURL[0 : len(fmtURL)-1]
	}

	return fmtURL
}
