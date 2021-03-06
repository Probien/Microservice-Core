package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type ICategoryRepository interface {
	GetById(c *gin.Context) (*domain.Category, error)
	GetAll() (*[]domain.Category, error)
	Create(c *gin.Context) (*domain.Category, error)
	Delete(c *gin.Context) (*domain.Category, error)
	Update(c *gin.Context) (*domain.Category, error)
}
