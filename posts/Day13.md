# Day 13: Read Middleware Post

**Date**: 29 June 2023

## What I did

I kept it super light today and focused on reading [this post]. The post
discusses using an `Adapter` type, which follows the [pattern by the same
name]. The essence of that pattern is it takes an object, and creates a new
object of the same type.

This is one snippet example:

```go
func Notify() Adapter {
  return func(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      log.Println("before")
      defer log.Println("after")
      h.ServeHTTP(w, r)
    }
  }
}
```

The post goes through ways to make this cleaner, even suggesting a function
that does the adapting for us instead. To maintain order, however, it should
reverse through them in the adapt function instead of ranging through them.

So how does this tie into middleware? The example of this post is using the
adapter technique in order to pass middleware through to `http.Handle`. But of
course, this is extendable to any middleware!

[this post]:
  https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
[pattern by the same name]: https://refactoring.guru/design-patterns/adapter
