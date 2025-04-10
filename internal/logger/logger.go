package logger

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
)

func SugarLogger(options ...Option) *LoggerProxy {
	var logger = logger.Sugar().With("release", runtime.Release, "env", runtime.GetEnvironment())

	for _, option := range options {
		logger = option(logger)
	}

	return &LoggerProxy{logger: logger}
}
