package main

import (
	"database/sql"
	"fmt"
	"go-alodokter/config/env"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
)

func main() {
	env.LoadEnv()

	condb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		env.Conf.DBUser,
		env.Conf.DBPass,
		env.Conf.DBHost,
		env.Conf.DBPort,
		env.Conf.DBName)

	db, err := sql.Open("mysql", condb)
	if err != nil {
		log.Errorf("error database %v", err)
		defer db.Close()
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Errorf("not connect database %v", err)
		os.Exit(1)
	}

	fmt.Println("connect database success")
	fmt.Println("start migrate")

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Errorf("error database %v", err)
		defer db.Close()
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "./migration_files"),
		"mysql",
		driver,
	)
	if err != nil {
		log.Errorf("error migration %v", err)
		defer db.Close()
		os.Exit(1)
	}

	m.Steps(2)

	src, dbe := m.Close()
	if dbe != nil || src != nil {
		log.Errorf("error migration %v %v", src, dbe)
		defer db.Close()
		os.Exit(1)
	}

	fmt.Println("migrate success")
}
