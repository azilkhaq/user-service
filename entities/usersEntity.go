package entities

import "time"

type StdUser struct {
	Uid            string `gorm:"primary_key;size:255;not null;" json:"uid"`
	EmailAddress   string `gorm:"size:255;null;unique" json:"email_address"`
	PhoneNumber    string `gorm:"size:255;null;unique" json:"phone_number"`
	Password       string `gorm:"size:255;null;" json:"password"`
	Role           string `gorm:"size:255;not null;" json:"role"`
	Status         string `gorm:"size:255;not null;" json:"status"`
	EmailActivated uint32 `gorm:"size:1;not null;" json:"email_activated"`
	PhoneActivated uint32 `gorm:"size:1;not null;" json:"phone_activated"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
}
