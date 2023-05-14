package song

type Song struct {
	ID            string `json:"id"`
	TrackID       string `json:"track_id"`       // spotify track ID
	AudioFeatures string `json:"audio_features"` // spotify audio features
	CategoryID    string `json:"category_id"`
	CreatedAt     string `json:"created_at"`
}

type NewSong struct {
	TrackID       string
	AudioFeatures string
	CategoryID    string
}
