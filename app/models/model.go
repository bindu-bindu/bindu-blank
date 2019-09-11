package models

import "time"

// ModelUint64 Auto gen ID Primary, CreatedAt, UpdatedAt, DeletedAt Property
type ModelUint64 struct {
	ID        uint64 `gorm:"primary_key" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
