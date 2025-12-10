# Topic J: Unit Testing & Mocking

> **Theory:** Table-driven tests. Test coverage. Mocking with interfaces. Benchmarks. TestMain.

## Assignments

### 1. Table-Driven Tests
Write tests for a `Calculate` function using table-driven approach with subtests.

### 2. Test Coverage
Run `go test -cover` and achieve 80%+ coverage. Use `go tool cover -html` to visualize.

### 3. Mocking with Interfaces
Create an interface for external dependency. Mock it in tests.

### 4. Benchmarks
Write benchmarks comparing two implementations. Use `go test -bench`.

### 5. TestMain
Use `TestMain` for setup/teardown. Initialize test database before tests.

---

## Progress
- [ ] Assignment 1: Table-Driven Tests
- [ ] Assignment 2: Test Coverage
- [ ] Assignment 3: Mocking
- [ ] Assignment 4: Benchmarks
- [ ] Assignment 5: TestMain

---

## Key Commands

```bash
# Run tests
go test -v ./...

# Run with race detector
go test -race ./...

# Test coverage
go test -cover

# Coverage with HTML report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. -benchmem

# Run specific test
go test -run TestFunctionName
```
