package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Jeswyrne/pig-latin/helper/response"
	"github.com/Jeswyrne/pig-latin/repository"

	"github.com/Jeswyrne/pig-latin/controller/utils"
	"github.com/Jeswyrne/pig-latin/models"
)

const vowels = "aeiou"

type PigLatinInterface interface {
	PostHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
}

type PigLatin struct {
	repo repository.DatabaseInterface
}

var _ PigLatinInterface = &PigLatin{}

func NewPigLatin(db repository.DatabaseInterface) PigLatinInterface {
	return &PigLatin{repo: db}
}

func (pl *PigLatin) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		response.Error(w,
			http.StatusNoContent,
			errors.New("method not allowed"),
		)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	var input models.BodyJsonRequest
	err = json.Unmarshal(body, &input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	output := pl.transform(input.Data)
	db_data := models.SaveObject{
		Input:  input.Data,
		Output: output,
	}

	res, err := pl.repo.Save(db_data)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.ToJSON(w, http.StatusOK, res)
}

func (pl *PigLatin) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		response.Error(w,
			http.StatusNoContent,
			errors.New("method not allowed"),
		)
		return
	}

	res, err := pl.repo.List()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.ToJSON(w, http.StatusOK, res)
}

func (pl *PigLatin) transform(in string) string {
	lower := strings.ToLower(in)
	words := strings.Split(lower, " ")

	for i, w := range words {
		words[i] = pl.translate(w)
	}

	return strings.Join(words, " ")

}

func (pl *PigLatin) translate(input string) string {
	if !utils.CheckHasVowels(input, vowels) {
		return input[2:] + input[:2] + "ay"
	}

	var idx int
	for k, v := range input {
		if strings.Contains(vowels, string(v)) {
			idx += k
			break
		}
	}

	if idx == 0 {
		return input + "way"
	}

	return input[idx:] + input[:idx] + "ay"
}
