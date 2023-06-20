# Day 4: [Create a Go Module]

**Date**: 20 June 2023

## What I did

### Tutorial: [Create a Go Module]

This is a great starting point after walking through the key language features.
In this tutorial I'll be doing the following:

1. Write a small module that can be imported in other modules
2. Import and use the module
3. Handle errors
4. Return a random greeting using slices (Go's dynamically-sized arrays)
5. Return greetings for multiple people
6. Add a test!
7. Compile and install the application

I've done bits and bobs from this list in days 2 and 3, but this will take it a
step further especially with unit testing. I'm super excited about that.
Testing is fundamental to any application, and I equate it at the same level of
importance as source code.

The way Go code is structured is via packages. Think of anything that requires
`package main` at the top, for instance. These packages are then grouped into
modules. A module will list out required dependencies, which could consist of
other modules and the Go version necessary to run the code properly.

This happens to be the inverse in terms of nomenclature of Python. In Python, a
package is a collection of modules. Semantics!

On day 1 I learned that you can only export things with capital letters:

> In Go, a function whose name starts with a capital letter can be called by a
> function not in the same package. This is known in Go as an exported name.

This likely allows for some private functionality assuming the function is
lowercase.

#### Import structure

It's not explicitly mentioned in the tutorial, but I noticed a space between
the two imports. I really love to group imports by sort order (built into
`gofmt`) as well as have a hierarchy depending on the language.

```go
package main

// notice the space! fmt is part of Go's standard library and comes first
import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Eleni")
	fmt.Println(message)
}
```

When writing C++, you should always use the `#include` directive starting with
the header file that matches the source file. Say you have `hello.h` as your
header file, and `hello.cpp` as your source file. The first included file
should be the header:

```cpp
// hello.cpp
#include <hello.h> // or "hello.h" depending on where the preprocessor searches

// any other required inclusions

// the code
```

After the first header is included, the order should generally also be sorted
(using `clang-format` is the best for this!). Then you generally follow the
hierarchy of within your project (code you or other teams have written), then
you include open source packages.

The opposite happens in Python and JavaScript.

To be clear, you don't _need_ to do these things. But it keeps your code
consistent especially if you use formatters and linters. For C++ it at least
allows you to fail compilation early should something be wrong with your header
file or the other dependencies. You might get weird, unhelpful errors
otherwise.

#### Overriding module location

```bash
$ go mod edit -replace example.com/greetings=../greetings
```

This command allows you to import the module using the string
`"example.com/greetings"` but will point to the directory where we defined it
rather than looking online for a package published to this location.

When you do this, you get this output in `go.mod:

```go.mod
module example.com/hello

go 1.20

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

Cool 😎

The module now has this default semantic-versioning output. If this were a
published package, we'd get the latest version or be able to explicitly set
which version we wanted or required.

#### Private functions

Turns out my assumption earlier was correct! When you create functions starting
with a lowercase letter, the function is only accessible within the package where
it is defined.

[create a go module]: https://go.dev/doc/tutorial/create-module
