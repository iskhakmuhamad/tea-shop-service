package model

import "time"

type Menu struct {
	ID        int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(255);column:name" json:"name"`
	Price     int64     `gorm:"type:integer;column:price" json:"price"`
	ImageUrl  string    `gorm:"type:varchar(255);column:image_url" json:"image_url"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
