package usecase

import (
	"errors"
	"go-alodokter/config"
	"go-alodokter/model"
	"go-alodokter/module/v1/order/repo"
	"go-alodokter/utl/password"
	"strings"
	"time"
)

func OrderList(conf config.Configuration) (users []model.User, err error) {
	db := conf.MysqlDB
	return repo.GetOrderList(db)
}

// get order detail
func UserDetail(conf config.Configuration, userId int) (users model.User, err error) {
	db := conf.MysqlDB
	return repo.GetUserDetail(db, userId)
}

// register new user
func UserRegister(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Password) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	hash, err := password.Encrypt(usr.Password)
	if err != nil {
		return *usr, err
	}

	payload.Password = string(hash)
	payload.Name = usr.Name
	payload.Email = usr.Email

	_, err = repo.RegisterUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}

// update user
func UserUpdate(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	payload.Name = usr.Name
	payload.Email = usr.Email
	payload.Id = usr.Id

	_, err = repo.UpdateUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}

// delete user
func UserDelete(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	payload.Id = usr.Id
	payload.Deleted = &now

	_, err = repo.DeleteUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}
