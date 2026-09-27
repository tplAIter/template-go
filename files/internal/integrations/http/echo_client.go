package httpx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type EchoClient struct {
	baseURL string
	client  *http.Client
	logger  *slog.Logger
}

func NewEchoClient(baseURL string, logger *slog.Logger) (*EchoClient, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("base URL is empty")
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &EchoClient{baseURL: baseURL, client: http.DefaultClient, logger: logger}, nil
}
func (c *EchoClient) Echo(ctx context.Context, message string) (string, error) {
	body, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, &bytesReader{data: body})
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("echo status %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var v map[string]string
	if err := json.Unmarshal(data, &v); err != nil {
		return "", err
	}
	return v["message"], nil
}

type bytesReader struct {
	data []byte
	off  int
}

func (r *bytesReader) Read(p []byte) (int, error) {
	if r.off >= len(r.data) {
		return 0, io.EOF
	}
	n := copy(p, r.data[r.off:])
	r.off += n
	return n, nil
}
func (r *bytesReader) Close() error { return nil }
