package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Id       int    `gorm:"primaryKey" form:"id"`
	InCharge string `form:"inCharge" json:"inCharge"`
	Content  string `form:"content" json:"content"`
	Status   int    `form:"status" json:"status"`
}
