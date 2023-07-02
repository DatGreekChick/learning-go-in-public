# Day 15: Continued Reading Effective Go

**Date**: 1 July 2023

## What I Did

### [Effective Go]

Sections / subsections read:

- Switch
- Type Switch
- Functions
  - Multiple return values
  - Named result parameters
  - Defer

#### Type Switch

It's so wild to me that you can define variables within switches (and in other
areas!) in Go! It's so flexible!!

#### Defer

I was struggling with understanding `defer` and its apparent asynchronicity.
This guide made it so much clearer:

> `defer` schedules a function call (the _deferred_ function) to be run
> immediately before the function executing the `defer` returns

If that still doesn't make sense, the example they use is most helpful:

```go
// Contents returns the file's contents as a string
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }

    defer f.Close()  // f.Close will run when we're finished

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later

        if err != nil {
            if err == io.EOF {
                break
            }

            return "", err  // f will be closed if we return here
        }
    }

    return string(result), nil // f will be closed if we return here
}
```

They mention it's good practice to have a `defer` statement so close to opening
the file so that you don't forget to close it. Really great practice.

This example was further helpful because of the _comments_! They all mark when
the file will be closed at each `return`, which means that the `defer`
statement executes whatever we've been deferring and then returns what we're
expecting at each point.

Remember as well that `defer` statements run in LIFO order -- like a stack.

[effective go]: https://go.dev/doc/effective_go
