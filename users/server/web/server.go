package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turao/topics/lib/web/middleware"
	apiV1 "github.com/turao/topics/users/api/v1"
)

type server struct {
	*http.Server
	userService apiV1.Users
}

func NewServer(userService apiV1.Users) *server {
	router := mux.NewRouter()
	headerValidator := middleware.HeaderValidator(
		middleware.HeaderExists("x-user-uuid"),
		middleware.HeaderExists("x-tenancy"),
	)
	router.Use(mux.MiddlewareFunc(headerValidator))

	s := &server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		userService: userService,
	}

	s.registerRoutes(router)

	return s
}

func (s *server) registerRoutes(router *mux.Router) {
	router.HandleFunc("/user/{id}", s.handleGetUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}", s.handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/user", s.handleRegisterUser).Methods("POST")
	http.Handle("/", router)
}

func (s *server) handleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := s.userService.GetUserInfo(r.Context(), apiV1.GetUserInfoRequest{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s *server) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var request apiV1.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := s.userService.RegisterUser(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s *server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := s.userService.DeleteUser(r.Context(), apiV1.DeleteUserRequest{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
