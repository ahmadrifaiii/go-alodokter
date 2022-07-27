package env

import (
	"os"

	"github.com/golobby/dotenv"
)

type ConfigurationEnvironment struct {
	ApplicationEnvironment
	DatabaseEnvironment
	AuthEnvironment
}

type ApplicationEnvironment struct {
	AppEnv   string `env:"APP_ENV" envDefault:"development"`
	HttpPort string `env:"PORT" envDefault:"8080"`
}

type DatabaseEnvironment struct {
	MaxIdle     int    `env:"DB_MAX_IDLE_CONN" envDefault:"10"`
	MaxOpenConn int    `env:"DB_MAX_OPEN_CONN" envDefault:"100"`
	DBName      string `env:"DB_NAME_MYSQL"`
	DBUser      string `env:"DB_USER_MYSQL"`
	DBPass      string `env:"DB_PASSWORD_MYSQL"`
	DBHost      string `env:"DB_HOST_MYSQL"`
	DBPort      string `env:"DB_PORT_MYSQL"`
}

type AuthEnvironment struct {
	TokenAge string `env:"TOKEN_AGE" envDefault:"24h30m"`
}

var Conf = ConfigurationEnvironment{}

func LoadEnv() {
	file, err := os.Open("system.env")
	if err != nil {
		panic(err.Error())
	}

	err = dotenv.NewDecoder(file).Decode(&Conf)
	if err != nil {
		panic(err.Error())
	}
}
