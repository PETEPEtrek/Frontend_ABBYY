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

func (h *Handler) getCommentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

		vars := mux.Vars(r)
		commentID, err := strconv.ParseUint(vars["comment_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing user_id to uint, Error: %v", err.Error())
			return
		}

		if authUserID != commentID {
			jsend.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		comment, err := h.userService.Get(uint(commentID))
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				jsend.Error(w, err.Error(), http.StatusNotFound)
			default:
				jsend.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Error occureed in handler.getCommentByID, Error:", err.Error())
			}
			return
		}

		jsend.Success(w, comment, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
