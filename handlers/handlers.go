package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github/abourne1/mybitly/db"
	"github/abourne1/mybitly/lib"
	"github/abourne1/mybitly/models"
)

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	var rb models.LinkReqBody
	err = decoder.Decode(&rb)
	if err != nil {
		// TODO: return 400 Bad Request
		panic(err)
	}

	if rb.Slug != nil {
		err := db.MakeShortLink(rb.URL, rb.Slug)
		if err != nil {
			// TODO: return 403 Forbidden (can't override existing URLs for now)
			panic(err)
		}
		// TODO: return 200
		return
	}

	err = db.MakeShortLink(rb.URL, nil)
	if err != nil {
		// TODO: return 500
		panic(err)
	}
}

// TODO: Consider making this three separate endpoints
// Instructions say "Provide a route for returning stats on a given short link"
// I'm not sure if I have to make ONE route or if I can make many routes
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	var rb models.StatsReqBody
	err = decoder.Decode(&rb)
	if err != nil {
		// TODO: return 400 Bad Request
		panic(err)
	}

	// Make stats calls in parallel and write results
	countChan := make(chan int64)
	histChan := make(chan map[string]int64)
	createChan := make(chan int64)

	go func(c chan int64) {
		count, err := db.GetShortLinkVisitCount(rb.Slug, rb.StartTime, rb.EndTime)
		if err != nil {
			// TODO: return 500
			panic(err)
		}
		c <- count
	}(countChan)

	go func(c chan map[string]int64) {
		histogram, err := db.GetShortLinkVisitHistogram(rb.Slug, rb.StartTime, rb.EndTime)
		if err != nil {
			// TODO: return 500
			panic(err)
		}
		c <- histogram
	}(histChan)

	go func(c chan int64) {
		createdAt, err := db.GetShortLinkCreationDate(rb.Slug)
		if err != nil {
			// TODO: return 500
			panic(err)
		}
		c <- createdAt
	}(createChan)

	resp, err := json.Marshal(models.StatsRespBody{
		Count: <-countChan,
		Histogram: <-histChan,
		CreatedAt: <-createChan,
	})
	if err != nil {
		// TODO: return 500
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	slug, ok := mux.Vars(r)["slug"]
	if !ok {
		// TODO: return 400 bad request
		panic("400 - Bad Request")
	}

	shortLink, err := db.GetShortLink(slug)
	if err != nil {
		// TODO: return 404 not found
		panic(err)
	}

	go db.MakeShortLinkVisit(shortLink.Slug)

	restOfURL := lib.GetURISuffix(r.URL.RequestURI())
	http.Redirect(w, r, shortLink.URL + restOfURL, http.StatusSeeOther)
}
