package model

type UrlRequestCreate struct {
	LongUrl string `json:"long_url"`
}

type UrlResponse struct {
	Id       string `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
	UserId   string `json:"user_id"`
}
