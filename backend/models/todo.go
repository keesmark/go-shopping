package models

type Todo struct {
	Id       int
	InCharge string `form:"inCharge"`
	Content  string `form:"content"`
	Status   int    `form:"status"`
}
