# Go Practice

## Table of Contents
- [Go Practice](#go-practice)
  - [Table of Contents](#table-of-contents)
  - [Roadmap](#roadmap)
  - [Notes](#notes)
  - [Specific Things](#specific-things)
    - [Keyword defer](#keyword-defer)
    - [Goroutines](#goroutines)

## Roadmap

- [x] [Syntax](variables/variables.go)
- [x] [Conditional branches](functions/functions.go)
- [x] [Functions](functions/functions.go)
- [x] [Arrays](arrays/arrays.go)
- [x] [Maps](maps/maps.go)
- [x] [Loops](arrays/arrays.go)
- [x] [Structs](structs/structs.go)
- [x] [Interfaces](interfaces/interfaces.go)
- [x] [Pointers](pointers/pointers.go)
  - [x] [Pass by Reference](pass-by-reference/pass-by-reference.go)
- [x] Channels
- [ ] Generics
- [ ] Go for embedded
- [ ] Tests in go

Specific Things
- [x] [Keyword: defer](#keyword-defer)
- [x] [Goroutines](#goroutines)
- [ ] tidy and vendor

## Notes

- `len()` returns the **byte size** of a string. So if char is not in ASCII table it could stored as more than one byte.
- If you don't specify the length of the array, it will be a **`slice`**
- Use make() to initialize things with declaring capacity
- Go has only one looping construct, the **for loop**
  - But for loop can be used as "regular for loop", "while", or "foreach"
- Map **always returns value** even if the `key` doesn't exist
  - But it returns a **second value** to indicate if key exists.
- Switch cases can be used with or without variables. If variable isn't passed, it can be used like if branch.
- Functions can return **more than one value**
> [!TIP]
> Don't forget to use `pass by reference` when passing large arrays.
- You can also make **in-line structs** that can be used one time.
- Interfaces is similar to **abstract classes** in C++
- Pointers are exactly like C

> [!IMPORTANT]
> Exception for copying values without pointers is slices.
> 
> Because slices use pointers to address values and if you copy a slice, you copy the addresses.

<br><hr>

## Specific Things

### Keyword defer

- A defer statement defers the execution of a function until the surrounding function returns
  - souce [here](https://go.dev/tour/flowcontrol/12)

### Goroutines

- concurrency != parallel execution
- For example:
  - Program can execute other processes while waiting for database to respond
- Device need to have a multicore CPU to execute processes in parallel

- Usage
  - use `go` keyword to use goroutines (async functions)
  - import sync
    - use `WaitGroup`: a counter for tasks
    - use `wg.Add(1)` before using async method to increment wait group
    - use `wg.Done()` to indicate the task is ended and decrease the counter
    - use `wg.Wait()` to wait till wait group has ended so the rest of the program can continue executing

- Example code
  - Example implemented [here](goroutines/goroutines.go)

> [!TIP]
> Check `sync.Mutex` and `Lock()-Unlock()` methods to prevent accessing the same memory at the same time