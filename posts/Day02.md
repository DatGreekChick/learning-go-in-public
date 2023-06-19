# Day 2: Continue Tour of Go

**Date**: 18 June 2023

## What I did

### More [Tour of Go]

Again, all Go snippets are lifted from this tutorial sometimes with minor
modifications / commentary.

#### Control statements

Now this is pretty cool: you can define a variable in an `if` statement, and
it's still available within the `else` statement. The closure is wrapped around
the full control statements; outside of one a variable doesn't exist.

This isn't exclusive to Go, however. More recent versions of C++ also allow
similar behavior.

```go
package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  // note: defining v in the if and using it in the else
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}
```

#### Exercise: Loops and functions

See the [loops and functions exercise].

You can run Go programs by using the `go run` invocation:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  go run code/day02/exercise-loops-and-functions.go
1.4142135623730951
3
5
12
```

The `go run` command will compile and run the Go program in one shot. You can
run `go help` from the command line to see other helpful Go command-line tools.
You can also use your IDE to run the command, which I did first (see play
button):

![run from ide]

#### Other flow control statements

There are `switch` statements:

```go
package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd, plan9, windows...
    fmt.Printf("%s.\n", os)
  }
}
```

`break`s are built-in to `switch`es; no need to explicitly state them. In other
words, if you reach `case "linux":`, then you print out `"Linux."` and exit the
`switch`. You can also use `switch` statements in long `if-then-else`
statements by creating a `switch` without a condition (same as `switch true`).

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}
```

#### Defer, Panic, and Recover

There's a great [blog post] linked in [Tour of Go] to explain this much better
than I can!

#### Pointers

Like C++, Go has pointers, which hold a memory address of a value. You can
modify a value directly through a pointer via "dereferencing or indirecting."

```go
package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i         // point to i
  fmt.Println(*p) // read i through the pointer
  *p = 21         // set i through the pointer
  fmt.Println(i)  // see the new value of i

  p = &j         // point to j
  *p = *p / 37   // divide j through the pointer
  fmt.Println(j) // see the new value of j
}
```

#### Structs

A struct is a collection of fields, similar to an object-literal in JavaScript
or a `dict` in Python. Structs in C++ are the same as C++ classes, but their
default accessibility is public, whereas classes are inherently private. These
examples in the tour don't go into much detail, but my guess is we're able to
create struct methods in addition to having a collection of fields. This is how
we can define a struct in Go:

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1, 2})
}
```

You can access struct fields via dot notation:

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  v.X = 4 // like so
  fmt.Println(v.X) // and like so
  fmt.Println(v)
}
```

We can also make pointers that reference struct fields, and use the shorthand
dot notation instead of a more cumbersome and explicit way.

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v

  // we don't need to write (*p).X, but that's what's happening under the hood
  p.X = 1e9
  fmt.Println(v)
}
```

You can also create a struct literal using `Name:` syntax:

```go
package main

import "fmt"

type Vertex struct {
  X, Y int // using the shorthand type notation learned in Day 1
}

var (
  v1 = Vertex{1, 2}  // has type Vertex
  // using Name: notation for X, Y takes the default numerical value (0)
  v2 = Vertex{X: 1}  // Y:0 is implicit
  v3 = Vertex{}      // X:0 and Y:0
  p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
```

#### Arrays and slices

Go arrays are mutable no matter how you _slice_ and dice them. The way you
instantiate an array is to use `var` followed by the name of the array, bracket
notation indicating the size of the array, and the type of array.

```go
// this is an integer array of size 10
var a [10]int
```

To make a slice of the array you use `a[low : high]`, where `low` is the lowest
index and `high` is the highest index of the array you want to slice. The
`high` value is exclusive, while the `low` value is inclusive.

```go
// a slice of array a from index 1 to 3 inclusive (1 to 4 exclusive)
a[1:4]
```

Slices only take a reference of an already defined array, and if you change a
slice's value at some index, the original array at that index will also change.
This is different from a language like JavaScript, which takes a copy and
returns a new array with `Array.prototype.slice`.

Here's what happens in JavaScript:

```js
const arr = [1, 2, 3];

const slice = arr.slice(); // [1, 2, 3]
slice[1] = 4;

console.log(arr); // still [1, 2, 3]
console.log(slice); // [1, 4, 3]
```

Now, here's what happens in Go:

```go
package main

import "fmt"

func main() {
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names) // [John Paul George Ringo]

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b) // [John Paul] [Paul George]

  // this not only changes b, it also changes a (another slice of names) and names
  // at the relevant indices
  b[0] = "XXX"
  fmt.Println(a, b) // [John XXX] [XXX George]
  fmt.Println(names) // [John XXX George Ringo]
}
```

The behavior in Go is similar to the behavior in C++. Here's my favorite part
of slicing in Go, though, which is similar to how we can slice or refer to
indices in Python:

```go
// you don't need to explicitly pass the size -- it will figure that out
// from the number of values you pass at instantiation
s := []int{2, 3, 5, 7, 11, 13}

// explicitly slicing from index 1 to index 4 exclusive (s[1] to s[3])
s = s[1:4]
fmt.Println(s) // [3, 5, 7]

// slice from 0 to 2 exclusive (s[0] to s[1])
s = s[:2]
fmt.Println(s) // [3, 5]

// slice from 1 to length of array
s = s[1:]
fmt.Println(s) // [5]

// slice of the whole array
s = s[:]
fmt.Println(s) // [5]
```

Slices have both length and capacity.

```go
package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it zero length.
  s = s[:0]
  printSlice(s)

  // Extend its length.
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

