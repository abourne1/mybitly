package handlers

import (
	"net/http"
	"encoding/json"
	"log"

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

func writeResponse(w http.ResponseWriter, status int, resp []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if resp != nil {
		w.Write(resp)
	}
}

func (h *Handler) Link(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var err error
	var rb models.LinkReqBody
	err = decoder.Decode(&rb)
	if err != nil {
		log.Printf("[Error] Link - decoder.Decode: %v", err.Error())
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	shortLink, err := h.DB.MakeShortLink(rb.URL, rb.Slug)
	if err != nil {
		log.Printf("[Error] Link - h.DB.MakeShortLink: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp, err := json.Marshal(shortLink)
	if err != nil {
		log.Printf("[Error] Link - json.Marshal: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	writeResponse(w, http.StatusCreated, resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug, ok := mux.Vars(r)["slug"]
	if !ok {
		log.Printf("[Error] Redirect - Unable to parse URL params")
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	shortLink, err := h.DB.GetShortLink(slug)
	if err != nil {
		log.Printf("[Error] Redirect - h.DB.GetShortLink: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	if shortLink == nil {
		log.Printf("[Error] Redirect - no short link with slug: %v", slug)
		writeResponse(w, http.StatusNotFound, nil)
		return
	}

	go func() {
		err := h.DB.MakeShortLinkVisit(shortLink.Slug)
		if err != nil {
			log.Printf("Failed to record visit: %s", err.Error())
		}
	}()

	restOfURL := lib.GetURISuffix(r.URL.RequestURI())
	http.Redirect(w, r, shortLink.URL + restOfURL, http.StatusSeeOther)
}
