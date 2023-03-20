package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func Create(db *gorm.DB) {
	db.Create(&User{Name: "test", Age: 2})
}

func Read(db *gorm.DB) User {
	var user User
	db.First(&user, 1)
	return user
}
