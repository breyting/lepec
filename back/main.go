package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type CodeRequest struct {
	Code string `json:"code"`
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	// Autoriser les requêtes cross-origin (si frontend séparé)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire le body JSON
	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Créer un dossier temporaire
	tmpDir, err := os.MkdirTemp("", "code-runner-*")
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir) // Nettoyage du dossier

	// Chemin du fichier
	scriptPath := filepath.Join(tmpDir, "script.py")

	// Écrire le code dedans
	if err := os.WriteFile(scriptPath, []byte(req.Code), 0644); err != nil {
		http.Error(w, "Erreur d'écriture", http.StatusInternalServerError)
		return
	}

	// Exécuter le conteneur Docker
	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", tmpDir),
		"-w", "/app",
		"python:3",
		"python", "script.py",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Même si erreur, on renvoie stdout + stderr
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func main() {
	// Servir le fichier index.html
	http.Handle("/", http.FileServer(http.Dir("../front")))

	http.HandleFunc("/run", runHandler)
	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
