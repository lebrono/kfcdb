package models

import (
	"time"
)

type Category struct {
	Category_Id int 
	Category_Name  string 
	IsSync int 
	IsDeleted int 
	IsActive int 
	Date_Created time.Time 
	Date_Updated time.Time 
}