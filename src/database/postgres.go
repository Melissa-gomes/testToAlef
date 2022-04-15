package database

import (
	"fmt"

	"github.com/Melissa-gomes/testGCP/src/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	DB *gorm.DB
}

func NewPostgresDatabase() *gorm.DB {
	env := config.LoadEnv()
	DbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		env.POSTGRES_USERNAME,
		env.POSTGRES_PASSWORD,
		env.POSTGRES_HOST,
		env.POSTGRES_PORT,
		env.POSTGRES_DATABASE,
	)

	db, err := gorm.Open(postgres.Open(DbConfig), &gorm.Config{})
	if err != nil {
		panic(errors.Wrap(err, "Error to connect at Database"))
	}
	return db
}
