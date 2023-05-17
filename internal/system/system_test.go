package system_test

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/dddsphere/martello/internal/config"
	mlog "github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
)

const (
	name   = "test"
	notify = true
)

var (
	cfg = &config.Config{}
	log = mlog.NewLogger(mlog.LogLevel.Error, false)
)

// Dummy tasks for testing
func dummyTask(ctx context.Context) error {
	return nil
}

func failingTask(ctx context.Context) error {
	return errors.New("failed task")
}

// Dummy teardown function for testing
func dummyTeardown() {
	// do nothing
}

func TestSupervisor(t *testing.T) {
	// Create a supervisor instance
	sv := system.NewSupervisor(name, cfg, log, notify)

	// AddTask tasks and teardown functions
	sv.AddTasks(dummyTask, failingTask)
	sv.AddShutdownTasks(dummyTeardown)

	// Create a context with timeout for testing
	ctx, cancel := context.WithTimeout(sv.Context(), 2*time.Second)
	defer cancel()

	// Run the supervisor
	err := sv.Wait()

	// Check if the supervisor returned an error
	if err == nil {
		t.Error("expected error, but got nil")
	}

	// Check if the supervisor canceled the context
	if ctx.Err() != context.Canceled {
		t.Errorf("expected context cancellation, but got %v", ctx.Err())
	}
}

func TestSupervisorWithSignal(t *testing.T) {
	// Create a supervisor instance with signal notification
	sv := system.NewSupervisor(name, cfg, log, notify)

	// Capture the os.Interrupt signal to simulate termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// AddTask a dummy task
	sv.AddTasks(dummyTask)

	// Run the supervisor in a separate goroutine
	go func() {
		err := sv.Wait()
		if err != nil {
			t.Errorf("supervisor returned an error: %v", err)
		}
	}()

	// Simulate receiving the os.Interrupt signal
	c <- os.Interrupt

	// Allow some time for the supervisor to handle the signal
	time.Sleep(100 * time.Millisecond)

	// Check if the supervisor canceled the context
	if sv.Context().Err() != context.Canceled {
		t.Errorf("expected context cancellation, but got %v", sv.Context().Err())
	}

	// Cleanup
	signal.Stop(c)
	close(c)
}

func TestSupervisorWithContextCancel(t *testing.T) {
	// Create a parent context with a timeout
	parentCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create a supervisor instance with the parent context
	sv := system.NewSupervisor(name, cfg, log, notify)

	// AddTask a dummy task
	sv.AddTasks(dummyTask)

	// Run the supervisor
	err := sv.Wait()

	// Check if the supervisor returned an error
	if err != nil {
		t.Errorf("supervisor returned an error: %v", err)
	}

	// Check if the parent context is canceled
	if parentCtx.Err() != context.Canceled {
		t.Errorf("expected parent context cancellation, but got %v", parentCtx.Err())
	}
}
