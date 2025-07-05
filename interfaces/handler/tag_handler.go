package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-gemini/usecase"
)

// TagHandler handles HTTP requests related to Tag items.
type TagHandler struct {
	TagUseCase *usecase.TagUseCase
}

// NewTagHandler creates a new TagHandler.
func NewTagHandler(uc *usecase.TagUseCase) *TagHandler {
	return &TagHandler{TagUseCase: uc}
}

// CreateTagRequest represents the request body for creating a Tag.
type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateTag handles the creation of a new Tag item.
func (h *TagHandler) CreateTag(c *gin.Context) {
	var req CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.TagUseCase.CreateTag(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// GetTag handles the retrieval of a single Tag item by ID.
func (h *TagHandler) GetTag(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	tag, err := h.TagUseCase.GetTagByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, tag)
}

// GetAllTags handles the retrieval of all Tag items.
func (h *TagHandler) GetAllTags(c *gin.Context) {
	tags, err := h.TagUseCase.GetAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// UpdateTagRequest represents the request body for updating a Tag.
type UpdateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateTag handles the update of an existing Tag item.
func (h *TagHandler) UpdateTag(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTag, err := h.TagUseCase.GetTagByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	existingTag.Name = req.Name

	updatedTag, err := h.TagUseCase.UpdateTag(existingTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTag)
}

// DeleteTag handles the deletion of a Tag item by ID.
func (h *TagHandler) DeleteTag(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.TagUseCase.DeleteTag(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
