package external

import (
	"Nutrition-App-Api/graph/external/entry"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entry.DishData{}, &entry.Menu{}, &entry.Recipe{}, &entry.Foodstuff{}, &entry.Nutrition{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
