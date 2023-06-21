package controller

import (
	"dictionary/repository"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Controller struct {
	Repository *repository.Repository
}

type RequestBody struct {
	Word       string
	Definition string
}

func (c *Controller) Create(writer http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()

	var requestBody RequestBody
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if len(requestBody.Word) == 0 {
		http.Error(writer, "You should specify a word", http.StatusBadRequest)
		return
	}

	if len(requestBody.Definition) == 0 {
		http.Error(writer, "You should specify a definition", http.StatusBadRequest)
		return
	}

	c.Repository.Add(requestBody.Word, requestBody.Definition)

	_, err := c.Repository.Get(requestBody.Word)
	if err != nil {
		http.Error(writer, "There was an issue when adding your word to the dictionary", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Definition added successfully for word `" + requestBody.Word + "`"))
}

func (c *Controller) Get(writer http.ResponseWriter, request *http.Request) {
	word := chi.URLParam(request, "word")

	defer request.Body.Close()

	if len(word) == 0 {
		http.Error(writer, "You should specify a word", http.StatusBadRequest)
		return
	}

	entry, err := c.Repository.Get(word)
	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(writer).Encode(entry)
}
