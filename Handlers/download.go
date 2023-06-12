package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func Download(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/download" {
		http.Error(w, "error 404 ", http.StatusNotFound)
		log.Println(http.StatusText(http.StatusNotFound) + " -Repertoires: Handlers -Fichiers: Home.go ")
		return
	}
	tmpl, err := template.ParseFiles("./Templates/ascii-art.tmpl")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + "Error (500)" + " -Repertoires: Handlers -fichiers: Ascii-Art.go  ")
		return
	}
	Format := r.FormValue("fileformat")

	T, _ := os.Open("result." + Format)
	defer T.Close()

	File, _ := T.Stat()
	Filesize := File.Size()

	sFileSize := strconv.Itoa(int(Filesize))

	w.Header().Set("Content-Disposition", "attachment; filename=fichier."+Format)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", sFileSize)

	io.Copy(w, T)

	tmpl.Execute(w, nil)
}
