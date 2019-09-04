package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jaakidup/project/service"
)

// Service ...
var Service = service.CreateService()

// NewWebServer ...
func NewWebServer(name string, listenAddress string) *WebServer {
	return &WebServer{
		Name:          name,
		ListenAddress: listenAddress,
	}
}

// WebServer ...
type WebServer struct {
	Name          string
	ListenAddress string
}

// Serve ...
func (ws *WebServer) Serve() {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)     // POST /users
		r.Get("/", getAllUsers)     // GET /users
		r.Get("/{id}", getUserByID) // GET /users/ouqw-423452345-asdfasfda
	})
	// http.ListenAndServe(":3333", r)
	log.Println("Starting " + ws.Name + " on : " + ws.ListenAddress)
	log.Fatalln(http.ListenAndServe(ws.ListenAddress, r))

}
