# System Design Scenarios (Go Specific)

> Architecture using Go primitives.

## Scenarios

### 91. Design a Rate Limiter
- Level 1: `map[IP]int` + Mutex (Memory leak?)
- Level 2: Token Bucket using Channels (Background ticker filling bucket)
- Level 3: `golang.org/x/time/rate`
- Level 4: Distributed Redis + Lua Script

### 92. Design a Web Crawler
- Concurrency: Worker pool
- Dedup: `sync.Map` or Bloom Filter for visited URLs
- Politeness: Per-host rate limiting
- Stop condition: `errgroup` or `WaitGroup`

### 93. Design a URL Shortener (High Read)
- Cache: `sync.RWMutex` map in memory vs Redis
- ID Gen: Pre-generated block allocation to avoid DB contention

### 94. Design a Job Scheduler
- Priority: Heap (Priority Queue)
- Dispatch: `time.Timer` for next job
- Persistence: Write-Ahead Log (WAL)

### 95. Design a Logging Library
- Zero-Alloc: Use `[]byte` buffer pools
- Async: Write to channel, background flush to disk
- API: Functional options for configuration

---

## Progress
- [ ] Rate Limiter
- [ ] Web Crawler
- [ ] URL Shortener
- [ ] Job Scheduler
- [ ] Logging Library

