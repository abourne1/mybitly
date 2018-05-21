package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github/abourne1/mybitly/db"
	"github/abourne1/mybitly/lib"
	"github/abourne1/mybitly/models"
)

type Handler struct {
	DB *db.DB
}

func New(db *db.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) Link(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	var rb models.LinkReqBody
	err = decoder.Decode(&rb)
	if err != nil {
		// TODO: return 400 Bad Request
		panic(err)
	}

	shortLink, err := h.DB.MakeShortLink(rb.URL, rb.Slug)
	if err != nil {
		// TODO: return 500
		panic(err)
	}

	resp, err := json.Marshal(shortLink)
	if err != nil {
		// TODO: return 500
		panic(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// TODO: Consider making this three separate endpoints
// Instructions say "Provide a route for returning stats on a given short link"
// I'm not sure if I have to make ONE route or if I can make many routes
func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
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
		count, err := h.DB.GetShortLinkVisitCount(rb.Slug, rb.StartTime, rb.EndTime)
		if err != nil {
			// TODO: return 500
			panic(err)
		}
		c <- count
	}(countChan)

	go func(c chan map[string]int64) {
		histogram, err := h.DB.GetShortLinkVisitHistogram(rb.Slug, rb.StartTime, rb.EndTime)
		if err != nil {
			// TODO: return 500
			panic(err)
		}
		c <- histogram
	}(histChan)

	go func(c chan int64) {
		createdAt, err := h.DB.GetShortLinkCreationDate(rb.Slug)
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

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug, ok := mux.Vars(r)["slug"]
	if !ok {
		// TODO: return 400 bad request
		panic("400 - Bad Request")
	}

	shortLink, err := h.DB.GetShortLinkBySlug(slug)
	if err != nil {
		// TODO: return 404 not found
		panic(err)
	}
	if shortLink == nil {
		// TODO: return 404 not found
		return
	}

	go h.DB.MakeShortLinkVisit(shortLink.Slug)

	restOfURL := lib.GetURISuffix(r.URL.RequestURI())
	http.Redirect(w, r, shortLink.URL + restOfURL, http.StatusSeeOther)
}
