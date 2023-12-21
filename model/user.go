package model

import "time"

type User struct {
	ID        int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(255);column:name" json:"name"`
	Email     string    `gorm:"type:varchar(255);column:email" json:"email"`
	Address   string    `gorm:"type:text;column:address" json:"address"`
	Password  string    `gorm:"type:varchar(255);column:password" json:"password"`
	WANumber  string    `gorm:"type:varchar(255);column:wa_number" json:"wa_number"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
