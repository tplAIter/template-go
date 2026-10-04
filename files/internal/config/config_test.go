package config

import (
	"testing"
	"time"
)

func TestEnvironmentConfiguration(t *testing.T) {
	t.Setenv("APP_LISTEN_ADDR", "127.0.0.1:9191")
	t.Setenv("TEMPORAL_HOSTPORT", "temporal:7233")
	t.Setenv("TEMPORAL_NAMESPACE", "local")
	t.Setenv("TEMPORAL_TASK_QUEUE", "local-queue")
	t.Setenv("TEMPORAL_DIAL_TIMEOUT", "3s")
	t.Setenv("TEMPORAL_WORKER_STOP_TIMEOUT", "4s")
	cfg, err := Load()
	if err != nil { t.Fatal(err) }
	if cfg.ListenAddr != "127.0.0.1:9191" || cfg.Temporal.HostPort != "temporal:7233" ||
		cfg.Temporal.Namespace != "local" || cfg.Temporal.TaskQueue != "local-queue" ||
		cfg.Temporal.DialTimeout != 3*time.Second || cfg.Temporal.WorkerStopTimeout != 4*time.Second {
		t.Fatalf("environment not applied: %+v", cfg)
	}
}

func TestRejectUnboundedTimeout(t *testing.T) {
	for _, key := range []string{"APP_READ_TIMEOUT", "APP_WRITE_TIMEOUT", "APP_IDLE_TIMEOUT", "APP_SHUTDOWN_TIMEOUT", "TEMPORAL_DIAL_TIMEOUT", "TEMPORAL_WORKER_STOP_TIMEOUT"} {
		t.Run(key, func(t *testing.T) {
			t.Setenv(key, "0s")
			if _, err := Load(); err == nil { t.Fatal("zero timeout accepted") }
		})
	}
}
