// main.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"estiam/dictionary"

	"github.com/gorilla/mux"
)

func main() {
	// Créer une instance de Dictionary en utilisant la fonction New définie dans votre package
	d := dictionary.New("dictionary.json")

	// Initialiser le routeur Gorilla Mux
	r := mux.NewRouter()

	// Enregistrer les routes avec le routeur
	d.RegisterHandlers(r)

	// Routes existantes
	r.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	}).Methods("GET")

	// Utiliser le routeur Gorilla Mux
	http.Handle("/", r)

	// Démarrer le serveur HTTP
	fmt.Println("Serveur en cours d'exécution sur le port :8080")

	// Log pour indiquer que le serveur démarre
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
