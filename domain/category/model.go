package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID string
	Name string
	Playlists [] Playlist
}

type Playlist struct {
	gorm.Model
	ID string
	Items string
}
