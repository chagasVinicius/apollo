package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Load() {
	database, _ = gorm.Open(sqlite.Open("./domain/database/database.db"), &gorm.Config{})
	database.AutoMigrate(&Category{})
	database.AutoMigrate(&Playlists{})
}

type Category struct {
	gorm.Model
	ID string
	Name string
	PlaylistsID []string
}

type Playlists struct {
	gorm.Model
	ID string
	PlaylistItems string // filepath for now
}

func CreateCategory(category Category) {
	database.Create(category)
}

func UpdateCategory(category Category) {
	database.Save(category)
}
