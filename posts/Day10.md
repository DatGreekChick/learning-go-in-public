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

##### Update on above

It turns out I did need the `v2`. I'm not sure why my IDE was assuming the
import was unused, but this solved the issue above. Now I see this output:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  curl -i --cookie-jar cj --cookie cj localhost:4000/put
HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Set-Cookie: session=ps8U4xO6WzTm6oDokJCGIbjHXTAsr3k7fF5cYBt_5ic; Path=/; Expires=Thu, 29 Jun 2023 00:41:02 GMT; Max-Age=86400; HttpOnly; SameSite=Lax
Vary: Cookie
Date: Wed, 28 Jun 2023 00:41:01 GMT
Content-Length: 0

Elenis-MacBook-Pro : ~/learning-go-in-public (main *)
🤓  curl -i --cookie-jar cj --cookie cj localhost:4000/get
HTTP/1.1 200 OK
Vary: Cookie
Date: Wed, 28 Jun 2023 00:41:09 GMT
Content-Length: 21
Content-Type: text/plain; charset=utf-8

Hello from a session!
```

[here]: ../code/day10
[go web examples]: https://gowebexamples.com/
[sessions]: https://gowebexamples.com/sessions/
[`scs`]: https://github.com/alexedwards/scs
