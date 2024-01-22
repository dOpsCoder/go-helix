package util

import (
	"io"
	"log/slog"
	"sync"
)

var once sync.Once
var logger *slog.Logger

func NopLogger() *slog.Logger {

	once.Do(func() {
		if logger == nil {
			logger = slog.New(slog.NewTextHandler(io.Discard, nil))
		}
	})
	return logger
}
