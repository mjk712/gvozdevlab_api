package handlers

import (
	"encoding/json"
	"fmt"
	"gvozdevLabApi/internal/models"
	"gvozdevLabApi/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) GetVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["id"]
	video, err := h.service.GetVideo(videoId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in Get Video"))
		return
	}
	res, _ := json.Marshal(video.Url)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *handler) AddVideo(w http.ResponseWriter, r *http.Request) {
	video := &models.Video{}
	utils.ParseBody(r, video)
	err := h.service.AddVideo(video.Url)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in add video"))
		return
	}
	f := fmt.Sprintf("add new video:%s", video.Url)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["id"]
	err := h.service.DeleteVideo(videoId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in delete Video"))
		return
	}
	f := fmt.Sprintf("Delete Video id-%s", videoId)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
