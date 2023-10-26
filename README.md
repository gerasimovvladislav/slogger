# slogger

This logger is an add-on to the standard logger syllable from Golange version 1.21

Examples:
```go
package main

import (
	"log/slog"
    "github.com/gerasimovvladislav/slogger"
)

func main() {
    slog.SetDefault(logger.New("DEBUG", true))

    slog.Info("INFO")
    slog.Warn("WARN")
    slog.Error("ERROR")
    slog.Debug("DEBUG")
}
```