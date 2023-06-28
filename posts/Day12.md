# Day 12: Completed Go Web Examples

**Date**: 28 June 2023

## What I did

Code is [here]

### Completed [Go Web Examples]

Most of the examples I elected to complete were recommending `gorilla` modules.
While that was fine because the repository is still publicly accessible, I
wanted to challenge myself a bit more and use repositories that were actively
maintained.

#### [Sockets]

This was probably the hardest example for me out of all the ones I completed
because I elected to not tie it to a frontend. I was able to get the server
running and work, but everything was effectively an error because I was passing
the `cURL` request incorrectly (as programmed). I only wanted a simple example
to test it out with Go, rather than something that worked end-to-end, so I was
okay with this approach though it wasn't particularly satisfying.

When running the server I get this output:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day12 (main *)
🤓  go run sockets.go
2023/06/28 19:48:13 Hello Go Web Examples, you're doing great!
2023/06/28 19:48:13 Error in wsutil.ReadClientData: EOF
```

Which was due to this `cURL` request:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day12 (main *)
🤓  curl -v --include --no-buffer \
   --header "Connection: Upgrade" \
   --header "Upgrade: websocket" \
   --header "Host: localhost:8080" \
   --header "Origin: https://localhost:8080" \
   http://localhost:8080/echo
```

I explicitly didn't pass a `"Sec-WebSocket-Key"` header as well as a version.

#### [Passwords]

There wasn't much in terms of this example as it was entirely copy-paste. It
just takes a predefined `password` and encrypts and decrypts it, ensuring both
match. This is pretty standard and something I've seen before.

When running the server I see this output:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day12 (main *)
🤓  sudo go run passwords.go
Password: secret
Hash:     $2a$14$S21ex7nYwyd9TQMmYiNfO.mub./mT4Y76BhaXeaEpZItv6rumEhPe
Match:    true
```

Tomorrow, I'll be focusing on reading material!

[here]: ../code/day12
[go web examples]: https://gowebexamples.com/
[sockets]: https://gowebexamples.com/websockets/
[passwords]: https://gowebexamples.com/password-hashing/
