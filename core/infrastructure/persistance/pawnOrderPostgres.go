package persistance

import (
	"errors"
	"log"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PawnOrderRepositoryImpl struct {
	database *gorm.DB
}

func NewPawnOrderRepositoryImpl(db *gorm.DB) repository.IPawnOrderRepository {
	return &PawnOrderRepositoryImpl{database: db}
}

func (r *PawnOrderRepositoryImpl) GetById(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Products").Preload("Employee").Preload("Customer").Preload("Status").Find(&pawnOrder, c.Param("id")).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if pawnOrder.ID == 0 {
		return nil, errors.New("pawn order not found")
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() (*[]domain.PawnOrder, error) {
	var pawnOrders []domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Products").Preload("Customer").Preload("Employee").Preload("Status").Preload("Endorsements").Find(&pawnOrders).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &pawnOrders, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := c.ShouldBindJSON(&pawnOrder); err != nil || pawnOrder.CustomerID == 0 {
		log.Print(err)
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.PawnOrder{}).Omit("Employee").Omit("Customer").Omit("Status").Create(&pawnOrder).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*domain.PawnOrder, error) {
	patch, pawnOrder := map[string]interface{}{}, domain.PawnOrder{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("Products").Omit("Endorsements").Updates(&patch).Find(&pawnOrder).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if pawnOrder.ID == 0 {
		return nil, errors.New("pawn order not found or json data does not match ")
	}

	return &pawnOrder, nil
}
