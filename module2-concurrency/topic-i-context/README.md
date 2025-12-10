# Topic I: Context

> **Theory:** Cancellation propagation. Request scoping.

## Assignments

### 1. Timeout Wrapper
Function sleeps 5s. Context timeout 2s. Return error immediately on timeout.

### 2. HTTP Request
`http.NewRequestWithContext`. Call a slow URL. Cancel request if it takes too long.

### 3. Tree Cancel
Cancel parent context. Verify child contexts are also cancelled.

### 4. Value Transport
Pass a "TraceID" via context through 3 function layers.

### 5. DB Loop
Run a loop doing work. Check `ctx.Err()` every iteration to abort early.

---

## Progress
- [ ] Assignment 1: Timeout Wrapper
- [ ] Assignment 2: HTTP Request
- [ ] Assignment 3: Tree Cancel
- [ ] Assignment 4: Value Transport
- [ ] Assignment 5: DB Loop

