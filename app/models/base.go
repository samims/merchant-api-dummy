package models

import (
	"time"
)

// BaseModel holds the basic attributes requires for models
type BaseModel struct {
	Id        int64     `json:"id" orm:"column(id);auto"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

// Pagination is common models to contain pagination data
type Pagination struct {
	Page *int64 `json:"page"`
	Size *int64 `json:"size"`
}
