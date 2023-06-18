# Day 1: Getting Started

**Date**: 17 June 2023

## What I did

### New IDE, who dis?

Well, first things first: I installed [GoLand]. I'm a huge fan of JetBrains, so
trying it out to get Go language support is the best first step!

With this IDE, I gain access to [gofmt] out the box.

I also [downloaded Go]. If you don't see the below output, try refreshing your
console.

```bash
Elenis-MacBook-Pro: ~/learning-go-in-public (main *)
🤓  Go version
go version go1.20.5 darwin/arm64

# case insensitive! 😎
Elenis-MacBook-Pro: ~/learning-go-in-public (main *)
🤓  go version
go version go1.20.5 darwin/arm64
```

### Tour

Go documentation has a [Tour of Go] selected tutorial that's interactive. I'm
writing out my thoughts in real time as I _go_. (See what I did there? 🤓) The
reality might be different, but that's part of my learning process given my
experience!

Code snippets are lifted directly from the tutorial, though I occasionally add
my own commentary or change things ever so slightly.

#### Imports and exports

You can write multiple `import` statements on each line, but the preferred
style is to use something called a _factored import_ statement. This type of
statement just wraps all the imports with parens.

You can only export things with a capital letter! I like that there's a
standard for this, but it does seem restrictive. In all the dynamic languages
with which I've worked, this is not a requirement. Go is a statically typed,
compiled language, but it's clear this requirement is a stylistic choice from
the creators. I don't believe C++, for instance, has such a requirement.
Regardless, this leads me to the conclusion that Go is going to be heavier on
object-oriented programming (OOP).

If you've never tried OOP, I highly recommend it! I love a good mix of
functional programming (FP) and OOP!

```go
package main

// factored import
import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(math.pi) // this doesn't work
  fmt.Println(math.Pi) // note the capitalization!
}
```

Turns out you can also use "factored" variables (referenced below):

```go
var (
  ToBe bool = false
  MaxInt uint64 = 1 << 64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)
```

#### Typing

When I first learned to code, I didn't understand the importance of typing
because I was writing exclusively in JavaScript. While I have my opinions on
TypeScript, it's so much more helpful to know what types are allowed as
function arguments. Since [my journey to becoming a software engineer], I've
worked and dabbled in many languages - most of them typed.

I'll tell you it's game-changing and significantly helps with debugging. Go's
typing is reminiscent of Python's, but there's one shortcut that doesn't seem
helpful as a Go beginner: shortening types.

```go
package main

import "fmt"

// full types
func add(x int, y int) int {
  return x + y
}

// shortened types
func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
```

As long as the types match, you only need to explicitly state the type of the
last parameter. I'd imagine this becomes more complicated the more parameters
in a function, though it probably forces engineers to think more carefully on
the size of the functions. In other words, Go might encourage smaller functions
because of this type shorthand. In the wild, that likely plays out in several
ways (ignoring or configuring linters to name a couple).

#### Named return values

This stuff is reminiscent of Ruby's implicit returns. While Go still requires
an explicit `return` statement, it allows you to forego explicitly passing what
you're returning. The caveat is that you need to state what the values you are
returning as part of the function definition in named variables.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))
}
```

The tutorial states this isn't recommended the larger a function grows, and I
completely agree. Honestly, I'm not too hot on this option in general, but I'll
see how that opinion shifts (or doesn't) over time!

#### Variable declaration

The shortcuts continue! This makes me think a lot about JavaScript and
TypeScript, but Python even more.

Basically, you have many ways to declare variables:

- Implicit typing
- Explicit typing
- Factored assignment (shown above)
  - Can use any combination of the above (depending on the closure)

```go
package main

import "fmt"

func main() {
  // explicit typing
  var i, j int = 1, 2

  // implicit typing
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
```

Explicit typing involves you stating the type, while implicit typing requires
the short assignment construct (`:=`), is only available within function
closures, and assumes the type based on the value you assign to the variable.
This is also called type inference.

This reminds me of JavaScript / TypeScript because you can define many
variables on the same line. It reminds me of Python because you can do the same
while also assigning different values on the same line like you can in Go. I
used to love this when I was starting out, but then I started using linters and
formatters more heavily and decided that using one line per variable
declaration and / or assignment was more readable and maintainable.

Pick what works best for you!

```js
// an example of this in JavaScript
let first, second = true;
```

```python
# an example of this in Python
one, two = 1, 2
```

However, there is something I do like about not explicitly assigning a value to
a variable: defaults! Go will pick the default value if you don't assign
anything to a variable at first. For numeric types the default value is 0, for
`bool`s the default value is false, and for `string`s the default value is an
empty string `""`. While other languages set a default, they're usually some
version of `undefined`, `nil`, or `None`, etc. I'm not sure (yet!) if Go has a
concept of truthiness or falseness, but if it does, it's clear the default
values are falsy.

#### Looping

There are multiple ways to loop in Go, but they all start with `for`! That
keeps it super simple. While braces are required, parentheses are not.

```go
package main

import "fmt"

func main() {
  // a traditional style for-loop
  // takes init, condition, and post statements
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  // a while loop takes only the condition statement (init and post are optional)
  sum2 := 1
  for sum2 < 1000 {
    sum2 += sum2
  }

  for {} // infinite loop! 😱
}
```

I'll continue this tomorrow!

[goland]: https://www.jetbrains.com/go/
[gofmt]: https://pkg.go.dev/cmd/gofmt
[downloaded go]: https://go.dev/dl/
[tour of go]: https://go.dev/tour/
[my journey to becoming a software engineer]:
  https://codeburst.io/five-ways-becoming-a-software-engineer-made-me-a-wizard-de1060fc04d4
