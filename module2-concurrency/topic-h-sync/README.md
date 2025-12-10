# Topic H: Sync Package

> **Theory:** Mutex for state. Channels for flow.

## Assignments

### 1. Safe Counter
Fix the race condition (F3) using `sync.Mutex`.

### 2. RWMutex
Create a Cache. 100 readers, 1 writer. Use `RLock` vs `Lock`.

### 3. Singleton
Use `sync.Once` to ensure `InitDB()` runs exactly once despite concurrent calls.

### 4. Atomic
Replace Mutex with `atomic.AddInt64` for the counter. Benchmark the speed difference.

### 5. Cond
Use `sync.Cond` to broadcast a "Start" signal to 10 waiting runners.

---

## Progress
- [ ] Assignment 1: Safe Counter
- [ ] Assignment 2: RWMutex
- [ ] Assignment 3: Singleton
- [ ] Assignment 4: Atomic
- [ ] Assignment 5: Cond

