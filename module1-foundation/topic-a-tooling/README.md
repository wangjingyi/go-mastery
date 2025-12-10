# Topic A: Go Mod, Tooling & Workspace

> **Theory:** `go.mod` is the source of truth. `go.sum` ensures integrity. Private modules.

## Assignments

### 1. The Initialization Drill
Initialize a module `github.com/yourname/calculator`. Create a `calc` package. Import it in `cmd/main.go`.

### 2. The Dependency Injection
Import `github.com/google/uuid`. Generate a UUID. Run `go mod tidy`. Inspect `go.sum`.

### 3. The Version Pinning
Force a specific older version of `github.com/sirupsen/logrus` (v1.4.0) using `go get`. Upgrade to latest.

### 4. The Local Replace
Create a local module `my-logger`. Use `replace` in `go.mod` to redirect dependency resolution to the local folder (Essential for microservices).

### 5. The Linter Fix
Write code that intentionally breaks rules (unused vars). Configure `.golangci.yml` to be strict. Fix errors until it passes.

---

## Progress
- [ ] Assignment 1: Initialization Drill
- [ ] Assignment 2: Dependency Injection
- [ ] Assignment 3: Version Pinning
- [ ] Assignment 4: Local Replace
- [ ] Assignment 5: Linter Fix

