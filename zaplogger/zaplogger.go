package zaplogger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

var ZapLog *zapLog = &zapLog{zap.L()}

type zapLog struct {
	*zap.Logger
}

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	ZapLog = &zapLog{logger}
	zap.ReplaceGlobals(logger)
}

func (zaplog *zapLog) Infof(format string, message ...interface{}) {
	zaplog.Info(fmt.Sprintf(format, message...))
}

func (zaplog *zapLog) Warnf(format string, message ...interface{}) {
	zaplog.Warn(fmt.Sprintf(format, message...))
}

func (zaplog *zapLog) Errorf(format string, message ...interface{}) {
	zaplog.Error(fmt.Sprintf(format, message...))
}
