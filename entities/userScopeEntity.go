package entities

import "time"

type StdUserScope struct {
	ID        uint32    `gorm:"size:11;primary_key;not null;auto_increment" json:"id"`
	Scope     string    `gorm:"size:255;not null;" json:"scope"`
	UID       string    `gorm:"size:255;not null;" json:"uid"`
	IsActive  bool      `gorm:"default:false;" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
}
