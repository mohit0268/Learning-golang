# Go Structs — Complete Guide (Industry Perspective)

A **struct** in Go is a composite data type that groups related variables together. It is the primary way to model data in Go applications and acts as the foundation for building APIs, services, and database models.

Unlike traditional OOP languages, Go does not use classes. Instead, it relies on **structs + methods + interfaces** to build scalable systems.

This guide explains:

* What structs are
* How they work internally
* Memory behavior
* Industry use cases
* Bottlenecks and fixes
* Best practices used in production systems

---

# Table of Contents

1. What is a Struct
2. Struct Syntax
3. Memory Representation
4. Struct Initialization Patterns
5. Struct Methods
6. Pointer vs Value Receivers
7. Struct Tags
8. Struct Composition (Embedding)
9. Real Industry Examples
10. Common Use Cases
11. Performance Bottlenecks
12. Common Pitfalls
13. Struct vs Map
14. Production Architecture
15. Best Practices

---

# 1. What is a Struct

A **struct** is a custom data type that groups multiple fields together.

Think of it as a **blueprint for objects**.

Example:

```go
type User struct {
    ID    int
    Name  string
    Email string
}
```

Creating an instance:

```go
user := User{
    ID: 1,
    Name: "Mohit",
    Email: "mohit@email.com",
}
```

Conceptual representation:

```
User
 ├── ID
 ├── Name
 └── Email
```

---

# 2. Struct Syntax

Basic structure definition:

```go
type StructName struct {
    FieldName DataType
}
```

Example:

```go
type Product struct {
    ID    int
    Price float64
}
```

---

# 3. Memory Representation (Behind the Scenes)

Struct fields are stored **contiguously in memory**.

Example:

```go
type Product struct {
    ID    int
    Price float64
}
```

Memory layout:

```
| ID (8 bytes) | Price (8 bytes) |
```

Go may add **padding** for alignment.

Example:

```go
type BadStruct struct {
    A bool
    B int64
}
```

Memory layout:

```
| bool (1 byte) | padding (7 bytes) | int64 (8 bytes) |
```

Optimized version:

```go
type GoodStruct struct {
    B int64
    A bool
}
```

Rule:

Always arrange fields **largest → smallest** to minimize memory waste.

---

# 4. Struct Initialization Patterns

## Standard Initialization

```go
user := User{
    ID: 1,
    Name: "Mohit",
}
```

---

## Zero Value Initialization

```go
var user User
```

Default values:

```
int    -> 0
string -> ""
bool   -> false
```

---

## Pointer Initialization

```go
user := &User{
    ID: 1,
    Name: "Mohit",
}
```

Memory concept:

```
Stack
 user → pointer

Heap
 User object
```

Pointers help avoid **large struct copying**.

---

# 5. Struct Methods

Structs can have methods attached to them.

Example:

```go
type User struct {
    Name string
}

func (u User) Greet() string {
    return "Hello " + u.Name
}
```

Usage:

```go
user := User{Name: "Mohit"}
fmt.Println(user.Greet())
```

Output:

```
Hello Mohit
```

---

# 6. Pointer vs Value Receivers

## Value Receiver

```go
func (u User) UpdateName(name string) {
    u.Name = name
}
```

This **does not modify the original struct** because Go passes a copy.

---

## Pointer Receiver

```go
func (u *User) UpdateName(name string) {
    u.Name = name
}
```

This **modifies the original struct**.

Use pointer receivers when:

* Struct is large
* You need mutation
* Performance matters

---

# 7. Struct Tags

Struct tags provide metadata for frameworks and libraries.

Example:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

Used for serialization:

```go
json.Marshal(user)
```

Output:

```json
{
 "id":1,
 "name":"Mohit",
 "email":"mohit@email.com"
}
```

Common tag use cases:

* JSON APIs
* Database ORM
* Validation
* Configuration parsing

---

# 8. Struct Composition (Embedding)

Go uses **composition instead of inheritance**.

Example:

```go
type Address struct {
    City    string
    Country string
}

type User struct {
    Name string
    Address
}
```

Usage:

```go
user := User{
    Name: "Mohit",
    Address: Address{
        City: "Delhi",
        Country: "India",
    },
}
```

Access fields directly:

```
user.City
```

---

# 9. Real Industry Example (Backend API)

Example model for an API:

```go
type Order struct {
    ID        string
    UserID    string
    Amount    float64
    CreatedAt time.Time
}
```

Typical backend flow:

```
HTTP Request
     ↓
Controller
     ↓
Service Layer
     ↓
Repository Layer
     ↓
Database
```

Structs act as **data contracts across layers**.

---

# 10. Common Struct Use Cases

## API Request Models

```go
type LoginRequest struct {
 Email string
 Password string
}
```

---

## Database Models

```go
type User struct {
 ID uint
 Name string
 CreatedAt time.Time
}
```

---

## Configuration

```go
type Config struct {
 Port int
 DBUrl string
}
```

---

## Event Messaging (Microservices)

```go
type PaymentEvent struct {
 OrderID string
 Amount float64
 Status string
}
```

---

# 11. Performance Bottlenecks

## Large Struct Copies

Bad:

```go
func process(u User)
```

Fix:

```go
func process(u *User)
```

---

## Poor Field Ordering

Incorrect ordering increases padding.

Correct ordering:

```
int64
int32
bool
```

---

## Unnecessary Heap Allocation

Bad:

```go
user := &User{}
```

If pointer not required, use value type.

---

## Slow JSON Serialization

Large structs with reflection can be slow.

Solutions used in production:

* Faster JSON libraries
* Code-generated serializers

---

# 12. Common Pitfalls

## Nil Pointer Panic

Bad:

```go
var user *User
user.Name = "Mohit"
```

Fix:

```go
user := &User{}
```

---

## Unexported Fields

Lowercase fields are private.

Bad:

```go
type user struct {
 name string
}
```

Correct:

```go
type User struct {
 Name string
}
```

---

## Struct Comparison Limitations

Structs containing slices or maps cannot be compared directly.

Example:

```go
type User struct {
 Tags []string
}
```

This struct cannot be used in equality comparison.

---

# 13. Struct vs Map

## Struct

Pros:

* Type safety
* Faster
* Compile-time validation
* Better IDE support

## Map

Pros:

* Flexible
* Dynamic keys

Example map:

```go
map[string]interface{}
```

Industry rule:

**Use structs when schema is known.**

---

# 14. Production Architecture Example

Typical Go project structure:

```
project/
│
├── models/
│   ├── user.go
│   └── order.go
│
├── services/
│   └── user_service.go
│
├── repositories/
│   └── user_repository.go
│
├── handlers/
│   └── user_handler.go
│
└── main.go
```

Structs define **domain models used across layers**.

---

# 15. Best Practices

✔ Use pointer receivers for large structs
✔ Keep structs small and focused
✔ Order fields for memory alignment
✔ Use struct tags properly
✔ Avoid circular dependencies
✔ Separate API models and database models
✔ Prefer composition over inheritance

---

# Mental Model

Think of structs as **data contracts across your system**.

```
Database
   ↓
Repository Struct
   ↓
Service Struct
   ↓
API Response Struct
   ↓
Client
```

Structs form the backbone of Go applications and are essential for building scalable backend systems.

---

# Conclusion

Structs are one of the most important building blocks in Go. When combined with **interfaces, methods, and composition**, they enable developers to build simple, maintainable, and highly performant systems.

Understanding struct memory layout, pointer usage, and serialization behavior is critical when building production-grade services.
