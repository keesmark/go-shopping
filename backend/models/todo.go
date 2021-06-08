package models

type Todo struct {
	Id       int    `form:"id"`
	InCharge string `form:"inCharge"`
	Content  string `form:"content"`
	Status   int    `form:"status"`
}
