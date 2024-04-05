package sqlserver

import (
	"fmt"
	"src/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	dsn := "sqlserver://gorm:gormTest1234@DESKTOP-SS75E3K\\TEW_SQLEXPRESS?database=articles"
	var err error
	Database,
		err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	Database.AutoMigrate(&models.Article{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
