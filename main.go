package main

import (
	"net/http"
)

func main() {
	// Servir le fichier index.html
	http.Handle("/", http.FileServer(http.Dir("./")))
	// Démarrer le serveur sur le port 8080
	http.ListenAndServe(":8080", nil)
}
