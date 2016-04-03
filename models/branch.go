package models

import (
	"time"
)

type Branch struct {
	Branch_Id int `json:"branch_id"`
	Branch_Code  string `json:"branch_code"`
	Branch_Name string `json:"branch_name"`
	Branch_Address string `json:"branch_address"`
	Branch_Email string `json:"branch_email"`
	Branch_Contact_No string `json:"branch_contact_no"`
	Latitute string `json:"latitude"`
	Longitute string `json:"longitude"`
	Date_Created time.Time `json:"date_created"`
	Date_Updated time.Time `json:"date_updated"`
}