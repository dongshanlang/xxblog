package logger

import (
	"errors"
	"testing"
)

func TestMain(m *testing.M) {
	c := New()
	c.SetDivision("time")                                           // 设置归档方式，"time"时间归档 "size" 文件大小归档，文件大小等可以在配置文件配置
	c.SetTimeUnit(Day)                                              // 时间归档 可以设置切割单位
	c.SetEncoding("console")                                        // 输出格式 "json" 或者 "console"
	c.SetInfoFile("/Users/lichen/data/logs/common/server_info.log") // 设置info级别日志
	c.SetCaller(true)
	c.Debug() // 是否开启debug模式
	c.InitLogger()

	Debug("this is a debug log")
	Info("this is a info log")
	Warn("this is a warn log")
	Fatal("this is a fatal log")

	Debugf("this is a debug format log :%s", "content")
	Infof("this is a info format log :%s", "content")
	Warnf("this is a warn format log :%s", "content")
	Fatalf("this is a fatal format log :%s", "content")

	Info("this is a log", With("Trace", "12345677")) // 带参数
	Debug("debug level test", With("Trace", "1111111"))
	Error("this is a log", WithError(errors.New("this is a new error"))) // 带error
}
