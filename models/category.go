package models

import (
	"time"
)

type Category struct {
	Category_Id int `json:"category_id"`
	Category_Name  string `json:"category_name"`
	IsSync int `json:"is_sync"`
	IsDeleted int `json:"is_deleted"`
	IsActive int `json:"is_active"`
	Date_Created time.Time `json:"date_created"`
	Date_Updated time.Time `json:"date_updated"`
}