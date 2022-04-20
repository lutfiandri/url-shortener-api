package model

type UrlRequestCreate struct {
	LongUrl string `json:"long_url"`
	Title   string `json:"title"`
}

type UrlResponse struct {
	Id       string `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
	Title    string `json:"title"`
}
