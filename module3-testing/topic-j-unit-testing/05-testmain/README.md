# Assignment 5: TestMain

## Goal
Use `TestMain` for setup/teardown. Initialize test database before tests.

## Instructions

### Step 1: Create a module
```bash
cd module3-testing/topic-j-unit-testing/05-testmain
go mod init github.com/yourname/testmain-demo
```

### Step 2: Create a simple "database" to test

**db.go:**
```go
package main

import (
    "errors"
    "sync"
)

// DB is a simple in-memory database
type DB struct {
    mu   sync.RWMutex
    data map[string]string
}

var (
    globalDB    *DB
    ErrNotFound = errors.New("key not found")
)

// InitDB initializes the global database
func InitDB() {
    globalDB = &DB{
        data: make(map[string]string),
    }
}

// CloseDB closes the database
func CloseDB() {
    globalDB = nil
}

// Set stores a key-value pair
func Set(key, value string) error {
    if globalDB == nil {
        return errors.New("database not initialized")
    }
    globalDB.mu.Lock()
    defer globalDB.mu.Unlock()
    globalDB.data[key] = value
    return nil
}

// Get retrieves a value by key
func Get(key string) (string, error) {
    if globalDB == nil {
        return "", errors.New("database not initialized")
    }
    globalDB.mu.RLock()
    defer globalDB.mu.RUnlock()
    
    val, ok := globalDB.data[key]
    if !ok {
        return "", ErrNotFound
    }
    return val, nil
}

// Delete removes a key
func Delete(key string) error {
    if globalDB == nil {
        return errors.New("database not initialized")
    }
    globalDB.mu.Lock()
    defer globalDB.mu.Unlock()
    delete(globalDB.data, key)
    return nil
}

func main() {}
```

### Step 3: Create tests with TestMain

**db_test.go:**
```go
package main

import (
    "fmt"
    "os"
    "testing"
)

// TestMain runs before and after all tests
func TestMain(m *testing.M) {
    // === SETUP ===
    fmt.Println("Setting up test database...")
    InitDB()
    
    // Seed test data
    Set("existing-key", "existing-value")
    
    // === RUN TESTS ===
    exitCode := m.Run()
    
    // === TEARDOWN ===
    fmt.Println("Cleaning up test database...")
    CloseDB()
    
    // Exit with the test result code
    os.Exit(exitCode)
}

func TestSet(t *testing.T) {
    err := Set("test-key", "test-value")
    if err != nil {
        t.Fatalf("Set failed: %v", err)
    }
    
    // Verify it was set
    val, err := Get("test-key")
    if err != nil {
        t.Fatalf("Get failed: %v", err)
    }
    if val != "test-value" {
        t.Errorf("Expected 'test-value', got '%s'", val)
    }
}

func TestGet_ExistingKey(t *testing.T) {
    // This uses data seeded in TestMain
    val, err := Get("existing-key")
    if err != nil {
        t.Fatalf("Get failed: %v", err)
    }
    if val != "existing-value" {
        t.Errorf("Expected 'existing-value', got '%s'", val)
    }
}

func TestGet_NotFound(t *testing.T) {
    _, err := Get("nonexistent-key")
    if err != ErrNotFound {
        t.Errorf("Expected ErrNotFound, got %v", err)
    }
}

func TestDelete(t *testing.T) {
    // First set a value
    Set("to-delete", "value")
    
    // Delete it
    err := Delete("to-delete")
    if err != nil {
        t.Fatalf("Delete failed: %v", err)
    }
    
    // Verify it's gone
    _, err = Get("to-delete")
    if err != ErrNotFound {
        t.Errorf("Expected ErrNotFound after delete, got %v", err)
    }
}
```

### Step 4: Run tests and observe setup/teardown
```bash
go test -v
```

**Output:**
```
Setting up test database...
=== RUN   TestSet
--- PASS: TestSet (0.00s)
=== RUN   TestGet_ExistingKey
--- PASS: TestGet_ExistingKey (0.00s)
=== RUN   TestGet_NotFound
--- PASS: TestGet_NotFound (0.00s)
=== RUN   TestDelete
--- PASS: TestDelete (0.00s)
PASS
Cleaning up test database...
ok      github.com/yourname/testmain-demo    0.002s
```

## Advanced: Subtests with Setup/Teardown

```go
func TestWithSubtests(t *testing.T) {
    // Setup for this test group
    Set("group-key", "group-value")
    
    t.Run("SubTest1", func(t *testing.T) {
        val, _ := Get("group-key")
        if val != "group-value" {
            t.Error("unexpected value")
        }
    })
    
    t.Run("SubTest2", func(t *testing.T) {
        // Each subtest can have its own setup
        Set("sub-key", "sub-value")
        val, _ := Get("sub-key")
        if val != "sub-value" {
            t.Error("unexpected value")
        }
    })
    
    // Teardown for this test group
    Delete("group-key")
}
```

## Using t.Cleanup (Go 1.14+)

```go
func TestWithCleanup(t *testing.T) {
    // Setup
    key := "cleanup-test"
    Set(key, "value")
    
    // Register cleanup (runs after test completes)
    t.Cleanup(func() {
        Delete(key)
    })
    
    // Test logic
    val, err := Get(key)
    if err != nil {
        t.Fatal(err)
    }
    if val != "value" {
        t.Error("unexpected value")
    }
    
    // Cleanup runs automatically after this
}
```

## Key Learnings
- `TestMain` runs once before/after ALL tests in a package
- Must call `m.Run()` and `os.Exit()` with its result
- Use for expensive setup (DB connections, Docker containers)
- Use `t.Cleanup()` for per-test cleanup (Go 1.14+)
- Subtests can have their own setup/teardown

