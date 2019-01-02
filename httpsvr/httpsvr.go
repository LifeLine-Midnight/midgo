package httpsvr

import (
	"fmt"
	"net/http"
)

// MidgoSvr 结构定义
type MidgoSvr struct {
	urlCtrlMap map[string]*ctrlMethodInfo
}

// GetMidgoSvr 获取 midgo api svr 指针实例
func GetMidgoSvr() *MidgoSvr {
	return &MidgoSvr{
		urlCtrlMap: make(map[string]*ctrlMethodInfo),
	}
}

// Run server
func (svr *MidgoSvr) Run(ipAddr string, port uint32) {
	http.HandleFunc("/", svr.apiHandler)
	host := fmt.Sprintf("%s:%d", ipAddr, port)
	fmt.Printf("Server run %s\n", host)
	http.ListenAndServe(host, nil)
}
