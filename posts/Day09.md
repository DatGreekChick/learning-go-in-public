# Day 9: Middleware

**Date**: 25 June 2023

## What I did

Code is [here]!

### Continued [Go Web Examples]

#### [Basic Middleware]

I mainly lifted from the example, but I extracted a function out to log the
request.

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day09 (main *)
🤓  go run .
2023/06/25 17:40:04 /bar
2023/06/25 17:40:08 /foo
```

#### [Advanced Middleware]

Ooh, a new argument notation!

```go
func someFunc(middlewares ...Middleware) {}
```

The official definition from Go's language specification on [operators]:

> The final incoming parameter in a function signature may have a type prefixed
> with `...`. A function with such a parameter is called variadic and may be
> invoked with zero or more arguments for that parameter.

This is a _variadic_ function, and this operator can only appear for a
parameter when it's the last one of a function. It can incorporate zero or more
arguments for that parameter.

For this example, all I added was error handling. No matter the route, we get
the same `"hello world"` to appear.

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day09 (main *)
🤓  go run advanced-middleware.go
2023/06/25 17:51:06 /foo 71.875µs
2023/06/25 17:51:11 /hello 14.375µs
2023/06/25 17:51:24 / 19.208µs

# different terminal tab
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day09 (main *)
🤓  curl -s http://localhost:8080/
hello world

Elenis-MacBook-Pro : ~/learning-go-in-public/code/day09 (main *)
🤓  curl -s -XPOST http://localhost:8080/
Bad Request
```

That's all for today folks! See you tomorrow :muscle:

[here]: ../code/day09
[go web examples]: https://gowebexamples.com/
[basic middleware]: https://gowebexamples.com/basic-middleware/
[advanced middleware]: https://gowebexamples.com/advanced-middleware/
[operators]: https://go.dev/ref/spec#Operators
