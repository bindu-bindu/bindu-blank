package models

import "time"

// DefaultProperties Auto gen ID Primary, CreatedAt, UpdatedAt, DeletedAt Property
type defaultProperties struct {
	ID        uint64 `gorm:"primary_key" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
