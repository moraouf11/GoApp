package repositories

import (
	// sqlserver "src/database"
	// models "src/models"
	// utils "src/utils"

	"gorm.io/gorm"
)

// type Repository interface {
// 	GetByID(id int) (*models.Job, error)
// 	GetAll() ([]models.Job, error)
// 	Create(user *models.Job) error
// 	Update(user *models.Job) error
// 	Delete(id int) error
// }

type DBRepository struct {
	dbConn *gorm.DB
}

// func  GetByID(db *DBRepository, id int)(*models.Job, error) {
// 	//id := mux.Vars(r)["jobKey"]

// 	var bids models.Bid
// 	//db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)

// 	//result := sqlserver.Database.Where("Id = ?", id).First(&job)
// 	result := sqlserver.Database.Joins("Job", sqlserver.Database.Where(&models.Job{JobKey: id})).Find(&bids)
// 	if result.Error == nil {
// 		utils.JsonResponse(w, bids)
// 	}
// }

// func (db *DBRepository) GetAll() ([]models.Job, error) {
// 	var Jobs []models.Job
// 	sqlserver.Database.Find(&Jobs)//.Where()
// 	return Jobs, nil
// }
// func (db *DBRepository) Create(id int) (*models.Job, error) {
// 	// Implementation
// }
// func (db *DBRepository) Update(id int) (*models.Job, error) {
// 	// Implementation
// }
// func (db *DBRepository) Delete(id int) (*models.Job, error) {
// 	// Implementation
// }
