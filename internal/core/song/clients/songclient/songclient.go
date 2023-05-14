package songclient

import (
	"context"
	"fmt"
	spotify "github.com/zmb3/spotify/v2"
)

type SongClient struct {
	songClient *spotify.Client
}

func NewSongClient(songClient *spotify.Client) SongClient {
	return SongClient{
		songClient: songClient,
	}
}

func (c SongClient) GetPlaylists(ctx context.Context, categoryID string, categoryName string) ([]spotify.SimplePlaylist, error) {

	results, err := c.songClient.Search(ctx, categoryName, spotify.SearchTypePlaylist)
	if err != nil {
		return []spotify.SimplePlaylist{}, fmt.Errorf("error searching for category playlists [category=%s]: %w", categoryName, err)
	}

	return results.Playlists.Playlists, nil
}
