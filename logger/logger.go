package logger

import (
	"fmt"
	"os"
	"time"
)

var LogPath string = "global.log"

// SetLogPath 设置 LogPath
func SetLogPath(path string) {
	LogPath = path
}

// Info 输出正常 Log 信息，同时打印到 log 文件中
func Info(format string, a ...interface{}) {
	fmtStr := fmt.Sprintf(format, a...)
	var t = timestampToStr(time.Now().Unix())
	infoStr := fmt.Sprintf("[I] [%s] %s\n", t, fmtStr)
	writeToLogFile(infoStr)

	fmt.Printf("%c[0;0;36m%s%c[0m", 0x1B, infoStr, 0x1B)
}

// Warn 输出警告 Log 信息，同时打印到 log 文件中
func Warn(format string, a ...interface{}) {
	fmtStr := fmt.Sprintf(format, a...)
	var t = timestampToStr(time.Now().Unix())
	infoStr := fmt.Sprintf("[W] [%s] %s\n", t, fmtStr)

	fmt.Printf("%c[0;0;33m%s%c[0m", 0x1B, infoStr, 0x1B)
}

// Error 输出错误 Log 信息，同时打印到 log 文件中
func Error(format string, a ...interface{}) {
	fmtStr := fmt.Sprintf(format, a...)
	var t = timestampToStr(time.Now().Unix())
	infoStr := fmt.Sprintf("[E] [%s] %s\n", t, fmtStr)

	fmt.Printf("%c[0;0;31m%s%c[0m", 0x1B, infoStr, 0x1B)
}

func writeToLogFile(logStr string) error {
	t := timestampToStr(time.Now().Unix())
	filePath := fmt.Sprintf("%s.%s", LogPath, t[0:10])
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	var buf = []byte(logStr)
	_, err = fd.Write(buf)
	fd.Close()

	return err
}

func timestampToStr(ts int64) string {
	var timeLayout = "2006-01-02 15:04:05"
	timeStr := time.Unix(ts, 0).Format(timeLayout)

	return timeStr
}
