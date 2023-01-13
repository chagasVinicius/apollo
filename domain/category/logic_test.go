package category

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"testing"

	spotify "github.com/zmb3/spotify/v2"
)

func TestCreateCategory(t *testing.T) {
	playlists := []spotify.SimplePlaylist{
		spotify.SimplePlaylist{
			Collaborative: false,
			Description:   "vibez pra treinar",
			ExternalURLs:  map[string]string{
				"spotify": "https://open.spotify.com/playlist/5mMcyD2KALPqAUy53Pcaic",
			},
			Endpoint: "https://api.spotify.com/v1/playlists/5mMcyD2KALPqAUy53Pcaic",
			ID:       "testeID1",
			Images:   []spotify.Image{
				spotify.Image{
					Height: 0,
					Width:  0,
					URL:    "https://i.scdn.co/image/ab67706c0000bebb5a00b041989033631e5145e3",
				},
			},
			Name:  "🔥 treino 🔥",
			Owner: spotify.User{
				DisplayName:  "minawinkel",
				ExternalURLs: map[string]string{
					"spotify": "https://open.spotify.com/user/minawinkel",
				},
				Followers: spotify.Followers{
					Count:    0,
					Endpoint: "",
				},
				Endpoint: "https://api.spotify.com/v1/users/minawinkel",
				ID:       "minawinkel",
				Images:   []spotify.Image(nil),
				URI:      "spotify:user:minawinkel",
			},
			IsPublic:   false,
			SnapshotID: "NTAsNjA5MTZhYzFjMjE2OWZhYWZlODExMGU3MzQ3ZGM1YmQ0OGZhNGQ3OA==",
			Tracks:     spotify.PlaylistTracks{
				Endpoint: "https://api.spotify.com/v1/playlists/5mMcyD2KALPqAUy53Pcaic/tracks",
				Total:    110,
			},
			URI: "spotify:playlist:5mMcyD2KALPqAUy53Pcaic",
		},

		spotify.SimplePlaylist{
			Collaborative: false,
			Description:   "vibez pra treinar",
			ExternalURLs:  map[string]string{
				"spotify": "https://open.spotify.com/playlist/5mMcyD2KALPqAUy53Pcaic",
			},
			Endpoint: "https://api.spotify.com/v1/playlists/5mMcyD2KALPqAUy53Pcaic",
			ID:       "testeID2",
			Images:   []spotify.Image{
				spotify.Image{
					Height: 0,
					Width:  0,
					URL:    "https://i.scdn.co/image/ab67706c0000bebb5a00b041989033631e5145e3",
				},
			},
			Name:  "🔥 treino 🔥",
			Owner: spotify.User{
				DisplayName:  "minawinkel",
				ExternalURLs: map[string]string{
					"spotify": "https://open.spotify.com/user/minawinkel",
				},
				Followers: spotify.Followers{
					Count:    0,
					Endpoint: "",
				},
				Endpoint: "https://api.spotify.com/v1/users/minawinkel",
				ID:       "minawinkel",
				Images:   []spotify.Image(nil),
				URI:      "spotify:user:minawinkel",
			},
			IsPublic:   false,
			SnapshotID: "NTAsNjA5MTZhYzFjMjE2OWZhYWZlODExMGU3MzQ3ZGM1YmQ0OGZhNGQ3OA==",
			Tracks:     spotify.PlaylistTracks{
				Endpoint: "https://api.spotify.com/v1/playlists/5mMcyD2KALPqAUy53Pcaic/tracks",
				Total:    110,
			},
			URI: "spotify:playlist:5mMcyD2KALPqAUy53Pcaic",
		},
	}

	searchResult := &spotify.SearchResult{
		Playlists: &spotify.SimplePlaylistPage {
			Playlists: playlists,
		},
	}

	newCategory := CreateCategory("teste", searchResult)
	categoryName := "teste"
	categoryPlaylistsID := []string{"testeID1", "testeID2"}

	if newCategory.Name != categoryName {
		t.Errorf("got %q, wanted %q", categoryName, newCategory.Name)
	}

	if !reflect.DeepEqual(newCategory.PlaylistsID, categoryPlaylistsID) {
	 	t.Errorf("got %q, wanted %q", categoryPlaylistsID, newCategory.PlaylistsID)
	}
}

func TestPlaylistPath(t *testing.T) {
	categoryID := uuid.New()
	playlistID := "d2814209-3c95-4e4e-99f4-0ed7b915ec1e"
	expectedPath := fmt.Sprintf("/persistent/category/%s/d2814209-3c95-4e4e-99f4-0ed7b915ec1e.json", categoryID.String())

	path := PlaylistPath(categoryID, playlistID)

	if !reflect.DeepEqual(expectedPath, path) {
	 	t.Errorf("got %q, wanted %q", expectedPath, path)
	}
}
