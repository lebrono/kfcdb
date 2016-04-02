package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"	
)

type Category struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CategoryName  string `form:"category_name" binding:"required"`
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}