package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model

	PollID      int    `gorm:"not null;column:poll_id"`
	Poll        *Poll  `gorm:"foreignKey:PollID"`
	Description string `gorm:"not null; column:description"`
	Votes       int    `gorm:"default:0; column:votes"`
}
