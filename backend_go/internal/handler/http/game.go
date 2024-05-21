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

func (h *Handler) getAllGames(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		games, err := h.gameService.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error when getting all games, Error: %v\n", err.Error())
			return
		}

		jsend.Success(w, games, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *Handler) getGameByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		vars := mux.Vars(r)
		gameID, err := strconv.ParseUint(vars["game_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing user_id to uint, Error: %v", err.Error())
			return
		}

		game, err := h.gameService.Get(uint(gameID))
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				jsend.Error(w, err.Error(), http.StatusNotFound)
			default:
				jsend.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Error occureed in handler.getGameByID, Error:", err.Error())
			}
			return
		}

		jsend.Success(w, game, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
