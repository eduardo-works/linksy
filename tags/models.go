package tags

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID     uint64    `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID   uuid.UUID `gorm:"type:char(36);uniqueIndex;not null"`
	Title  string    `gorm:"size:255;not null;uniqueIndex:idx_user_tag_title,priority:2"`
	UserID uint64    `gorm:"not null;uniqueIndex:idx_user_tag_title,priority:1" json:"-"`
	Color  string    `gorm:"type:char(6);not null;default:9FBE9A"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
