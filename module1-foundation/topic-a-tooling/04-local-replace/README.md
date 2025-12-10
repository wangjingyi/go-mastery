# Assignment 4: The Local Replace

## Goal
Create a local module `my-logger`. Use `replace` in `go.mod` to redirect dependency resolution to the local folder.

**This is ESSENTIAL for microservices development!**

## Instructions

### Step 1: Create directory structure
```bash
cd module1-foundation/topic-a-tooling/04-local-replace
mkdir -p my-logger
mkdir -p app
```

### Step 2: Create the local logger module

**my-logger/go.mod:**
```bash
cd my-logger
go mod init github.com/yourname/my-logger
```

**my-logger/logger.go:**
```go
package mylogger

import (
    "fmt"
    "time"
)

// Log prints a message with timestamp
func Log(message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s\n", timestamp, message)
}

// LogError prints an error message
func LogError(message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] ERROR: %s\n", timestamp, message)
}

// LogInfo prints an info message
func LogInfo(message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] INFO: %s\n", timestamp, message)
}
```

### Step 3: Create the app that uses the logger

**app/go.mod:**
```bash
cd ../app
go mod init github.com/yourname/app
```

**app/main.go:**
```go
package main

import mylogger "github.com/yourname/my-logger"

func main() {
    mylogger.Log("Application started")
    mylogger.LogInfo("Processing request...")
    mylogger.LogError("Something went wrong!")
    mylogger.Log("Application finished")
}
```

### Step 4: Add the replace directive

Edit **app/go.mod** to add:
```
module github.com/yourname/app

go 1.22

require github.com/yourname/my-logger v0.0.0

replace github.com/yourname/my-logger => ../my-logger
```

### Step 5: Run the app
```bash
cd app
go run main.go
```

## Expected Output
```
[2024-01-15 10:30:00] Application started
[2024-01-15 10:30:00] INFO: Processing request...
[2024-01-15 10:30:00] ERROR: Something went wrong!
[2024-01-15 10:30:00] Application finished
```

## Key Learnings
- `replace` redirects module resolution to local path
- Essential for developing multiple services together
- The replaced module doesn't need to be published
- Version can be `v0.0.0` for local development

## Real-World Usage
In microservices:
```
replace (
    github.com/company/shared-lib => ../shared-lib
    github.com/company/proto => ../proto
)
```

