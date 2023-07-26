package logger

import (
	"context"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	cfg := Config{
		LoggingLevel:          DebugLevel,
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    true,
		Directory:             "log",
		CallerEnabled:         true,
		CallerSkip:            1,
		Filename:              "",
		MaxSize:               100,
		MaxBackups:            10,
	}
	Configure(cfg)

	Debug("test debug")
	Info("test info")
	Warn("test warn")
	Error("test error")
}

func TestTraffic(t *testing.T) {
	cfg := TrafficLogConfig{
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    true,
		LoggingDirectory:      "log",
		Filename:              "data.log",
		MaxSize:               100,
		MaxBackups:            10,
	}
	ConfigureTrafficLog(cfg)

	ctx := context.Background()
	TrafficLoggerFromContext(ctx).DataWith(&Traffic{
		Typ:  TrafficTypRequest,
		Cmd:  "echo",
		Code: 0,
		Msg:  "ok",
		Cost: 2 * time.Millisecond,
		Req:  "Alice",
		Resp: map[string]any{
			"name": "Alice",
			"age":  "18",
		},
	}, Fields{
		"foo": "bar",
	})

	time.Sleep(1 * time.Second)
}
