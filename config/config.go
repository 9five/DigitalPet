package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func init() {
	DB = newDBConfig()
}

func newDBConfig() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=DigitalPet port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "DigitalPet."},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
