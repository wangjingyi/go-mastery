// Assignment 3: The Rich Error
//
// Goal: Create struct AppError with Code and Message.
//       Implement Error(). Use errors.As to retrieve the Code.
//
// Instructions:
// 1. Create AppError struct with Code, Message, and optional Cause
// 2. Implement the error interface (Error() method)
// 3. Use errors.As to extract the structured error
//
// Key insight: errors.As is for extracting error types, errors.Is for identity

package main

import (
	"errors"
	"fmt"
	"net/http"
)

// AppError is a rich error with structured information
type AppError struct {
	Code    int    // HTTP-like status code
	Message string // Human-readable message
	Op      string // Operation that failed
	Err     error  // Underlying error (for wrapping)
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %s - %v", e.Code, e.Op, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Op, e.Message)
}

// Unwrap allows errors.Is and errors.As to traverse the chain
func (e *AppError) Unwrap() error {
	return e.Err
}

// Helper constructors for common errors
func NotFoundError(op, message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Op: op, Message: message}
}

func BadRequestError(op, message string, cause error) *AppError {
	return &AppError{Code: http.StatusBadRequest, Op: op, Message: message, Err: cause}
}

func InternalError(op string, cause error) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Op: op, Message: "internal error", Err: cause}
}

// Simulated database operation
func GetUserFromDB(id int) (string, error) {
	if id <= 0 {
		return "", BadRequestError("GetUserFromDB", "invalid user ID", nil)
	}
	if id == 999 {
		return "", NotFoundError("GetUserFromDB", fmt.Sprintf("user %d not found", id))
	}
	if id == 500 {
		// Wrap an underlying error
		dbErr := errors.New("connection refused")
		return "", InternalError("GetUserFromDB", dbErr)
	}
	return "Alice", nil
}

func main() {
	testCases := []int{1, -1, 999, 500}

	for _, id := range testCases {
		fmt.Printf("\n=== Getting user %d ===\n", id)
		
		name, err := GetUserFromDB(id)
		if err != nil {
			fmt.Println("Error:", err)

			// Use errors.As to extract the AppError
			var appErr *AppError
			if errors.As(err, &appErr) {
				fmt.Printf("  Code: %d\n", appErr.Code)
				fmt.Printf("  Op: %s\n", appErr.Op)
				fmt.Printf("  Message: %s\n", appErr.Message)
				
				// You can now make decisions based on the code
				switch appErr.Code {
				case http.StatusNotFound:
					fmt.Println("  -> Return 404 to client")
				case http.StatusBadRequest:
					fmt.Println("  -> Return 400 to client")
				case http.StatusInternalServerError:
					fmt.Println("  -> Log this, return 500 to client")
				}
			}
		} else {
			fmt.Printf("Found user: %s\n", name)
		}
	}
}

