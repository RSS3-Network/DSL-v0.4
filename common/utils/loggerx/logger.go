package loggerx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

func Initialize(mode string) (err error) {
	var logger *zap.Logger

	switch mode {
	case ModeProduction:
		if logger, err = zap.NewProduction(); err != nil {
			return err
		}
	case ModeDevelopment:
		fallthrough
	default:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		if logger, err = config.Build(); err != nil {
			return err
		}
	}

	zap.ReplaceGlobals(logger)

	return nil
}

func Global() *zap.Logger {
	return zap.L()
}
