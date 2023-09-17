package dbManager

import (
	"gorm.io/gorm"
	"loco-transaction/middleware"
)

func GetDbInstance() (*gorm.DB, error) {
	var dbc = middleware.DBC
	return dbc, nil
}
