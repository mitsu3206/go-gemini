package domain

import "gorm.io/gorm"

// Tag represents a tag for a todo item.
type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex"`
	Todos []*Todo `json:"todos" gorm:"many2many:todo_tags;"`
}

// TagRepository defines the interface for tag data storage.
type TagRepository interface {
	Create(tag *Tag) (*Tag, error)
	FindByID(id uint) (*Tag, error)
	FindByName(name string) (*Tag, error)
	FindAll() ([]*Tag, error)
	Update(tag *Tag) (*Tag, error)
	Delete(id uint) error
}
