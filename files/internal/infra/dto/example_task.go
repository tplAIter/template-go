package dto

import (
	"fmt"
	"time"
)

type ExampleTaskStatus string

const (
	StatusPending    ExampleTaskStatus = "pending"
	StatusProcessing ExampleTaskStatus = "processing"
	StatusDone       ExampleTaskStatus = "done"
	StatusFailed     ExampleTaskStatus = "failed"
)
const (
	maxRetryCount = 5
	maxRetryDelay = 30 * time.Second
)

type ExampleTask struct {
	Email      string
	Status     ExampleTaskStatus
	RetryCount int
}

func NewExampleTask(email string) *ExampleTask {
	return &ExampleTask{Email: email, Status: StatusPending}
}
func (t *ExampleTask) setStatus(next ExampleTaskStatus) error {
	if next == t.Status {
		return nil
	}
	valid := t.Status == StatusPending && next == StatusProcessing || t.Status == StatusProcessing && (next == StatusDone || next == StatusFailed)
	if !valid {
		return fmt.Errorf("invalid task transition %s -> %s", t.Status, next)
	}
	t.Status = next
	return nil
}
func (t *ExampleTask) SetProcessing() error { return t.setStatus(StatusProcessing) }
func (t *ExampleTask) SetFailed() error {
	if t.Status == StatusFailed {
		return nil
	}
	return t.setStatus(StatusFailed)
}
func (t *ExampleTask) CalcBackoff() time.Duration {
	d := time.Second << min(t.RetryCount, 5)
	if d > maxRetryDelay {
		return maxRetryDelay
	}
	return d
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
