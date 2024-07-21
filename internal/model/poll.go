package model

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model

	Question string    `gorm:"not null; column:question"`
	Options  []*Option `gorm:"foreignKey:PollID"`
	Active   bool      `gorm:"column:active"`
	EndDate  time.Time `gorm:"column:end_date"`
}
