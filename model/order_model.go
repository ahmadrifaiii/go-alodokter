package model

import (
	"time"

	"github.com/google/uuid"
)

const TabelOrder = "orders"

type Order struct {
	Id        string  `json:"id" db:"order_id" fieldtag:"insert"`
	Number    string  `json:"name" db:"order_number" fieldtag:"insert,update"`
	DeletedAt *string `json:"deleted" db:"deleted_at" fieldtag:"insert,update,delete"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy *string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy *string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

func (o Order) Generate() *Order {

	if o.Id == "" {
		o.Id = uuid.New().String()
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	if o.CreatedAt == nil {
		o.CreatedAt = &now
	}

	if o.UpdatedAt == nil {
		o.UpdatedAt = &now
	}

	return &o
}
