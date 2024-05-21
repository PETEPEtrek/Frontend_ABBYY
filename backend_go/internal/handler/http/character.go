package handler

import (
	"backend_go/internal/entity"
	"clevergo.tech/jsend"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getAllCharacters(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		characters, err := h.characterService.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error when getting all characters, Error: %v\n", err.Error())
			return
		}

		jsend.Success(w, characters, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getCharacterByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		vars := mux.Vars(r)
		characterID, err := strconv.ParseUint(vars["character_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing user_id to uint, Error: %v", err.Error())
			return
		}

		character, err := h.characterService.Get(uint(characterID))
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				jsend.Error(w, err.Error(), http.StatusNotFound)
			default:
				jsend.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Error occureed in handler.getCharacterByID, Error:", err.Error())
			}
			return
		}

		jsend.Success(w, character, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
