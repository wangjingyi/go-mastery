# Assignment 2: The Dependency Injection

## Goal
Import `github.com/google/uuid`. Generate a UUID. Run `go mod tidy`. Inspect `go.sum`.

## Instructions

### Step 1: Create a new module
```bash
cd module1-foundation/topic-a-tooling/02-dependency-injection
go mod init github.com/yourname/uuid-demo
```

### Step 2: Create main.go
```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    // Generate a new UUID
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
    
    // Parse a UUID string
    parsed, err := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Parsed UUID:", parsed)
}
```

### Step 3: Download the dependency
```bash
go get github.com/google/uuid
```

### Step 4: Clean up dependencies
```bash
go mod tidy
```

### Step 5: Inspect the files

**Check go.mod:**
```bash
cat go.mod
```
You should see the `require` directive with the uuid package.

**Check go.sum:**
```bash
cat go.sum
```
This file contains cryptographic checksums for each module version.

### Step 6: Run the program
```bash
go run main.go
```

## Key Learnings
- `go get` downloads and adds dependencies
- `go mod tidy` removes unused dependencies and adds missing ones
- `go.mod` declares dependencies
- `go.sum` ensures reproducible builds with checksums

## Questions to Answer
1. What version of uuid was installed?
2. How many lines are in go.sum? Why?
3. What happens if you delete go.sum and run `go mod tidy`?

