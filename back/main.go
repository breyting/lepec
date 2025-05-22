package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Servir le fichier index.html
	http.Handle("/", http.FileServer(http.Dir("../front/html")))

	http.HandleFunc("/run", runHandler)
	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur : %v\n", err)
	}
}
