package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/turao/topics/lib/web/middleware"
	"github.com/turao/topics/lib/web/sse"
	apiV1 "github.com/turao/topics/users/api/v1"
)

// server is the implementation of the web server
type server struct {
	*http.Server
	userService       apiV1.Users
	userStreamService apiV1.UsersStream
}

// NewServer creates a new web server
func NewServer(userService apiV1.Users, usersStreamService apiV1.UsersStream) *server {
	router := mux.NewRouter()
	headerValidator := middleware.HeaderValidator(
	// middleware.HeaderExists("x-user-uuid"),
	// middleware.HeaderExists("x-tenancy"),
	)
	router.Use(mux.MiddlewareFunc(headerValidator))

	s := &server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		userService:       userService,
		userStreamService: usersStreamService,
	}

	s.registerRoutes(router)

	return s
}

// registerRoutes registers the routes for the web server
func (s *server) registerRoutes(router *mux.Router) {
	router.HandleFunc("/user/{id}", s.handleGetUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}", s.handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/user", s.handleRegisterUser).Methods("POST")
	router.HandleFunc("/sse/users", s.handleSSEUsers).Methods("GET")
	http.Handle("/", router)
}

// handleGetUserInfo handles the GET request for the user info
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

// handleRegisterUser handles the POST request for the user registration
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

// handleDeleteUser handles the DELETE request for the user deletion
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

func appendSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func (s *server) handleSSEUsers(w http.ResponseWriter, r *http.Request) {
	appendSSEHeaders(w)
	ctx := r.Context()

	keepAliveTicker := time.NewTicker(1 * time.Second)
	defer keepAliveTicker.Stop()

	response, err := s.userStreamService.StreamUsers(ctx, apiV1.StreamUsersRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flusher := w.(http.Flusher)
	for {
		select {
		case user := <-response.Users:
			data, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			event := sse.Event{
				Event: "user",
				Data:  data,
				ID:    &user.ID,
			}

			w.Write([]byte(event.String()))
			flusher.Flush()
		case <-keepAliveTicker.C:
			w.Write([]byte(sse.KeepAlive))
			flusher.Flush()
		case <-ctx.Done():
			http.Error(w, "context-exceeded", http.StatusRequestTimeout)
			return
		}
	}
}
