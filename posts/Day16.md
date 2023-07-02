# Day 16: More Effective Go

**Date**: 2 July 2023

## What I Did

### [Effective Go]

Sections / subsections read:

- Data
  - Allocation with `new`
  - Constructors and composite literals
  - Allocation with `make`
  - Arrays
  - Slices

#### Allocation with `new`

Using `new` is one of two built-in allocation primitives in Go. It allows
programmers to initialize types with their zero-values and returns the address
(a pointer) where that item is stored. This is helpful when you just want to
get up and running with a new instance of type `T`, but initializing a variable
with `var` of type `T` gets you almost the same result. `var` won't return a
pointer, but it's also zeroed.

I'm not sure why it would be more helpful to use `new` over `var` unless a
function signature is expecting a pointer; then you can initialize something
with `new` in one function and pass it to another. But that's only one use
case.

I was curious why Go developers use `new`, and it's basically something that
you don't use often but is there when you need it. See more discussion from
[this StackOverflow post].

#### Allocation with `make`

`make` is the other built-in allocation primitive in Go. It is different from
`new` because it returns an initialized (not zero) value of type `T` instead of
type `*T`. In other words, it returns the value instead of the reference. This
function is also exclusive to slices, maps, and channels. It can't be used to
initialize anything else.

Here are some examples that note the differences between `make` and `new`:

```go
// allocates slice structure; *p == nil; rarely useful
var p *[]int = new([]int)
// the slice v now refers to a new array of 100 ints
var v  []int = make([]int, 100)

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

[effective go]: https://go.dev/doc/effective_go
[this stackoverflow post]:
  https://softwareengineering.stackexchange.com/questions/210399/why-is-there-a-new-in-go
