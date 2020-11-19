package log

import "testing"

func TestInit(t *testing.T) {
	zf := ZapConfig{
		//LogFile:    "log.log", // 不传就不写文件
		LogLevel:   "debug",
		MaxAge:     1,
		MaxSize:    1,
		MaxBackups: 1,
		Compress:   true,
		JsonFormat: false,
	}
	Init(zf)
	Sugar.Info("log log", "success", true, 1)
	Sugar.Infof("log log success %t %d", true, 1)
	Sugar.Infow("log log", "success", true)
}
