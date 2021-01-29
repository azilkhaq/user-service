package entities

import "time"

type StdProfile struct {
	ID         uint32    `gorm:"size:11;primary_key;not null;auto_increment" json:"id"`
	Name       string    `gorm:"size:255;null;" json:"name"`
	Nim        string    `gorm:"size:255;null;" json:"nim"`
	Gender     string    `gorm:"type:enum('female','male');null" json:"gender"`
	ShortBio   string    `gorm:"text;null" json:"short_bio"`
	University string    `gorm:"size:255;null;unique;" json:"university"`
	Address    string    `gorm:"text;null;" json:"address"`
	UID        string    `gorm:"size:255;not null;unique;" json:"uid"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
}
