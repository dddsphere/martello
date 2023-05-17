package log_test

import (
	"bytes"
	"io"
	stdlog "log"
	"os"
	"testing"

	"github.com/dddsphere/martello/internal/log"
)

// MockLogger is a mock implementation of the Logger interface for testing.
type MockLogger struct {
	debug *bytes.Buffer
	info  *bytes.Buffer
	error *bytes.Buffer
}

func TestSimpleLoggerDebug(t *testing.T) {
	output := NewMockOutput()

	sl := log.NewLogger("debug", false)
	sl.SetDebugOutput(output.debug)

	sl.Debug("debug message")
	expectedOutput := "debug message\n"
	actualOutput := output.debug.String()
	if actualOutput != expectedOutput {
		t.Errorf("Expected debug output:\n%s\nBut got:\n%s", expectedOutput, actualOutput)
	}
}

func TestSimpleLoggerInfo(t *testing.T) {
	output := NewMockOutput()

	sl := log.NewLogger("info", false)
	sl.SetInfoOutput(output.info)

	sl.Info("info message")
	expectedOutput := "info message\n"
	actualOutput := output.info.String()
	if actualOutput != expectedOutput {
		t.Errorf("Expected info output:\n%s\nBut got:\n%s", expectedOutput, actualOutput)
	}
}

func TestSimpleLoggerError(t *testing.T) {
	output := NewMockOutput()

	sl := log.NewLogger("error", false)
	sl.SetInfoOutput(output.error)

	sl.Info("error message")
	expectedOutput := "error message\n"
	actualOutput := output.error.String()
	if actualOutput != expectedOutput {
		t.Errorf("Expected info output:\n%s\nBut got:\n%s", expectedOutput, actualOutput)
	}
}

func TestMain(m *testing.M) {
	// Redirect log output to io.Discard during tests
	stdlog.SetOutput(io.Discard)

	// Run the tests
	exitCode := m.Run()

	// Clean up any resources if required

	os.Exit(exitCode)
}

func NewMockOutput() *MockLogger {
	return &MockLogger{
		debug: &bytes.Buffer{},
		info:  &bytes.Buffer{},
		error: &bytes.Buffer{},
	}
}
