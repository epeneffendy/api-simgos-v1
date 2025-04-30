package database

import (
	"database/sql"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

var DbConnectionSatu *sql.DB

func DbMigrateSatu(db2 *sql.DB) {
	migrations_satu := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_satu_migration"),
	}

	n, err := migrate.Exec(db2, "postgres", migrations_satu, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	DbConnectionSatu = db2
	fmt.Println("Applied", n, "migrations!")
}
