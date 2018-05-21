package models

type UUID string

type ShortLink struct {
	UUID int64 `db:"id"`
	Slug string `db:"slug"`
	URL string `db:"url"`
	CreatedAt int64 `db:"created_at"`
}

// type CustomLink struct {
// 	UUID int64 `db:"id"`
// 	ShortLinkUUID int64 `db:"short_link_uuid"`
// 	Slug string `db:"slug"`
// 	URL string `db:"url"`
// 	CreatedAt int64 `db:"created_at"`
// }

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
	Slug string `json:"slug"`
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime *int64 `json:"end_time,omitempty"`
}

type StatsRespBody struct {
	Count int64 `json:"count"`
	Histogram map[string]int64 `json:"histogram"`
	CreatedAt int64 `json:"created_at"`
}