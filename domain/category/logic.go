package category

import (
	"github.com/google/uuid"
	"github.com/chagasVinicius/apollo/domain/database"
	"github.com/zmb3/spotify/v2"
)

func CreateCategory(categoryName string, playlists *spotify.SearchResult) *database.Category {
	var playlistsIDs []string
	categoryID := uuid.New().String()

	for _, playlist := range playlists.Playlists.Playlists {
		playlistsIDs = append(playlistsIDs, playlist.ID.String())
	}

	return &database.Category{
		ID: categoryID,
		Name: categoryName,
		PlaylistsID: playlistsIDs,
	}


}
