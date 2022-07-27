package repo

import (
	"database/sql"
	"go-alodokter/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

// get user role map detail
func GetUserRoleAccess(sqlx *sqlx.DB, roleIds []string) (result *[]model.UserRoleAccess, err error) {

	// sql builder
	sb := sqlbuilder.NewSelectBuilder()
	sb.From(model.TabelUserRoleAccess)
	sb.Select("user_role_id", "is_access_allowed", "user_role_access_url")

	sb.Where(
		sb.In("user_role_acceses.user_role_id", sqlbuilder.Flatten(roleIds)...),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return nil, err
		}

		return nil, err
	}

	r := []model.UserRoleAccess{}
	for rows.Next() {
		var py model.UserRoleAccess
		if err := rows.Scan(&py.RoleId, &py.IsAccess, &py.Url); err != nil {
			log.Error(err)
			continue
		}

		r = append(r, py)
	}

	return &r, nil

}
