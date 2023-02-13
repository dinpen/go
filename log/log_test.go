package log

import "testing"

func TestLog(t *testing.T) {
	Debug("debug message")
	DebugF("debug message %s", "arg")
	DebugJ("debug message")

	Info("info message")
	InfoF("info message %s", "arg")
	InfoJ("info message")

	Warn("warn message")
	WarnF("warn message %s", "arg")
	WarnJ("warn message")

	Error("error message")
	ErrorF("error message %s", "arg")
	ErrorJ("error message")

	// Fatal("fatal message")
	// FatalF("fatal message %s", "arg")
	// FatalJ("fatal message")

	Print("some message")
	Printf("some message %s", "arg")
}
