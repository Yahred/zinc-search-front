package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"zinc-index-back.com/models"
)

type SearchController struct{}

func (sc SearchController) SearchEmails(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if err != nil {
		pageSize = 10
	}

	emails, err := models.SearchEmails(query, page, pageSize)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emails)
}

func Route(router chi.Router) {
	sc := SearchController{}

	router.Get("/api/search", sc.SearchEmails)

}
