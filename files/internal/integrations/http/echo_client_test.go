package httpx

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoClientEcho(t *testing.T) {
	// External-service stub through httptest — testing at the boundary.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body) // echo: return the same {"message": ...} body
	}))
	defer srv.Close()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	client, err := NewEchoClient(srv.URL, logger)
	if err != nil {
		t.Fatalf("NewEchoClient: %v", err)
	}

	got, err := client.Echo(context.Background(), "ping")
	if err != nil {
		t.Fatalf("Echo: %v", err)
	}
	if got != "ping" {
		t.Errorf("Echo = %q, want %q", got, "ping")
	}
}
