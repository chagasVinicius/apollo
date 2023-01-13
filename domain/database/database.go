package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/chagasVinicius/domain/category"
)

var (
	database *gorm.DB
)

func Load() {
	database, _ = gorm.Open(sqlite.Open("./domain/database/database.db"), &gorm.Config{})
	database.AutoMigrate(&category.Category{})
	database.AutoMigrate(&category.Playlists{})
}

func CreateCategory(category category.Category) {
	database.Create(category)
}

func UpdateCategory(category category.Category) {
	database.Save(category)
}
