package models

import (
	"time"
)

type Transaction struct {
	Transaction_Id int `json:"transaction_id"`
	Transaction_No  string `json:"transaction_no"`
	Order string `json:"order"`
	Menu_Name string `json:"menu_name"`
	Quantity int `json:"quantity"`
	Amount string `json:"amount"`
	Total_Amount string `json:"total_amount"`
	Transaction_Grandtotal string `json:"trasanction_grandtotal"`
	IsPoint int `json:"is_point"`
	Points_Used string `json:"points_used"`
	Points_Added string `json:"points_added"`
	Fun_Id string `json:"fun_id"`
	IsOk_Kitchen int `json:is_ok_kitchen`
	IsOk_POS int `json:is_ok_pos`
	Order_Type string `json:"order_type"`
	Status string `json:"status"`
	Pickup_Datetime time.Time `json:"pickup_datetime"`
	User_Id string `json:"user_id"`
	User_Name string `json:"user_name"`
	Discount_Type string `json:"discount_type"`
	Discount_Total string `json:"discount_total"`
	Discount_AmountPercent string `json:"discount_amount_percent"`
	Branch_Code string `json:"branch_code"`
	Date_Created time.Time `json:"date_created"`
	Date_Updated time.Time `json:"date_updated"`
}