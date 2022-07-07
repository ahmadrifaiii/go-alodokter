package repo

import (
	"database/sql"
	"go-alodokter/config/database"
	"go-alodokter/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

// get user list
func GetUserList(sqlx *sqlx.DB) (users []model.User, err error) {
	users = make([]model.User, 0)
	var ModelUser model.User

	// sql builder
	st := sqlbuilder.NewStruct(ModelUser)
	sb := st.SelectFrom(model.TabelUser)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return users, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return users, err
		}
		return users, err
	}

	for rows.Next() {
		var usr model.User
		if err := rows.Scan(st.Addr(&usr)...); err != nil {
			log.Error(err)
			continue
		}

		users = append(users, usr)
	}

	return
}

// get user detail
func GetUserDetail(sqlx *sqlx.DB, userId int) (user model.User, err error) {
	var ModelUser model.User

	// sql builder
	st := sqlbuilder.NewStruct(ModelUser)
	sb := st.SelectFrom(model.TabelUser)
	sb.Where(
		sb.Equal("id", userId),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return user, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return user, err
		}
		return user, err
	}

	row.Scan(st.Addr(&user)...)

	return
}

// register user
func RegisterUser(tx *sql.Tx, p *model.UserPayload) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.UserPayload{})
	sb := st.InsertIntoForTag(model.TabelUser, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// update user
func UpdateUser(tx *sql.Tx, p *model.UserPayload) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.UserPayload{})
	sb := st.UpdateForTag(model.TabelUser, "update", *p)

	sb.Where(
		sb.Equal("id", p.Id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// delete user
func DeleteUser(tx *sql.Tx, p *model.UserPayload) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.UserPayload{})
	sb := st.UpdateForTag(model.TabelUser, "delete", *p)
	sb.Where(
		sb.Equal("id", p.Id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// get user detail by param
func GetUserDetailByParam(sqlx *sqlx.DB, param string, value interface{}) (user model.User, err error) {
	var ModelUser model.User

	// sql builder
	st := sqlbuilder.NewStruct(ModelUser)
	sb := st.SelectFrom(model.TabelUser)
	sb.Where(
		sb.Equal(param, value),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return user, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return user, err
		}
		return user, err
	}

	row.Scan(st.Addr(&user)...)

	return
}
