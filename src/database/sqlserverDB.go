package sqlserver

import (
	"fmt"
	"src/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	dsn := "sqlserver://gorm:gormTest1234@DESKTOP-SS75E3K?database=jobs"
	var err error
	Database,
		err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	Database.AutoMigrate(&models.Job{}, &models.Bid{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
