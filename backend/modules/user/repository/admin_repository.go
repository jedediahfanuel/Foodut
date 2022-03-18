package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func CreateAdmin(admin model.Admin) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&admin)

	return result
}
