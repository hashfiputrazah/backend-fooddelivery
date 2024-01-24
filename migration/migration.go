package migration

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func connectDB() *sql.DB {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func MigrateUp() {
	driver, err := postgres.WithInstance(connectDB(), &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/file",
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		v, d, _ := m.Version()
		if d {
			if v == 1 {
				m.Drop()
			} else {
				m.Force(int(v) - 1)
				panic(err)
			}
		}
	}

	fmt.Println("Migration run successfully")
}
