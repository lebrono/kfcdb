package models

import (
	"time"
)

type FunCard struct {
	Id int `json:"id"`
	Fun_Idnumber  string `json:"fun_id_number"`
	Fun_Name string `json:"fun_name"`
	Fun_Points string `json:"fun_points"`
	Date_Created time.Time `json:"date_created"`
	Date_Updated time.Time `json:"date_updated"`
}


