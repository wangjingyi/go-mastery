# Topic C: Structs & Composition (Not Inheritance)

> **Theory:** Embedding fields. Promoted methods. "Accept interfaces, return structs."

## Assignments

### 1. The JSON Modeler
Map a complex nested JSON response to Go structs using struct tags `json:"field"`. Unmarshal it.

### 2. The Promoted Field
Embed `BaseEntity` (ID, CreatedAt) into `User`. Access `user.ID` directly.

### 3. The Constructor Pattern
Create a private struct `server`. Create a public `NewServer(port int) *server`. Prevent direct initialization.

### 4. The Override Trap
Embed `Base` in `Child`. Give both a `Describe()` method. Call `child.Describe()`. Call `child.Base.Describe()`.

### 5. The Mixin
Create `Drivable` and `Flyable` structs. Embed both in `FlyingCar`. Use methods from both.

---

## Progress
- [ ] Assignment 1: JSON Modeler
- [ ] Assignment 2: Promoted Field
- [ ] Assignment 3: Constructor Pattern
- [ ] Assignment 4: Override Trap
- [ ] Assignment 5: Mixin

