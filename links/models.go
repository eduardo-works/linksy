package links

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID      uuid.UUID `gorm:"type:char(36);uniqueIndex;not null"`
	Title     string    `gorm:"size:255;not null"`
	URL       string    `gorm:"type:text;not null"`
	URLHash   string    `gorm:"type:char(64);not null;uniqueIndex:idx_user_url,priority:2" json:"-"`
	UserID    uint64    `gorm:"not null;uniqueIndex:idx_user_url,priority:1" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
