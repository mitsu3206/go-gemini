package repository

import (
	"go-gemini/domain"

	"gorm.io/gorm"
)

// GormTagRepository implements domain.TagRepository using GORM.
type GormTagRepository struct {
	DB *gorm.DB
}

// NewGormTagRepository creates a new GormTagRepository.
func NewGormTagRepository(db *gorm.DB) *GormTagRepository {
	return &GormTagRepository{DB: db}
}

// Create adds a new Tag to the database using GORM.
func (r *GormTagRepository) Create(tag *domain.Tag) (*domain.Tag, error) {
	result := r.DB.Create(tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return tag, nil
}

// FindByID retrieves a Tag item by its ID using GORM.
func (r *GormTagRepository) FindByID(id uint) (*domain.Tag, error) {
	var tag domain.Tag
	result := r.DB.First(&tag, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

// FindByName retrieves a Tag item by its name using GORM.
func (r *GormTagRepository) FindByName(name string) (*domain.Tag, error) {
	var tag domain.Tag
	result := r.DB.Where("name = ?", name).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

// FindAll retrieves all Tag items using GORM.
func (r *GormTagRepository) FindAll() ([]*domain.Tag, error) {
	var tags []*domain.Tag
	result := r.DB.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}

// Update updates an existing Tag item in the database using GORM.
func (r *GormTagRepository) Update(tag *domain.Tag) (*domain.Tag, error) {
	result := r.DB.Save(tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return tag, nil
}

// Delete deletes a Tag item by its ID using GORM.
func (r *GormTagRepository) Delete(id uint) error {
	result := r.DB.Delete(&domain.Tag{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
