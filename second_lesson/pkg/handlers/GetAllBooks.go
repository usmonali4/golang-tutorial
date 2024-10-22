package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/usmonali4/gotour/pkg/mocks"
	"github.com/usmonali4/gotour/pkg/models"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if res := h.DB.Find(&books); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}