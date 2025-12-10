// Assignment 4: Value Transport
//
// Goal: Pass a "TraceID" via context through 3 function layers.
//
// Instructions:
// 1. Use context.WithValue to attach request-scoped data
// 2. Pass context through function calls
// 3. Extract values at any layer
//
// Run: go run main.go

package main

import (
	"context"
	"fmt"
)

// Define typed keys to avoid collisions
type contextKey string

const (
	traceIDKey contextKey = "traceID"
	userIDKey  contextKey = "userID"
)

// Helper functions for type-safe access
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	if v := ctx.Value(traceIDKey); v != nil {
		return v.(string)
	}
	return "unknown"
}

func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) int {
	if v := ctx.Value(userIDKey); v != nil {
		return v.(int)
	}
	return 0
}

// Simulated application layers

func HandleRequest(ctx context.Context) {
	// Layer 1: HTTP Handler
	traceID := GetTraceID(ctx)
	fmt.Printf("[Handler] Processing request (trace: %s)\n", traceID)

	// Add user info
	ctx = WithUserID(ctx, 12345)

	// Call service layer
	ProcessOrder(ctx)
}

func ProcessOrder(ctx context.Context) {
	// Layer 2: Service
	traceID := GetTraceID(ctx)
	userID := GetUserID(ctx)
	fmt.Printf("[Service] Processing order for user %d (trace: %s)\n", userID, traceID)

	// Call repository layer
	SaveToDatabase(ctx)
}

func SaveToDatabase(ctx context.Context) {
	// Layer 3: Repository
	traceID := GetTraceID(ctx)
	userID := GetUserID(ctx)
	fmt.Printf("[Repository] Saving data for user %d (trace: %s)\n", userID, traceID)
}

func main() {
	// Simulate incoming request with trace ID
	ctx := context.Background()
	ctx = WithTraceID(ctx, "abc-123-xyz")

	fmt.Println("=== Request Processing ===")
	HandleRequest(ctx)

	fmt.Println("\n=== Another Request ===")
	ctx2 := context.Background()
	ctx2 = WithTraceID(ctx2, "def-456-uvw")
	HandleRequest(ctx2)
}

// Best practices for context.Value:
// 1. Use typed keys (not strings) to avoid collisions
// 2. Only use for request-scoped data (trace ID, user ID, auth tokens)
// 3. Don't use for optional parameters or configuration
// 4. Keep values immutable
// 5. Provide helper functions for type-safe access

