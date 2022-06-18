package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database *gorm.DB
}

func NewEndorsementRepositoryImpl(db *gorm.DB) repository.IEndorsementRepository {
	return &EndorsementRepositoryImpl{database: db}
}

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsement, c.Param("id")).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if endorsement.ID == 0 {
		return nil, errors.New(ENDORSEMENT_NOT_FOUND)
	}

	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll() (*[]domain.Endorsement, error) {
	var endorsements []domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsements).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	return &endorsements, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := c.ShouldBindJSON(&endorsement); err != nil || endorsement.PawnOrderID == 0 {
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.Endorsement{}).Create(&endorsement).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_INSERT, nil, &endorsement)
	return &endorsement, nil
}
