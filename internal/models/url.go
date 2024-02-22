package models

type Url struct {
	Url string `json:"url"`
}

type CreateUrl struct {
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

type UrlResponse struct {
	Id    int64  `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}
