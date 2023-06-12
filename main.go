package main

import (
	"fmt"
	"net/http"
	handlers "package/Handlers"
)

func main() {
	const port = ":8080"
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art", handlers.Ascii)
	http.HandleFunc("/download", handlers.Download)
	fmt.Println("link: http://localhost:8080 Server started port", port)
	http.ListenAndServe(port, nil)
}
