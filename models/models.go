package models

type UUID string

type ShortLink struct {
	UUID UUID
	Slug string
	URL string
	CreatedAt int64
}

type ShortLinkVisit struct {
	UUID UUID
	Slug string
	Datestr string
	CreatedAt int64
}

type LinkReqBody struct {
	URL string `json:"link"`
	Slug *string `json:"slug,omitempty"`
}

type StatsReqBody struct {
	Slug string `json:"slug"`
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime *int64 `json:"end_time,omitempty"`
}

type StatsRespBody struct {
	Count int64 `json:"count"`
	Histogram map[string]int64 `json:"histogram"`
	CreatedAt int64 `json:"created_at"`
}