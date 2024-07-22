package model

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model

	Id       int       `gorm:"primaryKey"`
	Question string    `gorm:"not null;column:question"`
	Options  []*Option `gorm:"foreignKey:poll_id"`
	Active   bool      `gorm:"column:active"`
	EndDate  time.Time `gorm:"column:end_date"`
}
