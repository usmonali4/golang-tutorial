package handlers

import (
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	"log"
	"net/http"

	"github.com/usmonali4/gotour/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	if res := h.DB.Create(&book); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}