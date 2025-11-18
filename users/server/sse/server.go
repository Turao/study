package sse

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/turao/topics/lib/web/middleware"
	apiV1 "github.com/turao/topics/users/api/v1"
)

type server struct {
	*http.Server
	usersStreamService apiV1.UsersStream
}

func NewServer(usersStreamService apiV1.UsersStream) *server {
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
		usersStreamService: usersStreamService,
	}

	s.registerRoutes(router)

	return s
}

// registerRoutes registers the routes for the web server
func (s *server) registerRoutes(router *mux.Router) {
	router.HandleFunc("/users", s.handleUsers).Methods("GET")
	http.Handle("/", router)
}

func appendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func (s *server) handleUsers(w http.ResponseWriter, r *http.Request) {
	appendHeaders(w)
	ctx := r.Context()

	keepAliveTicker := time.NewTicker(15 * time.Second)
	defer keepAliveTicker.Stop()
	response, err := s.usersStreamService.StreamUsers(ctx, apiV1.StreamUsersRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		select {
		case user := <-response.Users:
			log.Println("received a new user", user)
			return
		case <-keepAliveTicker.C:
			log.Println("tick")
			return
		case <-ctx.Done():
			log.Println("context cancelled")
			http.Error(w, "context-exceeded", http.StatusRequestTimeout)
			return
		}
	}
}
