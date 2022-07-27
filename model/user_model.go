package model

const TabelUser = "users"
const TabelUserRole = "user_roles"
const TabelUserRoleMap = "user_role_maps"
const TabelUserRoleAccess = "user_role_acceses"

type User struct {
	Id        string  `json:"id" db:"user_id"`
	Name      string  `json:"name" db:"user_name" fieldtag:"insert,update"`
	Email     string  `json:"email" db:"user_email" fieldtag:"insert,update"`
	Password  string  `json:"password" db:"user_password" fieldtag:"insert,update"`
	DeletedAt *string `json:"deleted" db:"deleted_at" fieldtag:"insert,update,delete"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy *string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy *string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type UserDetail struct {
	User
	Roles  *[]UserRoleMap    `json:"roles"`
	Access *[]UserRoleAccess `json:"access"`
}

type UserPayload struct {
	Id        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"user_name" fieldtag:"insert,update"`
	Email     string  `json:"email" db:"user_email" fieldtag:"insert,update"`
	Password  string  `json:"password" db:"user_password" fieldtag:"insert"`
	Deleted   *string `json:"deleted" db:"deleted" fieldtag:"insert,delete"`
	CreatedAt *string `json:"created_at" db:"created_at"`
	CreatedBy *string `json:"created_by" db:"created_by" fieldtag:"insert"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
	UpdatedBy *string `json:"updated_by" db:"updated_by" fieldtag:"insert,update"`
}

type UserRoleMap struct {
	Id       string `json:"id" db:"user_role_map_id"`
	UserId   string `json:"user_id" db:"user_id"`
	RoleId   string `json:"role_id" db:"user_role_id"`
	RoleName string `json:"role_name" db:"user_role_name"`
}

type UserRoleAccess struct {
	RoleId   string `json:"role_id" db:"user_role_id"`
	IsAccess bool   `json:"is_access" db:"is_access_allowed"`
	Url      string `json:"url" db:"user_role_access_url"`
}
