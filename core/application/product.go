package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type ProductInteractor struct {
}

func (PI *ProductInteractor) GetById(c *gin.Context) (*domain.Product, error) {
	repository := persistance.NewProductRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (PI *ProductInteractor) GetAll() (*[]domain.Product, error) {
	repository := persistance.NewProductRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (PI *ProductInteractor) Create(c *gin.Context) (*domain.Product, error) {
	repository := persistance.NewProductRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (PI *ProductInteractor) Update(c *gin.Context) (*domain.Product, error) {
	repository := persistance.NewProductRepositoryImpl(config.Database)
	return repository.Update(c)
}
