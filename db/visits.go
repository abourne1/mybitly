package db

import (
	"time"
)

const (
	datestrFormatStr = "2006-01-02"

	insertVisitStmt = `INSERT INTO short_link_visit (slug, datestr, created_at) VALUES ($1, $2, $3)`
	visitCountStmt = `SELECT count(*) FROM short_link_visit WHERE slug=$1 AND created_at BETWEEN $2 AND $3`
	histogramStmt = `SELECT datestr, count(*) FROM short_link_visit WHERE slug=$1 AND created_at BETWEEN $2 AND $3 GROUP BY datestr`
	creationDate = 	`SELECT created_at FROM short_link_visit WHERE slug=$1 ORDER BY created_at ASC LIMIT 1`
)

func resolveTime(timePtr *int64, fallback int64) int64 {
	if timePtr != nil {
		return *timePtr
	}
	return fallback
}

// MakeShortLinkVisit records that a short link was visited
func (db *DB) MakeShortLinkVisit(slug string) error {
	t := time.Now()
	_, err := db.Connection.Exec(insertVisitStmt, slug, t.Format(datestrFormatStr), t.Unix())
	if err != nil {
		return err
	}
	return nil
}

// GetShortLinkVisitCount returns the number of times a short link was visited
func (db *DB) GetShortLinkVisitCount(slug string, startPtr *int64, endPtr *int64) (*int64, error) {
	stmt, err := db.Connection.Prepare(visitCountStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var count int64
	// if startPtr or endPtr is null, choose earliest or current time respectively 
	start := resolveTime(startPtr, 0)
	end := resolveTime(endPtr, time.Now().Unix())
	err = stmt.QueryRow(slug, start, end).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// GetShortLinkVisitHistogram returns a per-day link visit count
// no entry is included if the link was not visited on a given day
func (db *DB) GetShortLinkVisitHistogram(slug string, startPtr *int64, endPtr *int64) (map[string]int64, error) {
	// if startPtr or endPtr is null, choose earliest or current time respectively 
	start := resolveTime(startPtr, 0)
	end := resolveTime(endPtr, time.Now().Unix())
	rows, err := db.Connection.Query(histogramStmt, slug, start, end)
	defer rows.Close()

	histogram := map[string]int64{}
	// for each date that the slug was visited, write count to histogram map
	for rows.Next() {
		var datestr string
		var count int64
		err = rows.Scan(&datestr, &count)
		histogram[datestr] = count
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return histogram, nil
}

func (db *DB) GetShortLinkCreationDate(slug string) (*int64, error) {
	stmt, err := db.Connection.Prepare(creationDate)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var createdAt int64
	err = stmt.QueryRow(slug).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	return &createdAt, nil
}