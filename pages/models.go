package pages

import (
	"time"

	"github.com/eduardo-works/linksy/links"
	"github.com/google/uuid"
)

type Page struct {
	ID     uint64       `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID   uuid.UUID    `gorm:"type:char(36);uniqueIndex;not null"`
	Title  string       `gorm:"size:255;not null;uniqueIndex:idx_user_page_title"`
	UserID uint64       `gorm:"not null;uniqueIndex:idx_user_page_title" json:"-"`
	Tags   []links.Link `gorm:"many2many:link_tags;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
