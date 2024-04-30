package api

import (
	"github.com/artemsmotritel/uni-architecture-lab4/service"
	"log"
	"net/http"
)

type LibraryServer struct {
	Logger      *log.Logger
	Service     *service.Service
	ListenAddrs string
	mux         *http.ServeMux
}

func NewLibraryServer(logger *log.Logger, service *service.Service, address string) *LibraryServer {
	server := &LibraryServer{
		Logger:      logger,
		Service:     service,
		ListenAddrs: address,
	}

	server.mux = server.getRouter()

	return server
}

func (s *LibraryServer) Listen() error {
	server := http.Server{
		Addr:    s.ListenAddrs,
		Handler: s.mux,
	}

	return server.ListenAndServe()
}

func (s *LibraryServer) getRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", s.handleAddBook)
	mux.HandleFunc("GET /books", s.handleGetBooks)
	mux.HandleFunc("DELETE /books/{id}", s.handleRemoveBook)
	mux.HandleFunc("GET /books/{id}", s.handleGetBook)
	mux.HandleFunc("POST /books/{id}/lend", s.handleLendBook)
	mux.HandleFunc("POST /books/{id}/return", s.handleReturnBook)

	mux.HandleFunc("GET /authors", s.handleGetAuthors)

	return mux
}
