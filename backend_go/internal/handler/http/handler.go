package handler

import (
	"backend_go/internal/entity"
	"backend_go/internal/pkg/auth"
	"backend_go/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	//service *service.Service
	userService      entity.UserService
	gameService      entity.GameService
	characterService entity.CharacterService
	peopleService    entity.PeopleService
	commentService   entity.CommentService
	auth             entity.AuthManager
}

func NewHandler(s *service.Service, auth *auth.AuthManager) *Handler {
	return &Handler{userService: s.User, gameService: s.Game, characterService: s.Character, peopleService: s.People, commentService: s.Comment, auth: auth}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.getAllGames).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/characters", h.getAllCharacters).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/people", h.getAllPeoples).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/user", h.getAllUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/{game_id}", h.getGameByID).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/characters/{character_id}", h.getCharacterByID).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/people/{people_id}", h.getPeopleByID).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/user/register", h.registerUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/login", h.loginUser).Methods(http.MethodPost, http.MethodOptions)

	r.Use(LoggingMiddleware(r))
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CustomCORSMiddleware(r))

	return r
}

func LoggingMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("Origin: %s | Forwarded: %s | Method: %s | RequestURI: %s", req.Header.Get("Origin"), req.Header.Get("Forwarded"), req.Method, req.RequestURI)

			next.ServeHTTP(w, req)
		})
	}
}

func CustomCORSMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Content-Type, Authorization, Host, Origin, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")

			next.ServeHTTP(w, req)
		})
	}
}
