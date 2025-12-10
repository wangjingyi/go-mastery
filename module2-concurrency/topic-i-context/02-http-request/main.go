// Assignment 2: HTTP Request
//
// Goal: http.NewRequestWithContext. Call a slow URL.
//       Cancel request if it takes too long.
//
// Instructions:
// 1. Create a context with timeout
// 2. Create HTTP request with context
// 3. If server is slow, request is automatically cancelled
//
// Run: go run main.go

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Create a context with 3 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// This URL simulates a delay (httpbin.org)
	// Change delay parameter to test timeout behavior
	url := "https://httpbin.org/delay/2" // 2 second delay (should succeed)
	// url := "https://httpbin.org/delay/5" // 5 second delay (should timeout)

	fmt.Printf("Fetching %s with 3s timeout...\n", url)

	start := time.Now()
	body, err := fetchWithContext(ctx, url)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error after %v: %v\n", elapsed, err)
	} else {
		fmt.Printf("✅ Success after %v\n", elapsed)
		fmt.Printf("Response length: %d bytes\n", len(body))
	}
}

func fetchWithContext(ctx context.Context, url string) ([]byte, error) {
	// Create request WITH context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Execute request - will be cancelled if context expires
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	return body, nil
}

// Key insight:
// - Context cancellation propagates to HTTP transport
// - Connection is closed immediately on cancel
// - No waiting for slow servers

// In production:
// - Always use context with timeouts for HTTP calls
// - Set sensible defaults (e.g., 30 seconds)
// - Consider per-endpoint timeouts based on expected latency

