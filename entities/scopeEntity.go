package entities

import "time"

type StdScope struct {
	ID        uint32    `gorm:"size:11;primary_key;not null;auto_increment" json:"id"`
	Scope     string    `gorm:"size:255;not null;" json:"scope"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
}
