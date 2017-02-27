package log

import (
	"testing"
	"time"
)

func TestConsoleLog(t *testing.T) {
	DefaultLogger = New(LevelDebug, 0, 0, NewTextEncoder("", "")).AddWriter(NewConsoleWriter(nil).DisableColor())
	Debug("aaa1")
	Info("aaa2")
	Warn("aaa4")
	Error("aaa3")
}

func TestFileLog(t *testing.T) {
	logger := New(LevelDebug, 0, 0, NewJSONEncoder(""))
	logger.AddWriter(NewConsoleWriter(nil))
	fw, err := NewFileWriter(logger.Level(), "logs", 1024*10, false)
	if err != nil {
		t.Fatal(err)
	}
	logger.AddWriter(fw)
	logger.Warn("DDDDDDDDDDDDDDDD")
	logger.Info("DDDDDDDDDDDDDDDD")
	logger.Debug("DDDDDDDDDDDDDDDD")
	logger.Close()
	time.Sleep(100 * time.Millisecond)
}
