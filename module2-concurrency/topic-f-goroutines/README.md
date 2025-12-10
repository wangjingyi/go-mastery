# Topic F: Goroutines

> **Theory:** M:N Scheduler. 2kb stack.

## Assignments

### 1. The Spawner
Launch 10,000 goroutines that print "Done". Measure execution time.

### 2. The WaitGroup
Sync the 10,000 goroutines using `sync.WaitGroup` so main doesn't exit early.

### 3. The Race Condition
Increment a global counter from 1000 goroutines. Run with `go run -race`.

### 4. The Loop Trap
Launch goroutines inside a `for i := 0` loop printing `i`. Observe they all print the same number. Fix by passing `i` as argument.

### 5. The Heartbeat
Background goroutine printing "Pulse" every 500ms. Stop it when main exits.

---

## Progress
- [ ] Assignment 1: Spawner
- [ ] Assignment 2: WaitGroup
- [ ] Assignment 3: Race Condition
- [ ] Assignment 4: Loop Trap
- [ ] Assignment 5: Heartbeat

