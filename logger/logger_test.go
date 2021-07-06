package logger

import (
	"fmt"
	"testing"
)

func TestFatalLogType(t *testing.T) {
	recoverUnhandledLogs := func() {
		if r := recover(); r != nil {
			fmt.Println("Handled Fatal logs - ", r)
		}
	}
	defer recoverUnhandledLogs()
	PushLogs("Fatal Test", Level.Fatal)
}

func TestDebugLogType(t *testing.T) {
	PushLogs("Debug log Test", Level.Debug)
}
