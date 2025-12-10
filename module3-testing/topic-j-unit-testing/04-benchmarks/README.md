# Assignment 4: Benchmarks

## Goal
Write benchmarks comparing two implementations. Use `go test -bench`.

## Instructions

### Step 1: Create a module
```bash
cd module3-testing/topic-j-unit-testing/04-benchmarks
go mod init github.com/yourname/bench-demo
```

### Step 2: Create two implementations to compare

**concat.go:**
```go
package main

import (
    "bytes"
    "strings"
)

// ConcatPlus uses + operator (slow for many strings)
func ConcatPlus(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}

// ConcatBuilder uses strings.Builder (fast)
func ConcatBuilder(strs []string) string {
    var builder strings.Builder
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}

// ConcatBuffer uses bytes.Buffer
func ConcatBuffer(strs []string) string {
    var buffer bytes.Buffer
    for _, s := range strs {
        buffer.WriteString(s)
    }
    return buffer.String()
}

// ConcatJoin uses strings.Join (idiomatic)
func ConcatJoin(strs []string) string {
    return strings.Join(strs, "")
}

func main() {}
```

### Step 3: Create benchmark tests

**concat_test.go:**
```go
package main

import (
    "testing"
)

// Generate test data
func generateStrings(n int) []string {
    strs := make([]string, n)
    for i := 0; i < n; i++ {
        strs[i] = "hello"
    }
    return strs
}

var result string // Prevent compiler optimization

func BenchmarkConcatPlus(b *testing.B) {
    strs := generateStrings(100)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        result = ConcatPlus(strs)
    }
}

func BenchmarkConcatBuilder(b *testing.B) {
    strs := generateStrings(100)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        result = ConcatBuilder(strs)
    }
}

func BenchmarkConcatBuffer(b *testing.B) {
    strs := generateStrings(100)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        result = ConcatBuffer(strs)
    }
}

func BenchmarkConcatJoin(b *testing.B) {
    strs := generateStrings(100)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        result = ConcatJoin(strs)
    }
}

// Sub-benchmarks for different sizes
func BenchmarkConcat(b *testing.B) {
    sizes := []int{10, 100, 1000}
    
    for _, size := range sizes {
        strs := generateStrings(size)
        
        b.Run("Plus/"+string(rune(size)), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                result = ConcatPlus(strs)
            }
        })
        
        b.Run("Builder/"+string(rune(size)), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                result = ConcatBuilder(strs)
            }
        })
    }
}
```

### Step 4: Run benchmarks
```bash
# Run all benchmarks
go test -bench=.

# Run with memory allocation stats
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkConcatBuilder

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5
```

### Step 5: Analyze results
```
BenchmarkConcatPlus-8       5000    234567 ns/op    123456 B/op    99 allocs/op
BenchmarkConcatBuilder-8  200000      8901 ns/op      1024 B/op     1 allocs/op
```

**Reading the output:**
- `8` = GOMAXPROCS (CPU cores used)
- `5000` = iterations run
- `234567 ns/op` = nanoseconds per operation
- `123456 B/op` = bytes allocated per operation
- `99 allocs/op` = allocations per operation

### Step 6: Compare results
```bash
# Install benchstat
go install golang.org/x/perf/cmd/benchstat@latest

# Save results
go test -bench=. -count=5 > old.txt
# Make changes...
go test -bench=. -count=5 > new.txt

# Compare
benchstat old.txt new.txt
```

## Key Learnings
- `b.N` is set automatically by the framework
- Use `b.ResetTimer()` after setup code
- `-benchmem` shows allocations (critical for performance)
- Lower ns/op and B/op is better
- strings.Builder is ~26x faster than `+` for 100 strings!

