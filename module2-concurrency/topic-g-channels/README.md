# Topic G: Channels

> **Theory:** "Share memory by communicating." Buffered vs Unbuffered.

## Assignments

### 1. Ping Pong
Two goroutines passing an int back and forth on a channel, incrementing it.

### 2. Worker Pool
5 workers. `jobs` channel. `results` channel. Process 100 items.

### 3. Select Timeout
Wait for a channel or `time.After(2 * time.Second)`. Print which happened first.

### 4. Fan-In
Merge 2 channels into 1.

### 5. Graceful Close
Producer sends 10 items then closes channel. Consumer loops using `range` until closed.

---

## Progress
- [ ] Assignment 1: Ping Pong
- [ ] Assignment 2: Worker Pool
- [ ] Assignment 3: Select Timeout
- [ ] Assignment 4: Fan-In
- [ ] Assignment 5: Graceful Close

