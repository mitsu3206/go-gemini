package usecase

import (
	"go-gemini/domain"

	"gorm.io/gorm"
)

// TagUseCase defines the use case for Tag operations.
type TagUseCase struct {
	TagRepo domain.TagRepository
}

// NewTagUseCase creates a new TagUseCase.
func NewTagUseCase(repo domain.TagRepository) *TagUseCase {
	return &TagUseCase{TagRepo: repo}
}

// CreateTag creates a new Tag item.
func (uc *TagUseCase) CreateTag(name string) (*domain.Tag, error) {
	tag := &domain.Tag{
		Name: name,
	}
	return uc.TagRepo.Create(tag)
}

// GetTagByID retrieves a Tag item by its ID.
func (uc *TagUseCase) GetTagByID(id uint) (*domain.Tag, error) {
	return uc.TagRepo.FindByID(id)
}

// GetTagByName retrieves a Tag item by its name.
func (uc *TagUseCase) GetTagByName(name string) (*domain.Tag, error) {
	return uc.TagRepo.FindByName(name)
}

// GetAllTags retrieves all Tag items.
func (uc *TagUseCase) GetAllTags() ([]*domain.Tag, error) {
	return uc.TagRepo.FindAll()
}

// UpdateTag updates an existing Tag item.
func (uc *TagUseCase) UpdateTag(tag *domain.Tag) (*domain.Tag, error) {
	return uc.TagRepo.Update(tag)
}

// DeleteTag deletes a Tag item by its ID.
func (uc *TagUseCase) DeleteTag(id uint) error {
	return uc.TagRepo.Delete(id)
}

// GetOrCreateTags ensures tags exist and returns them.
func (uc *TagUseCase) GetOrCreateTags(tagNames []string) ([]*domain.Tag, error) {
	var tags []*domain.Tag
	for _, name := range tagNames {
		tag, err := uc.TagRepo.FindByName(name)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Tag not found, create it
				newTag, createErr := uc.CreateTag(name)
				if createErr != nil {
					return nil, createErr
				}
				tags = append(tags, newTag)
			} else {
				return nil, err
			}
		} else {
			tags = append(tags, tag)
		}
	}
	return tags, nil
}
