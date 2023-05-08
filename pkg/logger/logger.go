package logger

import (
	"github.com/vsPEach/Kursovaya/config"
	"go.uber.org/zap"
)

func New(conf config.LoggerConf) *zap.SugaredLogger {
	logConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(conf.Level),
		DisableCaller:    true,
		Development:      true,
		Encoding:         conf.Encoding,
		OutputPaths:      conf.OutputPath,
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}
	return zap.Must(logConfig.Build()).Sugar()
}
