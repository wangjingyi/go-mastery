# Assignment 3: The Version Pinning

## Goal
Force a specific older version of `github.com/sirupsen/logrus` (v1.4.0) using `go get`. Then upgrade to latest.

## Instructions

### Step 1: Create a new module
```bash
cd module1-foundation/topic-a-tooling/03-version-pinning
go mod init github.com/yourname/logrus-demo
```

### Step 2: Install a SPECIFIC old version
```bash
go get github.com/sirupsen/logrus@v1.4.0
```

### Step 3: Verify the version
```bash
cat go.mod
```
You should see: `require github.com/sirupsen/logrus v1.4.0`

### Step 4: Create main.go
```go
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.Info("Hello from logrus!")
    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A walrus appears")
    
    log.Warn("This is a warning")
    log.Error("This is an error")
}
```

### Step 5: Run with old version
```bash
go run main.go
```

### Step 6: Upgrade to latest
```bash
go get github.com/sirupsen/logrus@latest
```

### Step 7: Check the version change
```bash
cat go.mod
```
The version should now be higher (e.g., v1.9.3).

### Step 8: Run with new version
```bash
go run main.go
```

## Key Learnings
- `@v1.4.0` syntax pins to a specific version
- `@latest` upgrades to the latest version
- go.mod records the exact version
- Semantic versioning: vMAJOR.MINOR.PATCH

## Bonus Challenge
1. Try installing a version that doesn't exist: `go get github.com/sirupsen/logrus@v99.0.0`
2. List all available versions: `go list -m -versions github.com/sirupsen/logrus`

