package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
}

func main() {
	ctx := context.Background()

	// Open a PostgreSQL database
	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it
	db := bun.NewDB(pgDb, pgdialect.New())

	// Print all queries to stdout
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	res, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)
	if err != nil {
		fmt.Errorf("there was an error creating the users table: %v", err)
	}

	fmt.Printf("Successfully created users table with res: %v\n", res)

	user := User{Name: "Eleni", CreatedAt: time.Now()}
	userRes, userErr := db.NewInsert().Model(&user).Exec(ctx)
	if userErr != nil {
		fmt.Errorf("there was an error adding the user: %v", userErr)
	}

	fmt.Printf("Successfully created user with res: %v\n", userRes)
}
