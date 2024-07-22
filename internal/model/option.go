package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model

	Id          int    `gorm:"primaryKey"`
	PollId      int    `gorm:"not null;column:poll_id"`
	Description string `gorm:"not null; column:description"`
	Votes       int    `gorm:"default:0; column:votes"`
}
