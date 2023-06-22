# Day 6: Splitting Up the RESTful API

**Date**: 22 June 2023

## What I did

[Yesterday] I worked on an interactive tutorial from Go's documentation. While
it exposed me to a few new things, it was lacking trial and error because the
entire tutorial was basically copy-paste. I wanted to build upon the knowledge
I've accumulated the past few days and bring it all together. So, I extracted
some of the setup code into an `api` package, and imported that package into
the `main` package named `albums.go`. I also tested the router a bit. You can
see my work [here].

## What I learned

### Packages

I had been exposed to this already, but this exercise really solidified it for
me: whatever you name the package is how you need to import your code. Also,
you need a separate `main` package in order to run and execute your code.

I was struggling in the beginning trying to name all the packages `albums`.
Once I named the `api` package and the `main` package, which imported the `api`
package, the code started compiling. It was so simple and in front of me, but
repetition is the most helpful when learning something new!

### Testing

There are great testing packages out there for Go. It's already built-in to the
language, but you can get mocking _for free_ from something like
`"net/http/httptest"`. I love how easy that was!

I also was happily surprised at the ability to use a private function in my
test. Any other language I've used that has private methods prevents use of
those methods explicitly. This private function was `setupRouter`. This
function didn't exist in the tutorial, but it was in a code snippet example for
testing, so I applied it to my use cases. It also made the code much cleaner.
For a real-world router, this is exactly what you'd do: split things up into
bite-sized chunks to make changes that much cleaner and easier to reason about.

I would have also really loved to have tested the actual output, but that could
have fit better in `api/albums_api_test.go`, something on which I didn't spend
time. What I focused on was more of an integration test: can I make a call to
an API endpoint and see the responses I expect?

Tomorrow is another day!

[yesterday]: ./Day05.md
[here]: ../code/day06
