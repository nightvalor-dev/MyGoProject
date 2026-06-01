package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid id")
		return 0, false
	}
	return id, true
}
