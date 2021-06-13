package model

import (
	"time"
)

type User struct {
	ID        uint      `form:"id" json:"id"`
	Email     string    `form:"email" json:"email"`
	Name      string    `form:"name" json:"name"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
