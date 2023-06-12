package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "error 404 ", http.StatusNotFound)
		log.Println(http.StatusText(http.StatusNotFound) + " -Repertoires: Handlers -Fichiers: Home.go ")
		return
	}
	tmpl, err := template.ParseFiles("./Templates/home.tmpl")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + "Error (500)" + " -Repertoires: Handlers -fichiers: Home.go ")
		return
	}
	Error_Execution := tmpl.Execute(w, nil)
	if Error_Execution != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError) + " Error (500)" + " Erreur d'execution")
		return
	}
	log.Println("Status Home.go " + http.StatusText(http.StatusOK))
}
