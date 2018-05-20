package db

import (
	"time"
	"database/sql"

	_ "github.com/lib/pq"

	"github/abourne1/mybitly/models"
)

type DB struct {
	Connection *sql.DB
}

// close DB when main func terminates
func New(connection *sql.DB) *DB {
	return &DB{
		Connection: connection,
	}
}

func MakeShortLink(url string, slug *string) error {
	// TODO: implement business logic with db calls when appropriate
	return nil
}

func GetShortLink(url string) (models.ShortLink, error) {
	// TODO: remove this fixture
	return models.ShortLink{
		UUID: "1",
		Slug: "temp",
		URL: "temp",
		CreatedAt: time.Now().Unix(),
	}, nil
}

func MakeShortLinkVisit(slug string) error {
	// TODO: impolement business logic with db calls when appropriate
	return nil
}

func GetShortLinkVisitCount(slug string, start *int64, end *int64) (int64, error) {
	// TODO: implement logic w/ db calls when appropriate
	//
	// SELECT count(*) 
	// FROM shortlinks 
	// WHERE slug = slug AND createdAt BETWEEN start AND end
	return 100, nil
}

func GetShortLinkVisitHistogram(slug string, start *int64, end *int64) (map[string]int64, error) {
	// TODO: implement logic w/ db calls when appropriate
	//
	// SELECT count(*) 
	// FROM shortlinks 
	// WHERE slug = slug AND createdAt BETWEEN start AND end 
	// GROUP BY datestr
	return map[string]int64{}, nil
}

func GetShortLinkCreationDate(slug string) (int64, error) {
	// TODO: implement db logic later
	// 
	// SELECT createdAt
	// FROM shortlinks 
	// WHERE slug = slug
	// ORDER BY createdAt ASC
	// LIMIT 1
	return time.Now().Unix(), nil
}