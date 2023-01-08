package api

import (
	"io/ioutil"
	"net/http"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	// imageName := r.URL.Query().Get("/user/img/profile")        // mengambil nama image dari query url
	// fileBytes, err := ioutil.ReadFile("./images/" + imageName) // membaca file image menjadi bytes
	fileBytes, err := ioutil.ReadFile("./assets/images/img-avatar-default.png") // membaca file image menjadi bytes
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("File not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes) // menampilkan image sebagai response
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	// TODO: answer here

	api.dashboardView(w, r)
}
