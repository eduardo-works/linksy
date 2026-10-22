package users

import (
	"time"

	"github.com/eduardo-works/linksy/pages"
	"github.com/google/uuid"
)

type User struct {
	ID           uint64       `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID         uuid.UUID    `gorm:"type:char(36);uniqueIndex;not null"`
	Email        string       `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string       `gorm:"size:255;not null" json:"-"`
	Pages        []pages.Page `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
