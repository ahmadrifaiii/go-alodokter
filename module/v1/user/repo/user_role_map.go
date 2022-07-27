package repo

import (
	"database/sql"
	"go-alodokter/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

// get user role map detail
func GetUserRoleMaps(sqlx *sqlx.DB, userId string) (result *[]model.UserRoleMap, err error) {

	// sql builder
	sb := sqlbuilder.NewSelectBuilder()
	sb.From(model.TabelUserRoleMap)
	sb.Select("user_role_map_id", "user_id", "user_roles.user_role_id", "user_role_name")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "user_roles", "user_roles.user_role_id = user_role_maps.user_role_id")

	sb.Where(
		sb.Equal("user_role_maps.user_id", userId),
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

	r := []model.UserRoleMap{}
	for rows.Next() {
		var py model.UserRoleMap
		if err := rows.Scan(&py.Id, &py.UserId, &py.RoleId, &py.RoleName); err != nil {
			log.Error(err)
			continue
		}

		r = append(r, py)
	}

	return &r, nil

}
