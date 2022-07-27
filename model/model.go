package model

import "context"

const (
	RequestId = "REQUEST_ID"
)

type ReqIDContextKey string

// func set context value
func NewContext(ctx context.Context, key interface{}, p interface{}) context.Context {
	return context.WithValue(ctx, key, p)
}

// func parse context get value
func ParseContext(ctx context.Context, key interface{}) interface{} {
	if v := ctx.Value(key); v != nil {
		return v
	}
	return nil
}

type Response struct {
	LogId   string      `json:"log_id"`
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type TableAttributes struct {
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type IsSuccess struct {
	IsSuccess bool `json:"is_success"`
}