If you try to extend the length of a slice beyond its capacity, you'll see the
following:

```go
  // Extend its length.
  s = s[:8]
  printSlice(s)
```

```bash
panic: runtime error: slice bounds out of range [:8] with capacity 6

goroutine 1 [running]:
main.main()
  /tmp/sandbox2661788614/prog.go:14 +0x86
```

But what if you need something more dynamic or want to append? You can use
`make`, a built-in Go function to create a dynamically-sized array (similar to
a C++ vector). `make` allocates an array full of zeroes (a zeroed array) and
returns a slice that refers to that array. The arguments you can pass to `make`
are: (1) array type like `[]int`, (2) length of array, and (3) an optional
capacity. If you don't pass in a capacity, `make` will assume it's equal to
length of the array.

```go
package main

import "fmt"

func main() {
  a := make([]int, 5)
  printSlice("a:", a) // a: len=5 cap=5 [0 0 0 0 0]

  b := make([]int, 0, 5)
  printSlice("b:", b) // b: len=0 cap=5 []

  c := b[:2]
  printSlice("c:", c) // c: len=2 cap=5 [0 0]

  d := c[2:5]
  printSlice("d:", d) // d: len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
```

For appending, you can use Go's built-in `append` method, which takes a slice
of type `T`, and any values you wish to append to the slice.

```go
package main

import "fmt"

func main() {
  var s []int
  printSlice(s) // len=0 cap=0 []

  // append works on nil slices
  s = append(s, 0)
  printSlice(s) // len=1 cap=1 [0]

  // the slice grows as needed
  s = append(s, 1)
  printSlice(s) // len=2 cap=2 [0 1]

  // we can add more than one element at a time
  s = append(s, 2, 3, 4)
  printSlice(s) // len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

My gut reaction here is mixed... I like the fact that you can pass in a lot of
values to the `append` function in order to add additional values to the slice.
I don't like that you're required to assign it to some value in order for the
appended values to appear in the slice. When I tried `append(s, 5)` I got this
error:

```bash
./prog.go:19:2: append(s, 5) (value of type []int) is not used
```

What is awesome about this is that anything that isn't used is considered an
error! That's great for code quality. I wonder how far that extends throughout
the language.

#### Ranges

You can use a `range` form of a loop to iterate over a slice or map. When you
range through a slice, two values are returned for each iteration:

1. the index
2. a copy of the element at that index

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}

// output:
// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 2**4 = 16
// 2**5 = 32
// 2**6 = 64
// 2**7 = 128
```

This again reminds me of [a Python `range`]:

```python
# will loop 5 times
for i in range(5):
  print(i)
```

Other ways to define a range in Go:

```go
for i, _ := range pow
for _, value := range pow
for i := range pow // this is most similar to the Python example
```

#### [Slice exercise]

Okay, we've got our first package to install! This was trial and error... I
didn't Google anything for this.

```bash
# 'go mod' provides access to operations on modules
# if you want to add / remove / upgrade / downgrade a dependency, you need to ues 'go get'

# 'go mod init' initializes a new module in the current directory
Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  go mod init learning-go-in-public
go: creating new go.mod: module learning-go-in-public
go: to add module requirements and sums:
        go mod tidy
```

Since I had a module that I needed for this exercise that wasn't part of the Go
standard library, I ran the recommended command `go mod tidy`.

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  go mod tidy
go: finding module for package golang.org/x/tour/pic
go: found golang.org/x/tour/pic in golang.org/x/tour v0.1.0
```

After running this, I got access to `go.mod` as a top-level directory within
`learning-go-in-public` and the following appeared in `go.mod/go.sum:

```go.sum
golang.org/x/tour v0.1.0 h1:OWzbINRoGf1wwBhKdFDpYwM88NM0d1SL/Nj6PagS6YE=
golang.org/x/tour v0.1.0/go.mod h1:DUZC6G8mR1AXgXy73r8qt/G5RsefKIlSj6jBMc8b9Wc=
```

I found this especially awesome and user-friendly!

After writing up the code and running it:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  go run code/day02/exercise-slices.go
IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAACK0lEQVR42uzTMQ1DIQBF0ZemEpgRgTZkoA0RzAw4aNO0YxX8f85E4E0k93nOq/ckKSV7p7Ws9Tl8b2rNnL+nJGPE2PhK40fgxgSAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAEAAIAAYAAQAAgABAACAAEAAIAAYAAQAAgABAACAAEAAIAAYAAQAAgABAACAAEAAIAAYAAQAAgABAACAAEAAIAAYAAQAAgABAACAAEAAIAAYAAQAAgABAACAAEAAIAAYAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAASAAHwBAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgABAACAAEAAIAAQAAgABgAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAAIAAQAAgABAACAAGAAEAA8M87AAD//9Ei0OJgFSnhAAAAAElFTkSuQmCC
```

That doesn't really indicate what the image should look like, but the [Tour of
Go] interactive editor outputs this bluescale image:

![bluescale image]

Alright, that's all for today folks!

[tour of go]: https://go.dev/tour
[loops and functions exercise]: ../code/day02/exercise-loops-and-functions.go
[run from ide]: ../img/day02/running-go-programs.png
[blog post]: https://go.dev/blog/defer-panic-and-recover
[a python `range`]: https://docs.python.org/3/library/stdtypes.html#ranges
[slice exercise]: ../code/day02/exercise-slices.go
[bluescale image]: ../img/day02/bluescale.png
