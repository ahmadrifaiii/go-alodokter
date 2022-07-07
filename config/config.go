package config

import "github.com/jmoiron/sqlx"

type Configuration struct {
	MysqlDB *sqlx.DB
}
