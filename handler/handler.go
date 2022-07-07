package handler

import (
	"go-alodokter/config"
	"go-alodokter/config/database"
	"go-alodokter/module/v1/auth"
	"go-alodokter/module/v1/user"
	authMid "go-alodokter/utl/middleware/auth"
)

type Service struct {
	MiddlewareAuth *authMid.Handle
	AuthModule     *auth.Module
	UserModule     *user.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB: MySQLConnection,
	}

	// set service modular
	middlewareAuth := authMid.InitAuthMiddleware(config)
	moduleAuth := auth.InitModule(config)
	moduleUser := user.InitModule(config)

	return &Service{
		AuthModule:     moduleAuth,
		UserModule:     moduleUser,
		MiddlewareAuth: middlewareAuth,
	}
}
