package db

import (
	"time"
)

const (
	insertVisitStmt = `INSERT INTO short_link_visit (slug, datestr, created_at) VALUES ($1, $2, $3)`
	visitCountStmt = `SELECT count(*) FROM short_link_visit WHERE slug=$1 AND created_at BETWEEN $2 AND $3`
	histogramStmt = `SELECT count(*) FROM short_link_visit WHERE slug=$1 AND created_at BETWEEN $2 AND $3 GROUP BY datestr`
	creationDate = 	`SELECT created_at FROM short_link_visit WHERE slug=$1 ORDER BY created_at ASC LIMIT 1`
)

// MakeShortLinkVisit records that a short link was visited
func (db *DB) MakeShortLinkVisit(slug string) error {
	t := time.Now()
	_, err := db.Connection.Exec(insertVisitStmt, slug, t.Format("2000-01-01"), t.Unix())
	if err != nil {
		return err
	}
	return nil
}

// GetShortLinkVisitCount returns the number of times a short link was visited
func (db *DB) GetShortLinkVisitCount(slug string, start *int64, end *int64) (*int64, error) {
	stmt, err := db.Connection.Prepare(visitCountStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var count *int64
	err = stmt.QueryRow(slug, start, end).Scan(count)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (db *DB) GetShortLinkVisitHistogram(slug string, start *int64, end *int64) (map[string]int64, error) {
	stmt, err := db.Connection.Prepare(histogramStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var histogram map[string]int64
	err = stmt.QueryRow(slug, start, end).Scan(histogram)
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

	var createdAt *int64
	err = stmt.QueryRow(slug).Scan(createdAt)
	if err != nil {
		return nil, err
	}

	return createdAt, nil
}