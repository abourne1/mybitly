package db

import (
	"log"
	"time"
	"database/sql"

	_ "github.com/lib/pq"

	"github/abourne1/mybitly/lib"
	"github/abourne1/mybitly/models"
)

const (
	uuidConversionBase = 62

	getBySlugStmt = `SELECT * FROM short_link WHERE slug=$1 LIMIT 1`
	getByURLStmt = `SELECT * FROM short_link WHERE url=$1 LIMIT 1`

	insertStmt = `INSERT INTO short_link (url, created_at) VALUES ($1, $2) RETURNING uuid`
	insertCustomStmt = `INSERT INTO short_link (url, slug, created_at) VALUES ($1, $2, $3) RETURNING uuid`

	addSlugStmt = `UPDATE short_link SET slug=$1 WHERE uuid=$2`
)

// MakeShortLink returns a short link if it already exists, else it creates a new one
// TODO: address race condition. Checking for short link and creating it should be one atomic action.
func (db *DB) MakeShortLink(url string, slug *string) (*models.ShortLink, error) {
	if slug != nil {
		// If slug already exists, return it
		shortLink, err := db.getShortLinkBySlug(*slug)
		if err != nil {
			log.Printf("[Error] MakeShortLink - db.getShortLinkBySlug: %v", err.Error())
			return nil, err
		}
		if shortLink != nil {
			return shortLink, nil
		}
		// If it does not exist, create a new one
		return db.makeCustomShortLink(url, *slug)

	}

	// If URL is already short linked, return it
	shortLink, err := db.getShortLinkByURL(url)
	if err != nil {
		return nil, err
	}
	if shortLink != nil {
		return shortLink, err
	}

	// If it does not exist, create a new short link
	return db.makeRandomShortLink(url)
}

func (db *DB) GetShortLink(slug string) (*models.ShortLink, error) {
	return db.getShortLinkBySlug(slug)
}

func (db *DB) getShortLinkByURL(url string) (*models.ShortLink, error) {
	stmt, err := db.Connection.Prepare(getByURLStmt)
	if err != nil {
		log.Printf("[Error] getShortLinkByURL - db.Connection.Prepare: %v", err.Error())
		return nil, err
	}
	defer stmt.Close()
	return db.getShortLink(url, stmt)
}

func (db *DB) getShortLinkBySlug(slug string) (*models.ShortLink, error) {
	stmt, err := db.Connection.Prepare(getBySlugStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return db.getShortLink(slug, stmt)
}

func (db *DB) getShortLink(searchParam string, stmt *sql.Stmt) (*models.ShortLink, error) {
	var uuid int64
	var dbSlug string
	var dbURL string
	var createdAt int64
	err := stmt.QueryRow(searchParam).Scan(&uuid, &dbSlug, &dbURL, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("[Error] getShortLink - stmt.QueryRow: %v", err.Error())
		return nil, err
	}
	return &models.ShortLink{
		UUID: uuid,
		Slug: dbSlug,
		URL: dbURL,
		CreatedAt: createdAt,
	}, nil
}

func (db *DB) makeRandomShortLink(url string) (*models.ShortLink, error) {
	stmt, err := db.Connection.Prepare(insertStmt)
	if err != nil {
		log.Printf("[Error] makeRandomShortLink - db.Connection.Prepare: %v", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var shortLinkUUID int64
	createdAt := time.Now().Unix()
	err = stmt.QueryRow(url, time.Now().Unix()).Scan(&shortLinkUUID)
	if err != nil {
		log.Printf("[Error] makeRandomShortLink - stmt.QueryRow: %v", err.Error())
		return nil, err
	}

	// Convert UUID to base62 and save back to db as the slug
	base62UUID, _ := lib.ConvertBase(shortLinkUUID, uuidConversionBase)
	_, err = db.Connection.Exec(addSlugStmt, *base62UUID, shortLinkUUID)
	if err != nil {
		log.Printf("[Error] makeRandomShortLink - db.Connection.Exec: %v", err.Error())
		return nil, err
	}

	return &models.ShortLink{
		UUID: shortLinkUUID,
		Slug: *base62UUID,
		URL: url, 
		CreatedAt: createdAt,
	}, nil
}

func (db *DB) makeCustomShortLink(url string, slug string) (*models.ShortLink, error) {
	stmt, err := db.Connection.Prepare(insertCustomStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var shortLinkUUID int64
	createdAt := time.Now().Unix()
	err = stmt.QueryRow(url, slug, createdAt).Scan(&shortLinkUUID)
	if err != nil {
		log.Printf("[Error] makeCustomShortLink - stmt.QueryRow: %v", err.Error())
		return nil, err
	}

	return &models.ShortLink{
		UUID: shortLinkUUID,
		Slug: slug,
		URL: url, 
		CreatedAt: createdAt,
	}, nil
}
