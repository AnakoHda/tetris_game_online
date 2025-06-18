package models

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Nickname  string    `db:"nickname" json:"nickname"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_time" json:"created_time"`
	UpdatedAt time.Time `db:"updated_time" json:"updated_time"`
}
