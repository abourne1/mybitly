package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"

	"github/abourne1/mybitly/models"
)

func prepareStatsRequest(r *http.Request) (*models.StatsReqBody, error) {
	decoder := json.NewDecoder(r.Body)
	var err error
	var rb models.StatsReqBody
	err = decoder.Decode(&rb)
	if err != nil {
		log.Printf("[Error] Stats - decoder.Decode: %v", err.Error())
		return nil, err
	}
	if rb.Slug == nil {
		return nil, fmt.Errorf("Stats request body must contain 'slug'")
	}
	return &rb, nil
}

// StatsCount writes the count of all visits to a given slug to the response
func (h *Handler) StatsCount(w http.ResponseWriter, r *http.Request) {
	rb, err := prepareStatsRequest(r)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}
	count, err := h.DB.GetShortLinkVisitCount(*rb.Slug, rb.StartTime, rb.EndTime)
	if err != nil {
		log.Printf("[Error] Stats - h.DB.GetShortLinkVisitCount: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp, err := json.Marshal(models.StatsCountRespBody{
		Count: count,
	})
	if err != nil {
		log.Printf("[Error] Stats - h.DB.GetShortLinkVisitHistogram: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	writeResponse(w, http.StatusCreated, resp)
}

// StatsHistogram writes the per-day count of all visits to a given slug to the response
func (h *Handler) StatsHistogram(w http.ResponseWriter, r *http.Request) {
	rb, err := prepareStatsRequest(r)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}
	histogram, err := h.DB.GetShortLinkVisitHistogram(*rb.Slug, rb.StartTime, rb.EndTime)
	if err != nil {
		log.Printf("[Error] Stats - h.DB.GetShortLinkVisitHistogram: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	resp, err := json.Marshal(models.StatsHistRespBody{
		Histogram: histogram,
	})
	if err != nil {
		log.Printf("[Error] Stats - json.Marshal: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	writeResponse(w, http.StatusCreated, resp)
}

// StatsCreatedAt writes the date a slug was first created to the response
func (h *Handler) StatsCreatedAt(w http.ResponseWriter, r *http.Request) {
	rb, err := prepareStatsRequest(r)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}
	createdAt, err := h.DB.GetShortLinkCreationDate(*rb.Slug)
	if err != nil {
		log.Printf("[Error] Stats - h.DB.GetShortLinkVisitHistogram: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	resp, err := json.Marshal(models.StatsCreatedAtRespBody{
		CreatedAt: createdAt,
	})
	if err != nil {
		log.Printf("[Error] Stats - json.Marshal: %v", err.Error())
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}
	writeResponse(w, http.StatusCreated, resp)
}
