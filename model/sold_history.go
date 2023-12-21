package model

import "time"

type SoldHistory struct {
	ID         int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Menu       int64     `gorm:"type:integer;column:menu" json:"menu"`
	Amount     int64     `gorm:"type:integer;column:amount" json:"amount"`
	TotalPrice int64     `gorm:"type:integer;column:total_price" json:"total_price"`
	CreatedAt  time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
