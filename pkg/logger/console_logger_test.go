package logger

import (
	"bytes"
	"log"
	"testing"
)

func TestNewConsoleLogger(t *testing.T) {
	t.Run("verbose mode on", func(t *testing.T) {
		logger := NewConsoleLogger(true)
		if logger.logLevel != TRACE {
			t.Errorf("expected loglevel %d, got %d", TRACE, logger.logLevel)
		}
	})

	t.Run("verbose mode off", func(t *testing.T) {
		logger := NewConsoleLogger(false)
		if logger.logLevel != INFO {
			t.Errorf("expected loglevel %d, got %d", INFO, logger.logLevel)
		}
	})
}

func TestConsoleLogger_log(t *testing.T) {
	var buf bytes.Buffer

	tests := []struct {
		name     string
		logFunc  func(*consoleLogger)
		logger   *consoleLogger
		expected string
	}{
		{
			name: "Info adds INFO prefix and joins arguments with spaces",
			logFunc: func(logger *consoleLogger) {
				logger.Info("hello", 123, 321)
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "INFO: hello123 321\n",
		},
		{
			name: "Infof adds INFO prefix and formats string correctly",
			logFunc: func(logger *consoleLogger) {
				logger.Infof("hello %s !", "dish")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "INFO: hello dish !\n",
		},
		{
			name: "Debug does not print if logLevel is INFO",
			logFunc: func(logger *consoleLogger) {
				logger.Debug("should not print")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  INFO,
			},
			expected: "",
		},
		{
			name: "Debug adds DEBUG prefix",
			logFunc: func(logger *consoleLogger) {
				logger.Debug("debug")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  DEBUG,
			},
			expected: "DEBUG: debug\n",
		},
		{
			name: "Debugf adds DEBUG prefix and formats string correctly",
			logFunc: func(logger *consoleLogger) {
				logger.Debugf("debug %d", 1)
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  DEBUG,
			},
			expected: "DEBUG: debug 1\n",
		},
		{
			name: "Warn prints with WARN prefix",
			logFunc: func(logger *consoleLogger) {
				logger.Warn("warn message")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "WARN: warn message\n",
		},
		{
			name: "Warnf prints formatted WARN message",
			logFunc: func(logger *consoleLogger) {
				logger.Warnf("warn %d", 42)
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "WARN: warn 42\n",
		},
		{
			name: "Error prints with ERROR prefix",
			logFunc: func(logger *consoleLogger) {
				logger.Error("error")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "ERROR: error\n",
		},
		{
			name: "Errorf prints formatted ERROR message",
			logFunc: func(logger *consoleLogger) {
				logger.Errorf("fail %s", "here")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "ERROR: fail here\n",
		},
		{
			name: "Trace prints with TRACE prefix",
			logFunc: func(logger *consoleLogger) {
				logger.Trace("trace")
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "TRACE: trace\n",
		},
		{
			name: "Tracef prints formatted TRACE message",
			logFunc: func(logger *consoleLogger) {
				logger.Tracef("trace %d", 1)
			},
			logger: &consoleLogger{
				stdLogger: log.New(&buf, "", 0),
				logLevel:  TRACE,
			},
			expected: "TRACE: trace 1\n",
		},
	}

	for _, tt := range tests {
		buf.Reset()

		tt.logFunc(tt.logger)

		output := buf.String()

		if output != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, output)
		}
	}
}

func TestConsoleLogger_log_Panic(t *testing.T) {
	logger := NewConsoleLogger(true)

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic but did not get one")
		}

		expected := "PANIC: could not start dish"
		if r != expected {
			t.Fatalf("expected panic message %s, got %s", expected, r)
		}
	}()

	logger.Panic("could not start dish")
}

func TestConsoleLogger_log_Panicf(t *testing.T) {
	logger := NewConsoleLogger(true)

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic but did not get one")
		}

		expected := "PANIC: could not start dish"
		if r != expected {
			t.Fatalf("expected panic message %s, got %s", expected, r)
		}
	}()

	logger.Panicf("could not start %s", "dish")
}
