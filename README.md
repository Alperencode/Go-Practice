# Go Practice

## Table of Contents
- [Go Practice](#go-practice)
  - [Table of Contents](#table-of-contents)
  - [To-do](#to-do)
  - [Notes](#notes)

## To-do
Go Language Roadmap
- [x] Syntax
- [x] Conditional branches (If-else, switch-case)
- [x] Functions
- [x] Arrays
- [x] Maps
- [x] Loops
- [ ] OOP
  - [ ] Structs
  - [ ] Interfaces
- [ ] Go for embedded
- [ ] Tests in go
- [ ] Go mod details (tidy and vendor)

Specific Things
- [x] Keyword: defer
  - A defer statement defers the execution of a function until the surrounding function returns ([source](https://go.dev/tour/flowcontrol/12))
- [ ] Interface
- [ ] Goroutines
- [ ] tidy and vendor

## Notes

- `len()` returns the **byte size** of a string. So if char is not in ASCII table it could stored as more than one byte.
- If you don't specify the length of the array, it will be called **`slice`**
- Use make() to initialize it with declaring capacity
- Map **always returns value** even if the `key` doesn't exists
  - But it returns a **second value** to indicate if key exists.
- Switch cases can use with or without variables. If variable isn't passed, it can be used like if branch.
- Functions can return more than one value