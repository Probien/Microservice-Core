package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) repository.ICategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := r.database.Model(&domain.Category{}).Preload("Products").Find(&category, c.Param("id")).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if category.ID == 0 {
		return nil, errors.New(CATEGORY_NOT_FOUND)
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll() (*[]domain.Category, error) {
	var categories []domain.Category

	if err := r.database.Model(&domain.Category{}).Find(&categories).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}
	return &categories, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.Category{}).Create(&category).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_INSERT, nil, &category)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Find(&category, c.Param("id"))
	if category.ID == 0 {
		return nil, errors.New(CATEGORY_NOT_FOUND)
	} else if len(category.Products) > 0 {
		return nil, errors.New(INVALID_ACTION)
	}

	if err := r.database.Model(&domain.Category{}).Unscoped().Delete(&category, &category.ID).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (*domain.Category, error) {
	patch, category, categoryOld := map[string]interface{}{}, domain.Category{}, domain.Category{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New(ERROR_BINDING)
	}

	r.database.Model(&domain.Category{}).Find(&categoryOld, patch["id"])

	if err := r.database.Model(&domain.Category{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&category).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if category.ID == 0 {
		return nil, errors.New(CATEGORY_NOT_FOUND)
	}

	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_UPDATE, &categoryOld, &category)
	return &category, nil
}
