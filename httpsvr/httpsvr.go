package httpsvr

import (
	"fmt"
	"net/http"

	"midgo/logger"
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
	logger.Info("Server run at %s", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		logger.Error("err: %s", err.Error())
	}
}
