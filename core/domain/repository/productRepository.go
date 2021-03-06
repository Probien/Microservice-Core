package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IProductRepository interface {
	GetById(c *gin.Context) (*domain.Product, error)
	GetAll() (*[]domain.Product, error)
	Create(c *gin.Context) (*domain.Product, error)
	Update(c *gin.Context) (*domain.Product, error)
}
