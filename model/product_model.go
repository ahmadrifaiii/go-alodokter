package model

import (
	"time"

	"github.com/google/uuid"
)

const TabelProduct = "products"

type Product struct {
	Id        string  `json:"id" db:"product_id" fieldtag:"insert"`
	Name      string  `json:"name" db:"product_name" fieldtag:"insert,update"`
	SKU       string  `json:"sku" db:"product_sku" fieldtag:"insert,update"`
	Price     float64 `json:"price" db:"product_price" fieldtag:"insert,update"`
	DeletedAt *string `json:"deleted" db:"deleted_at" fieldtag:"insert,update,delete"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy *string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy *string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

func (p *Product) Generate() *Product {

	if p.Id == "" {
		p.Id = uuid.New().String()
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	if p.CreatedAt == nil {
		p.CreatedAt = &now
	}

	if p.UpdatedAt == nil {
		p.UpdatedAt = &now
	}

	return p
}
