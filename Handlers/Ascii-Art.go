package handlers

import (
	"log"
	"net/http"
	asciiart "package/Code_Art_Art"
	"text/template"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "error 404 ", http.StatusNotFound)
		log.Println(http.StatusText(http.StatusNotFound) + " -Repertoires: Handlers -Fichiers: Ascii-Art.go ")
		return
	}
	tmpl, err := template.ParseFiles("./Templates/ascii-art.tmpl")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + "Error (500)" + " -Repertoires: Handlers -fichiers: Ascii-Art.go  ")
		return
	}
	Input := r.FormValue("Saisie")
	Banner := ""
	if Input == "" {
		http.Error(w, "error 400 ", http.StatusBadRequest)
		log.Println(http.StatusText(http.StatusBadRequest), " -Repertoires: Handlers -Fichiers: Ascii-Art.go")
		return
	}
	Banner = r.FormValue("Banner")
	resultat := asciiart.Affichage(Input, Banner)
	type Output struct {
		Input  string
		Banner string
		Sortie string
	}
	p := Output{
		Input:  Input,
		Banner: Banner,
		Sortie: resultat,
	}
	if resultat == "error 400" {
		http.Error(w, "error 400 ", http.StatusBadRequest)
		log.Println(http.StatusText(http.StatusBadRequest), " -Repertoires: Handlers -Fichiers: Ascii-Art.go")
		return
	}
	if resultat == "error 500" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + "Error (500) : fichier introuvable" + " -Repertoires: Handlers -fichiers: Ascii-Art.go ")
		return
	}
	Error_Execution := tmpl.Execute(w, p)
	if Error_Execution != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + "Error (500)" + " -Repertoires: Handlers -fichiers: Ascii-Art.go ")
		return
	}
	log.Println("Status Ascii-Art.go: "+http.StatusText(http.StatusOK)+" Input:", Input, " Banner: ", Banner)
}
