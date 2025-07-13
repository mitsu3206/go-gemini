package domain

import "gorm.io/gorm"

// TodoTag is the join table for Todo and Tag.
// It includes a gorm.DeletedAt field for soft deletes.
type TodoTag struct {
	TodoID    uint           `gorm:"primaryKey"`
	TagID     uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
