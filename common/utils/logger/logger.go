package logger

import (
	"go.uber.org/zap"
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
		if logger, err = zap.NewDevelopment(); err != nil {
			return err
		}
	}

	zap.ReplaceGlobals(logger)

	return nil
}

func Global() *zap.Logger {
	return zap.L()
}
