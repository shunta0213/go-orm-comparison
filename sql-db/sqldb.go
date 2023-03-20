package sqldb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithPQ() *sql.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name varchar NOT NULL,
		age int32 NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT(now())
	)
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
