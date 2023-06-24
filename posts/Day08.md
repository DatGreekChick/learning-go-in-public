# Day 8: Set up a Postgres Database

**Date**: 24 June 2023

## What I did

### Downloaded [Postgres]

It's been awhile since I had Postgres on my personal laptop. Last time I did
this was in 2017 at my coding bootcamp! At work, we don't use database GUIs,
but we have the option to connect to the database via our JetBrains IDEs. This
is what I've done here as well. I really like the simplicity of (most)
everything being accessible from my IDE, and viewing records this way is pretty
nice! The thing that's lacking? Tagging query performance and other rich
features. For the purposes of learning Go, it's sufficient for me to not have
all the available features.

### Used an [ORM]

At my former company, we used an open-sourced SQL-like database called
[`comdb2`]. It had no ORM associated with it, so every query was raw SQL
:scream:

Now that I've used [`ActiveRecord`], I don't want to go back to raw SQL except
when necessary. For Go, I found a few different options, but most of them
seemed to be outdated or soon-to-be outdated. As such, I've landed on [`bun`].
According to their website, `bun` is a

> Lightweight Golang ORM for PostgreSQL, MySQL, MSSQL, and SQLite.

### Set up a database in Go

See `bun`'s [quickstart] guide, but we'll need a few more modules:

- `context`
- `database/sql`
- `github.com/uptrace/bun/*`

Here's the code snippet provided in the quickstart guide:

```go
package main

import (
  "context"
  "database/sql"
  "fmt"

  "github.com/uptrace/bun"
  "github.com/uptrace/bun/dialect/pgdialect"
  "github.com/uptrace/bun/driver/pgdriver"
  "github.com/uptrace/bun/extra/bundebug"
)

func main() {
  ctx := context.Background()

  // Open a PostgreSQL database
  dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
  pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

  // Create a Bun db on top of it
  db := bun.NewDB(pgDb, pgdialect.New())

  // Print all queries to stdout
  db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

  var rnd float64

  // Select a random number
  if err := db.NewSelect().ColumnExpr("random()").Scan(ctx, &rnd); err != nil {
    panic(err)
  }

  fmt.Println(rnd)
}
```

### Continued [Go Web Examples]

If it wasn't already obvious, I focused on bringing up a database today! Code
for that is [here]. Since I'm heavily involved with migrations at work, I
elected to keep this simple and in application code.

#### Output

After successfully bringing up the database, creating a `users` table, and
inserting one record into the table, I see this on the command line:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day08 (main *)
🤓  go run .
[bun]  14:44:16.755   CREATE TABLE         26.199ms  CREATE TABLE "users" ("id" BIGSERIAL NOT NULL, "name" VARCHAR NOT NULL, "created_at" TIMESTAMPTZ NOT NULL, PRIMARY KEY ("id"))
Successfully created users table with res: {0x14000216000 0}
[bun]  14:44:16.756   INSERT                  711µs  INSERT INTO "users" ("id", "name", "created_at") VALUES (DEFAULT, 'Eleni', '2023-06-24 18:44:16.755573+00:00') RETURNING "id"
Successfully created user with res: 1
```

I see this using the CLI for Postgres:

```bash
Elenis-MacBook-Pro : ~/learning-go-in-public/code/day08 (main *)
🤓  psql -p5432 "postgres"
psql (15.3)
Type "help" for help.

postgres=# \d users
                                       Table "public.users"
   Column   |           Type           | Collation | Nullable |              Default
------------+--------------------------+-----------+----------+-----------------------------------
 id         | bigint                   |           | not null | nextval('users_id_seq'::regclass)
 name       | character varying        |           | not null |
 created_at | timestamp with time zone |           | not null |
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)

postgres=# select * from users limit 1;
 id | name  |          created_at
----+-------+-------------------------------
  1 | Eleni | 2023-06-24 14:44:16.755573-04
(1 row)
```

...and I see the same data from my IDE!

![postgres database]

## Looking Ahead

I've perused a bunch of the tutorials left in Go Web Examples, and since there
are HTML template examples, I'm deciding to skip those. I'm not interested
using Go for frontend code because I much prefer React for web development over
templates that are hard to read and outdated ways to develop web apps. I'll be
skipping the following examples:

- Templates
- Assets and Files
- Forms

But there are still 6 more examples!

[postgres]: https://www.postgresql.org/
[orm]:
  https://www.freecodecamp.org/news/what-is-an-orm-the-meaning-of-object-relational-mapping-database-tools/
[`comdb2`]: https://bloomberg.github.io/comdb2/overview_home.html
[`activerecord`]: https://guides.rubyonrails.org/active_record_basics.html
[`bun`]: https://bun.uptrace.dev/
[quickstart]: https://bun.uptrace.dev/#quickstart
[go web examples]: https://gowebexamples.com/
[here]: ../code/day08
[postgres database]: ../img/day08/postgres-database.png
