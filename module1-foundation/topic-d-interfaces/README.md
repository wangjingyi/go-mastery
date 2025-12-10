# Topic D: Interfaces

> **Theory:** Implicit interfaces. Duck typing. The `any` type.

## Assignments

### 1. The Shape Solver
Define `Shape` interface (Area method). Implement `Circle`, `Rectangle`. Write `PrintArea(s Shape)`.

### 2. The Writer Adapter
Create a `ConsoleWriter` struct. Implement `Write([]byte) (int, error)`. Pass it to `fmt.Fprintf` (which expects `io.Writer`).

### 3. The Type Switch
Create a `map[string]any`. Store int, string, struct. Iterate and use `switch v := val.(type)` to handle each.

### 4. The Mock Interface
Define `PaymentProcessor` interface. Create `StripeProcessor` (real) and `MockProcessor` (fake). Swap them in main.

### 5. The Interface Segregation
Take a massive `GodInterface`. Break it into `Reader`, `Writer`, `Closer`. Demonstrate combining them `interface{ Reader; Writer }`.

---

## Progress
- [ ] Assignment 1: Shape Solver
- [ ] Assignment 2: Writer Adapter
- [ ] Assignment 3: Type Switch
- [ ] Assignment 4: Mock Interface
- [ ] Assignment 5: Interface Segregation

