# Day 5: Developing a RESTful API with Go and Gin

**Date**: 21 June 2023

## What I did

### [Tutorial]: RESTful API using Go and Gin

This is an interactive tutorial, but the code is all [here]. This is what I'm
setting out to accomplish with this tutorial, as per the description:

1. Design API endpoints
2. Create a folder for your code
3. Create the data
4. Write a handler to return all items
5. Write a handler to add a new item
6. Write a handler to return a specific item

The task is to build an API that provides access to a store selling vintage
recordings on vinyl. I'll be implementing the following endpoints:

#### `/albums`

- `GET` – Get a list of all albums, returned as JSON
- `POST` – Add a new album from request data sent as JSON

#### `/albums/:id`

- `GET` – Get an album by its ID, returning the album data as JSON

#### When to use `package main` over `package <name>`?

According to this tutorial:

> A standalone program (as opposed to a library) is always in package main.

## What is [Gin]?

From their GitHub repo:

> Gin is a HTTP web framework written in Go (Golang). It features a
> [Martini]-like API with much better performance -- up to 40 times faster. If
> you need smashing performance, get yourself some Gin.

Martini is no longer maintained, but it was something Go developers would use
for web apps and services. I appreciate the continuity of the naming between
these two packages!

## Thoughts

I would have liked to have more autonomy in writing the code. This was just
copy-paste. It was fun to get to use `gin`, though! Obviously, this is an
example of how to write a RESTful API. In the real world, we wouldn't be
storing the data in local memory - we'd be using a database. Further, `main`
should only execute the interesting parts, while everything else should live
outside of it.

I'll try shifting things around tomorrow!

### A note on the content for today

Today was a hectic day all around, but I'm proud that I still carved out time
to learn Go. Every bit counts! So if anyone out there is reading this, know
that small steps make a huge impact so long as you don't stop.

[tutorial]:
  https://shell.cloud.google.com/?walkthrough_tutorial_url=https%3A%2F%2Fraw.githubusercontent.com%2Fgolang%2Ftour%2Fmaster%2Ftutorial%2Fweb-service-gin.md&pli=1&show=ide&environment_deployment=ide
[here]: ../code/day05
[gin]: https://github.com/gin-gonic/gin
[martini]: https://github.com/go-martini/martini
