package database

import (
	"errors"
	"fmt"
	"go-alodokter/config/env"
	"sync"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlDb     *sqlx.DB
	lockMysqlDb sync.Mutex
)

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string, maxIdle int, MaxOpen int) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", descriptor)
	if err != nil {
		log.Errorf("error database %v", err)
		defer db.Close()
		return db, err
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(MaxOpen)
	db.SetConnMaxLifetime(time.Second * 10)
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Errorf("not connect database %v", err)
		return db, err
	}

	return db, err
}

// MysqlDB function for creating database connection
func MysqlDB() (mysqlMaster *sqlx.DB, err error) {

	lockMysqlDb.Lock()
	defer lockMysqlDb.Unlock()

	if mysqlMaster == nil {
		mysqlMaster, err = CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", env.Conf.DBUser,
			env.Conf.DBPass, env.Conf.DBHost, env.Conf.DBPort, env.Conf.DBName), env.Conf.MaxIdle, env.Conf.MaxOpenConn)
	}

	return mysqlMaster, err

}

// MySQL handle error parse
func Error(err error) (e error) {
	m, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	if m.Number == 1062 {
		return errors.New("duplicated data")
	}

	return err
}
