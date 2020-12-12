package user

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"`
	Username string `gorm:column:name`
	Email string `gorm:column:email`
	CreatedAt timestamp.Timestamp `gorm:column:created_at`
	updatedAt timestamp.Timestamp `gorm:column:update_at`
}

