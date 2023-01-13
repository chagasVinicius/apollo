package category

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/zmb3/spotify/v2"
)

const basePath = "/persistence/playlist/items/%s.json"

func CreateCategory(categoryName string, playlists *spotify.SearchResult) *Category {
	var allPlaylists []Playlist
	categoryID := uuid.New().String()

	for _, playlist := range playlists.Playlists.Playlists {
		structedPlaylist := CratePlaylist(playlist)
		allPlaylists = append(allPlaylists, *structedPlaylist)
	}

	return &Category{
		ID: categoryID,
		Name: categoryName,
		Playlists: allPlaylists,
	}

}

func CratePlaylist(playlist spotify.SimplePlaylist) *Playlist {
	playlistID := uuid.New().String()
	return &Playlist{
		ID: playlistID,
		Items: PlaylistPath(playlistID),
	}
}

func PlaylistPath(playlistID string) string {
	return fmt.Sprintf(basePath, playlistID)
}
