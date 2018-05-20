package db

import (
	"time"
	"database/sql"

	_ "github.com/lib/pq"

	"github/abourne1/mybitly/lib"
	"github/abourne1/mybitly/models"
)

const (
	uuidConversionBase = 62
	getStmt = `SELECT * FROM short_link WHERE slug=$1 LIMIT 1`
	insertStmt = `INSERT INTO short_link (url, created_at) VALUES ($1, $2) RETURNING uuid`
	addSlugStmt = `UPDATE short_link SET slug=$1 WHERE uuid=$2`
)

type DB struct {
	Connection *sql.DB
}

func New(connection *sql.DB) *DB {
	return &DB{
		Connection: connection,
	}
}

func (db *DB) MakeShortLink(url string, slug *string) error {
	// TODO: handle custom url case
	stmt, err := db.Connection.Prepare(insertStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var shortLinkUUID int64
	err = stmt.QueryRow(url, time.Now().Unix()).Scan(&shortLinkUUID)
	if err != nil {
		return err
	}

	base62UUID, _ := lib.ConvertBase(shortLinkUUID, uuidConversionBase)
	_, err = db.Connection.Exec(addSlugStmt, *base62UUID, shortLinkUUID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetShortLink(slug string) (*models.ShortLink, error) {
	stmt, err := db.Connection.Prepare(getStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var uuid int64
	var dbSlug string
	var url string
	var createdAt int64
	err = stmt.QueryRow(slug).Scan(&uuid, &dbSlug, &url, &createdAt)
	if err != nil {
		return nil, err
	}

	return &models.ShortLink{
		UUID: uuid,
		Slug: dbSlug,
		URL: url,
		CreatedAt: createdAt,
	}, nil
}

func (db *DB) MakeShortLinkVisit(slug string) error {
	// TODO: impolement business logic with db calls when appropriate
	return nil
}

func (db *DB) GetShortLinkVisitCount(slug string, start *int64, end *int64) (int64, error) {
	// TODO: implement logic w/ db calls when appropriate
	//
	// SELECT count(*) 
	// FROM shortlinks 
	// WHERE slug = slug AND createdAt BETWEEN start AND end
	return 100, nil
}

func (db *DB) GetShortLinkVisitHistogram(slug string, start *int64, end *int64) (map[string]int64, error) {
	// TODO: implement logic w/ db calls when appropriate
	//
	// SELECT count(*) 
	// FROM shortlinks 
	// WHERE slug = slug AND createdAt BETWEEN start AND end 
	// GROUP BY datestr
	return map[string]int64{}, nil
}

func (db *DB) GetShortLinkCreationDate(slug string) (int64, error) {
	// TODO: implement db logic later
	// 
	// SELECT createdAt
	// FROM shortlinks 
	// WHERE slug = slug
	// ORDER BY createdAt ASC
	// LIMIT 1
	return time.Now().Unix(), nil
}