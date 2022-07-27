package usecase

import (
	"errors"
	"go-alodokter/config"
	"go-alodokter/model"
	"go-alodokter/module/v1/product/repo"

	"strings"
	"time"
)

func ProductList(conf config.Configuration) (users []model.Product, err error) {
	db := conf.MysqlDB
	r, err := repo.GetProductList(db)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return nil, errors.New("data not found")
	}

	return repo.GetProductList(db)
}

// get product detail
func ProductDetail(conf config.Configuration, userId int) (users model.User, err error) {
	db := conf.MysqlDB
	return repo.GetUserDetail(db, userId)
}

// create new product
func CreateNewProduct(conf config.Configuration, in *model.Product) (res *model.IsSuccess, err error) {
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return nil, err
	}

	in = in.Generate()

	_, err = repo.CreateNewProduct(tx, in)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return &model.IsSuccess{
		IsSuccess: true,
	}, nil
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
