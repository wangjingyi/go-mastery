# Assignment 5: The Linter Fix

## Goal
Write code that intentionally breaks linter rules (unused vars, etc.). Configure `.golangci.yml` to be strict. Fix all errors until it passes.

## Instructions

### Step 1: Create a new module
```bash
cd module1-foundation/topic-a-tooling/05-linter-fix
go mod init github.com/yourname/linter-demo
```

### Step 2: Create buggy code

**main.go** (with intentional linter errors):
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Error 1: Unused variable
    unusedVar := "I'm never used"
    
    // Error 2: Error not checked
    os.Open("nonexistent.txt")
    
    // Error 3: Ineffective assignment
    x := 5
    x = 10
    x = 15
    
    // Error 4: Variable shadowing
    err := doSomething()
    if true {
        err := doSomethingElse()  // Shadows outer err
        fmt.Println(err)
    }
    fmt.Println(err)
    
    fmt.Println(x)
}

func doSomething() error {
    return nil
}

func doSomethingElse() error {
    return nil
}
```

### Step 3: Create strict linter config

**.golangci.yml:**
```yaml
run:
  timeout: 5m

linters:
  enable:
    - errcheck      # Check for unchecked errors
    - gosimple      # Simplify code
    - govet         # Vet examines Go source code
    - ineffassign   # Detect ineffective assignments
    - staticcheck   # Static analysis
    - unused        # Check for unused code
    - gocritic      # Highly extensible linter
    - revive        # Fast, configurable linter

linters-settings:
  govet:
    check-shadowing: true
  errcheck:
    check-type-assertions: true
    check-blank: true
```

### Step 4: Run the linter
```bash
golangci-lint run
```

You should see errors like:
- `unusedVar declared but not used`
- `Error return value of os.Open is not checked`
- `ineffective assignment to x`
- `shadow: declaration of err shadows declaration`

### Step 5: Fix each error

1. **Unused variable**: Remove it or use it
2. **Unchecked error**: `f, err := os.Open(...); if err != nil {...}`
3. **Ineffective assignment**: Remove redundant assignments
4. **Shadowing**: Use `=` instead of `:=` or rename the variable

### Step 6: Run linter until clean
```bash
golangci-lint run
# Should output nothing (no errors)
```

## Fixed Code (Solution)
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Fixed: Removed unused variable
    
    // Fixed: Error is now checked
    _, err := os.Open("nonexistent.txt")
    if err != nil {
        fmt.Println("Expected error:", err)
    }
    
    // Fixed: Only one assignment
    x := 15
    
    // Fixed: No shadowing - reuse outer err
    err = doSomething()
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(x)
}

func doSomething() error {
    return nil
}
```

## Key Learnings
- Linters catch bugs before they reach production
- `golangci-lint` combines many linters in one tool
- Shadowing can cause subtle bugs
- Always handle errors explicitly

