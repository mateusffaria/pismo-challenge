package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	migratePG "github.com/golang-migrate/migrate/v4/database/postgres"
)

type DBConn struct {
	SSL      bool
	Port     int
	Host     string
	User     string
	Password string
	Dbname   string
}

func NewDatabaseConnection(dbc DBConn) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Dbname)
	if dbc.SSL {
		sqlInfo += " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("open database connection failed")
	}

	return db
}

func RunMigrations(db *sql.DB, migrationsPath string) {
	driver, err := migratePG.WithInstance(db, &migratePG.Config{})
	if err != nil {
		log.Fatal("migration WithInstance failed: ", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("could not get working directory: ", err)
	}

	if migrationsPath == "" {
		migrationsPath = "file://" + filepath.Join(wd, "db/migrations")
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatal("migration NewWithDatabaseInstance failed: ", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no new migrations to apply.")
			return
		}

		log.Fatal("migration Up failed: ", err)
	}
}
