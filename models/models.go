package models

type UUID string

type ShortLink struct {
	UUID int64
	Slug string
	URL string
	IsCustom bool
	CreatedAt int64
}

type ShortLinkVisit struct {
	UUID UUID
	Slug string
	Datestr string
	CreatedAt int64
}

type LinkReqBody struct {
	URL string `json:"url"`
	Slug *string `json:"slug,omitempty"`
}

type StatsReqBody struct {
	Slug *string `json:"slug"`
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime *int64 `json:"end_time,omitempty"`
}

type StatsCountRespBody struct {
	Count *int64 `json:"count"`
}

type StatsHistRespBody struct {
	Histogram map[string]int64 `json:"histogram"`
}

type StatsCreatedAtRespBody struct {
	CreatedAt *int64 `json:"created_at"`
}