# Go Practice

## Table of Contents
- [Go Practice](#go-practice)
  - [Table of Contents](#table-of-contents)
  - [To-do](#to-do)
  - [Notes](#notes)

## To-do
Go Language Roadmap
- [x] [Syntax](variables/variables.go)
- [x] [Conditional branches (If-else, switch-case)](functions/functions.go)
- [x] [Functions](functions/functions.go)
- [x] [Arrays](arrays/arrays.go)
- [x] [Maps](maps/maps.go)
- [x] [Loops](arrays/arrays.go)
- [x] OOP
  - [x] [Structs](structs/structs.go)
  - [x] [Interfaces](interfaces/interfaces.go)
- [x] [Pointers](pointers/pointers.go)
  - [x] [Pass by Reference](pass-by-reference.go)
- [ ] Channels
- [ ] Generics
- [ ] Go for embedded
- [ ] Tests in go

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
- Go has only one looping construct, the **for loop**
  - But for loop can be used as "regular for loop", "while", or "foreach"
- Map **always returns value** even if the `key` doesn't exist
  - But it returns a **second value** to indicate if key exists.
- Switch cases can use with or without variables. If variable isn't passed, it can be used like if branch.
- Functions can return **more than one value**
> [!TIP]
> Don't forget to use `pass by reference` when passing large arrays.
- You can also make **in-line structs** that can be used one time.
- Interfaces is similar to **abstract classes** in C++
- Pointers are exactly like C

> [!IMPORTANT]
> Exception for copying values without pointers is slices.
> 
> Because slices are using pointers to address values and if you copy a slice, you copy the addresses.
