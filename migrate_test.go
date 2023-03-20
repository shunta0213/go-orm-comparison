package main_test

import (
	"testing"

	"github.com/shunta0213/go-orm-comparison/gorm"
	sqldb "github.com/shunta0213/go-orm-comparison/sql-db"
)

func BenchmarkSQLDB(b *testing.B) {
	db := sqldb.ConnectWithPQ()
	b.StartTimer()
	sqldb.Migrate(db)
	b.StopTimer()

	defer db.Close()
}

func BenchmarkGORM(b *testing.B) {
	db := gorm.ConnectDatabase()
	b.StartTimer()
	gorm.Migrate(db)
	b.StopTimer()
}
