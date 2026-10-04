package main

import (
	"strings"
	"testing"
)

func TestInvalidListenAddressReturnsError(t *testing.T) {
	t.Setenv("APP_LISTEN_ADDR", "invalid::address")
	if err := run(); err == nil || !strings.Contains(err.Error(), "http server") {
		t.Fatalf("expected startup failure, got %v", err)
	}
}
