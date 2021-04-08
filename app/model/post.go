package model

import (
	"time"
)

type Post struct {
	Id        uint   `gorm:"primary_key;auto_increment;not_null"`
	Content   string `gorm:"size:256"`
	CreatedAt time.Time
}
