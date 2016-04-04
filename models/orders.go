package models

import ( 
	"encoding/json"
)

type Order struct {
	Menu_Id json.Number `json:"menu_id,Number"`
	Menu_Name  string `json:"menu_name"`
	Menu_Amount json.Number `json:"menu_amount,Number"`
	Qty json.Number `json:"qty,Number"`
	Total json.Number `json:"total,Number"`
}