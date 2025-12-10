# Topic B: Pointers & Memory Mechanics

> **Theory:** Stack vs Heap. Escape Analysis. Value vs Pointer semantics.

## Assignments

### 1. The Swap Function
Write `Swap(a, b *int)`. Verify variables change in main.

### 2. The Heavy Struct
Benchmark passing a `[10000]int` struct by Value vs Pointer. Measure the speed difference.

### 3. The Escape Analyst
Write a function returning a pointer to a local variable. Run `go build -gcflags="-m"` to verify it "escapes to heap".

### 4. The Mutation Bug
Create a method `func (u User) Birthday()` (value receiver). Call it. Why didn't age change? Fix it with a pointer receiver.

### 5. The Nil Panic Guard
Write a function accepting `*User`. If passed nil, return an error instead of panicking.

---

## Progress
- [ ] Assignment 1: Swap Function
- [ ] Assignment 2: Heavy Struct
- [ ] Assignment 3: Escape Analyst
- [ ] Assignment 4: Mutation Bug
- [ ] Assignment 5: Nil Panic Guard

