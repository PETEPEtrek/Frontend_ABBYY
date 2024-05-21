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

func (h *Handler) getAllPeoples(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		peoples, err := h.peopleService.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error when getting all peoples, Error: %v\n", err.Error())
			return
		}

		jsend.Success(w, peoples, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getPeopleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		vars := mux.Vars(r)
		peopleID, err := strconv.ParseUint(vars["people_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing people_id to uint, Error: %v", err.Error())
			return
		}

		people, err := h.peopleService.Get(uint(peopleID))
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				jsend.Error(w, err.Error(), http.StatusNotFound)
			default:
				jsend.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Error occureed in handler.getPeopleByID, Error:", err.Error())
			}
			return
		}

		jsend.Success(w, people, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
