package model

const TabelUser = "user"

type User struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"user_name" fieldtag:"insert,update"`
	Email     string `json:"email" db:"user_email" fieldtag:"insert,update"`
	Password  string `json:"password" db:"user_password" fieldtag:"insert,update"`
	Deleted   int    `json:"deleted" db:"deleted" fieldtag:"insert,update,delete"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type UserPayload struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"user_name" fieldtag:"insert,update"`
	Email     string `json:"email" db:"user_email" fieldtag:"insert,update"`
	Password  string `json:"password" db:"user_password" fieldtag:"insert"`
	Deleted   int    `json:"deleted" db:"deleted" fieldtag:"insert,delete"`
	CreatedAt string `json:"created_at" db:"created_at"`
	CreatedBy string `json:"created_by" db:"created_by" fieldtag:"insert"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	UpdatedBy string `json:"updated_by" db:"updated_by" fieldtag:"insert,update"`
}
