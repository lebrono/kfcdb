package models

import (
	"time"
)

type Menu struct {
	Menu_Id int `json:"menu_id"`
	Menu_Name  string `json:"menu_name"`
	Menu_Description string `json:"menu_description"`
	Menu_Amount string `json:"menu_amount"`
	Category_Name string `json:"category_name"`
	IsSync int `json:"isSync"`
	IsActive int `json:"isActive"`
	Date_Created time.Time `json:"date_created"`
	Date_Updated time.Time `json:"date_updated"`
}