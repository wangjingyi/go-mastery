# Machine Coding (Live Coding)

> Clean, runnable code expected in 30-45 mins.

## Challenges

### 96. Implement errgroup from scratch
- Cancel all other goroutines if one fails
- Return first error
- Wait for all to finish

### 97. Implement a Semaphore
- `Acquire(n)` and `Release(n)`
- Should block if not enough permits

### 98. Implement Thread-Safe LRU Cache
- O(1) Get/Put
- Max capacity eviction

### 99. Implement Fan-In Pattern
- Read from N channels, write to 1
- Close output only when all N are closed

### 100. Dining Philosophers Problem
- Avoid Deadlock

### 101. Implement a Connection Pool
- Get() returns connection
- Put() returns it
- Max Open Connections limit
- Timeout on Get()

### 102. Implement a Circuit Breaker
- States: Closed → Open → Half-Open
- Failure threshold and timeout

### 103. Implement Weighted Round Robin Load Balancer
- Thread-safe Next()
- Smooth weighted distribution

### 104. Implement Rolling Window Rate Limiter
- "Allow max 100 requests in last 1 minute"
- Accurate to the second

### 105. Implement Consistent Hashing
- AddNode, RemoveNode, GetNode
- Virtual nodes for balance

### 106. Implement Pub/Sub with Wildcards
- `Subscribe("user.*")` matches `Publish("user.created")`

### 107. Implement Delayed Job Queue
- Push(job, delay), Pop() blocks until ready

### 108. Implement In-Memory File System
- Mkdir, Cd, Ls, Touch

### 109. Implement Middleware Chaining
- `Chain(handler, ...middlewares)` returns `http.Handler`

### 110. Implement Concurrent Bitset
- Set(index), Clear(index), IsSet(index)
- Thread-safe

---

## Progress
- [ ] errgroup
- [ ] Semaphore
- [ ] LRU Cache
- [ ] Fan-In
- [ ] Dining Philosophers
- [ ] Connection Pool
- [ ] Circuit Breaker
- [ ] Load Balancer
- [ ] Rate Limiter
- [ ] Consistent Hashing
- [ ] Pub/Sub
- [ ] Delayed Queue
- [ ] File System
- [ ] Middleware Chain
- [ ] Concurrent Bitset

