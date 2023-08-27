package handlers

import (
	"encoding/json"
	"fmt"
	"gvozdevLabApi/internal/models"
	"gvozdevLabApi/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) GetPreview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	previewId := vars["id"]
	preview, err := h.service.GetPreview(previewId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in Get Preview"))
		return
	}
	res, _ := json.Marshal(preview.Url)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *handler) AddPreview(w http.ResponseWriter, r *http.Request) {
	preview := &models.Preview{}
	utils.ParseBody(r, preview)
	err := h.service.AddPreview(preview.Url)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in add preview"))
		return
	}
	f := fmt.Sprintf("add new preview:%s", preview.Url)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *handler) DeletePreview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	previewId := vars["id"]
	err := h.service.DeletePreview(previewId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in delete Preview"))
		return
	}
	f := fmt.Sprintf("Delete preview id-%s", previewId)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
