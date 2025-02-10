package models

type ShortenLink struct {
	FullLink string `json:"full_link"`
}

type ShortenedLink struct {
	ShortLink string `json:"short_link"`
}
