package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/utils"
)

func NewPostgresDSN(env *Env) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresName,
		env.PostgresHost,
		env.PostgresPort,
	)
}

func NewPostgresConnection(env *Env) *gorm.DB {
	dsn := NewPostgresDSN(env)
	postgresConfig := postgres.Open(dsn)
	gormConfig := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}

	db, err := gorm.Open(postgresConfig, &gormConfig)
	utils.ErrorPanic(err)

	db.Table("product").AutoMigrate(&model.Product{})

	return db
}
