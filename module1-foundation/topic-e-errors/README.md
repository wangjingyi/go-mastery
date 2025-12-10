# Topic E: Error Handling

> **Theory:** Errors are values. `errors.Is`, `errors.As`. Wrapping.

## Assignments

### 1. The Divider
Return custom error if dividing by zero. Handle it.

### 2. The Sentinel
Define `var ErrNotFound = errors.New(...)`. Return it. Check for it using `errors.Is`.

### 3. The Rich Error
Create struct `AppError` with Code and Message. Implement `Error()`. Use `errors.As` to retrieve the Code.

### 4. The Wrapper
Call a failing function. Return `fmt.Errorf("db failed: %w", err)`. Print the full chain.

### 5. The Safe Recovery
Write a web handler that panics. Use `defer` and `recover` to catch the panic and log it instead of crashing.

---

## Progress
- [ ] Assignment 1: Divider
- [ ] Assignment 2: Sentinel
- [ ] Assignment 3: Rich Error
- [ ] Assignment 4: Wrapper
- [ ] Assignment 5: Safe Recovery

