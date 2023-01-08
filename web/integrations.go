package web

import (
	"context"
	"log"
	"os"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	spotifyClient *spotify.Client
)


func Load() {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID: os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL: spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	spotifyClient = spotify.New(httpClient)
}

func CategoryPlaylists(category string) *spotify.SearchResult {
	ctx := context.Background()
	results, err := spotifyClient.Search(ctx, category, spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)

	if err != nil {
		log.Fatal(err)
	}

	return results

}

func GetPlaylistItems(playlistID spotify.ID) *spotify.PlaylistItemPage {
	ctx := context.Background()

	tracks, err := spotifyClient.GetPlaylistItems(
		ctx,
		spotify.ID(playlistID),
	)

	if err != nil {
		log.Fatal(err)
	}

	return tracks
}
