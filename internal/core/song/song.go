package song

import (
	"context"
	"fmt"

	spotify "github.com/zmb3/spotify/v2"
)

type SongClient interface {
	GetPlaylists(ctx context.Context, categoryID string, categoryName string) ([]spotify.SimplePlaylist, error)
}

type Storer interface {
	AddSong(ctx context.Context, song Song) error
}

type Core struct {
	client SongClient
}

func NewCore(songClient SongClient) Core {
	return Core{
		client: songClient,
	}
}

func (c Core) SearchPlaylists(ctx context.Context, categoryID string, categoryName string) error {
	playlists, err := c.client.GetPlaylists(ctx, categoryID, categoryName)
	if err != nil {
		return fmt.Errorf("error searching playlists for category [categoryName=%s]: %w", categoryName, err)
	}

	for _, playlist := range playlists {
		fmt.Println(playlist.Description)
	}
	return nil
}
