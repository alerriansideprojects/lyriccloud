package models

type Song struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ArtistID int    `json:"artist"`
	Lyrics   string `json:"lyrics"`
}
