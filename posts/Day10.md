# Day 10: Sessions

**Date**: 26 June 2023

## What I did

Code is [here]... Finally hit double digits!

### Continued [Go Web Examples]

#### [Sessions]

This example was once again recommending a repository that's now been archived.
I decided to look for replacements and found [`scs`]. I used the example as
provided in the `README` but had to take out the version in the import `v2`,
otherwise my IDE wouldn't recognize the package.

I'm also not convinced I properly installed the package despite seeing
`go mod tidy` output because when I try to execute the code I see this:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day10 (main *)
💀  go run sessions.go
# command-line-arguments
./sessions.go:13:24: undefined: scs.New
./sessions.go:26:2: undefined: sessionManager
./sessions.go:32:9: undefined: sessionManager
```

This doesn't make too much sense to me right now, but I'm calling it a night
and will debug this more tomorrow to figure out the issue and hopefully wrap up
the rest of these examples.

[here]: ../code/day10
[go web examples]: https://gowebexamples.com/
[sessions]: https://gowebexamples.com/sessions/
[`scs`]: https://github.com/alexedwards/scs
