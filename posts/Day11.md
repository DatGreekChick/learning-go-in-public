# Day 11: JSON

**Date**: 27 June 2023

## What I did

Code is [here]

### Continued [Go Web Examples]

#### [JSON]

JSON parsing is pretty standard, though I personally love the way `JSON.parse`
handles it in JavaScript. In Go, you need to first define an encoder or
decoder, and then encode or decode, respectively.

I used [John Lurie] as the example because the web example was using a
narcissistic billionaire instead, and I much prefer John's presence! (If you
haven't heard of him, please take some moments to check him out. It will be
well worth your time; he's so talented!)

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day10 (main *)
🤓  curl -s -XPOST -d'{"firstName":"John","lastName":"Lurie","age":70}' http://localhost:8080/decode
John Lurie is 70 years old!
```

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day10 (main *)
🤓  curl -s http://localhost:8080/encode
{"firstName":"John","lastName":"Doe","age":25}
```

### Tomorrow is another day

I've been doing minimal Go learning during the week because I have a full time
job! I'm about to wrap up with these Go web examples, however. So I'm feeling
great about that progress.

After these examples have been completed, I'll do some reading on Go to ensure
I continue learning with positive behaviors (best practices). Then, I'll write
a CLI app!

[here]: ../code/day11
[go web examples]: https://gowebexamples.com/
[json]: https://gowebexamples.com/json/
[john lurie]: https://www.johnlurieart.com/
