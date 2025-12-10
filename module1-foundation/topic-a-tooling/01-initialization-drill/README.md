# Assignment 1: The Initialization Drill

## Goal
Initialize a module `github.com/yourname/calculator`. Create a `calc` package. Import it in `cmd/main.go`.

## Instructions

### Step 1: Initialize the module
```bash
cd module1-foundation/topic-a-tooling/01-initialization-drill
go mod init github.com/yourname/calculator
```

### Step 2: Create the calc package
Create `calc/calc.go`:
```go
package calc

// Add returns the sum of two integers
func Add(a, b int) int {
    // TODO: implement
}

// Subtract returns the difference of two integers  
func Subtract(a, b int) int {
    // TODO: implement
}
```

### Step 3: Create the main program
Create `cmd/main.go`:
```go
package main

import (
    "fmt"
    "github.com/yourname/calculator/calc"
)

func main() {
    result := calc.Add(2, 3)
    fmt.Printf("2 + 3 = %d\n", result)
}
```

### Step 4: Run it
```bash
go run cmd/main.go
```

## Expected Structure
```
01-initialization-drill/
├── go.mod              ← You create this
├── calc/
│   └── calc.go         ← You create this
└── cmd/
    └── main.go         ← You create this
```

## Key Learnings
- `go mod init` creates a new module
- Package names match directory names
- Import paths are based on module name + directory path
- `cmd/` is a convention for executable entry points

## Verification
```bash
go run cmd/main.go
# Should output: 2 + 3 = 5
```

