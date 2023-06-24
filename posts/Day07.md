# Day 7: Read How to Write Go Code & Started Go Web Examples

**Date**: 23 June 2023

## What I did

### [How to Write Go Code]

I read this article from Go's documentation that had been linked throughout the
past few days I've been learning Go. Today, I finally make the time to read
through it!

#### Code organization

Code organization is super important, and it's imperative to understand how the
language you're working with expects you to organize everything.

I've already learned the distinction between _packages_ and _modules_ in Go,
but I've learned a bit more from this article.

There are four key terms to consider when structuring Go code:

1. Program
2. Package
3. Module
4. Repository

Go **programs** (AKA: source files) are organized into **packages**.

> A package is a collection of source files in the same directory that are
> compiled together.

Anything that is defined in one source file within a package is available to
another within the same package. That includes functions, types, variables, and
constants.

> A module is a collection of related Go packages that are released together.

**Modules** are a series of packages that are bundled together to create a
**program**.

> A repository contains one or more modules.

While a **repository** can contain more than one module, it typically doesn't.
Regardless, modules can be found at the root of a repository.

There are of course more things to remember when developing in Go, but the
structure lies in those four categories. While it isn't required to publish
your code to a remote repository, _it is_ a good habit to structure your code
assuming that some day you might publish it publicly.

#### Imports and paths

When you create a module, you use the `go mod init` command. The name you
provide to this command declares the module path and stores it in `go.mod`.

> The module path is the import path prefix for all packages within the module.

As noted, a module is a collection of packages, so a module contains the
packages in the directory tied to its `go.mod` file. It also includes
subdirectories of that directory until Go finds the next subdirectory with a
`go.mod` file, if present.

The module path is not only the import path prefix for its packages, however.
It's also an indicator of where go should search to download it.

> An import path is a string used to import a package.

An import path is the module path joined with its subdirectory within the
module. If I were to create a Go module on GitHub under my GitHub user account,
the URL for that GitHub repository would look something like this:
github.com/DatGreekChick/going. Let's say that module had a package in the
directory `gone`. What would be the package's import path?

```
github.com/DatGreekChick/going/gone
```

Note how this is different from any packages in the Go standard library. Those
don't have a module path prefix and can be imported directly (like `fmt`).

#### Writing a program

The first thing that always needs to appear on the very first line of a program
is the `package <name>`. If the program is meant to be executable, it needs to
be in `package main`.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
```

#### Installing packages

To install packages, you need the `go install` command:

```bash
$ go install example/user/hello
```

You don't need to pass in the path to the `go install` command if you're
installing from within the relevant working directory; it will apply the
context of the module containing that current working directory.

To further illustrate that point, the following three options for the command
are exactly the same:

```bash
$ go install example/user/hello
$ go install .
$ go install
```

This may fail if the working directory is not within the `example/user/hello`
module, however.

#### Testing

Go comes with a lightweight test framework: the `go test` command and the
`testing` package. Go files should be styled in lowercase snake case (example:
`my_go_program.go`) with the `.go` extension. Go _test_ files follow their
source file counterpart, but contain the additional `_test`. This naming is
required. Using our example, the test file would be `my_go_program_test.go`.

Within a test file, tests consist of functions named like so: `TestXXX`. All
test functions have the following signature: `func (t *testing.T)`

### [Go Web Examples]

Code for that is [here]. It appears as though these examples somewhat build
upon each other, and as the days go by I'll bring in what I've learned to the
examples.

One thing that this tutorial requests is to use `gorilla/mux` as the router.
It's still publicly available to use, but it's been archived and is no longer
maintained. Instead of using something that is no longer really meant to be
used, I'm going to switch to `gin`.

I've stopped at building the router and have written rudimentary code to handle
a GET request to `/books/:title/page/:page`. The next step is to create a MySQL
database. I use Postgres at work, so I'll be using that instead! I'll continue
that tomorrow! :smile:

[how to write go code]: https://go.dev/doc/code
[go web examples]: https://gowebexamples.com/
[here]: ../code/day07
