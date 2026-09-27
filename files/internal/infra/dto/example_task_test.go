package dto

import (
	"testing"
	"time"
)

func TestExampleTaskTransitions(t *testing.T) {
	tests := []struct {
		name    string
		from    ExampleTaskStatus
		to      ExampleTaskStatus
		wantErr bool
	}{
		{name: "pending→processing", from: StatusPending, to: StatusProcessing, wantErr: false},
		{name: "processing→done", from: StatusProcessing, to: StatusDone, wantErr: false},
		{name: "processing→failed", from: StatusProcessing, to: StatusFailed, wantErr: false},
		{name: "pending→done forbidden", from: StatusPending, to: StatusDone, wantErr: true},
		{name: "done→processing forbidden", from: StatusDone, to: StatusProcessing, wantErr: true},
		{name: "same status no-op", from: StatusPending, to: StatusPending, wantErr: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			task := &ExampleTask{Status: tc.from}
			err := task.setStatus(tc.to)
			if (err != nil) != tc.wantErr {
				t.Fatalf("setStatus(%s→%s) err=%v, wantErr=%v", tc.from, tc.to, err, tc.wantErr)
			}
			if !tc.wantErr && task.Status != tc.to {
				t.Errorf("status = %s, want %s", task.Status, tc.to)
			}
		})
	}
}

func TestExampleTaskBackoffIsExponentialAndCapped(t *testing.T) {
	task := NewExampleTask("user@example.com")

	var prev time.Duration
	for i := 0; i < maxRetryCount+2; i++ {
		task.RetryCount = i
		got := task.CalcBackoff()
		if got <= 0 || got > maxRetryDelay {
			t.Fatalf("retry=%d backoff=%s out of (0, %s]", i, got, maxRetryDelay)
		}
		if i > 0 && got < prev {
			t.Errorf("retry=%d backoff=%s decreased from %s", i, got, prev)
		}
		prev = got
	}
}

func TestExampleTaskSetFailedTerminates(t *testing.T) {
	task := NewExampleTask("user@example.com")
	if err := task.SetProcessing(); err != nil {
		t.Fatalf("SetProcessing: %v", err)
	}
	for i := 0; i < maxRetryCount; i++ {
		if err := task.SetFailed(); err != nil {
			t.Fatalf("SetFailed #%d: %v", i, err)
		}
	}
	if task.Status != StatusFailed {
		t.Errorf("after %d failures status = %s, want %s", maxRetryCount, task.Status, StatusFailed)
	}
}
