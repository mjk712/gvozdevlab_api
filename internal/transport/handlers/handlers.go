package handlers

import (
	"gvozdevLabApi/internal/service"

	"github.com/gorilla/mux"
)

type handler struct {
	service service.Service
}

type Handler interface {
	GvozdevLabRoutes(router *mux.Router)
}

func NewHandler(service service.Service) Handler {
	return &handler{service}
}

func (h *handler) GvozdevLabRoutes(router *mux.Router) {

	router.HandleFunc("/gvozdev_lab/video/{id}", h.GetVideo).Methods("GET")
	router.HandleFunc("/gvozdev_lab/add_video", h.AddVideo).Methods("POST")
	router.HandleFunc("/gvozdev_lab/delete_video/{id}", h.DeleteVideo).Methods("DELETE")

	router.HandleFunc("/gvozdev_lab/preview/{id}", h.GetPreview).Methods("GET")
	router.HandleFunc("/gvozdev_lab/add_preview", h.AddPreview).Methods("POST")
	router.HandleFunc("/gvozdev_lab/delete_preview/{id}", h.DeletePreview).Methods("DELETE")

}
