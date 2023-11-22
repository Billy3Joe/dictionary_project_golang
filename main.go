// main.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"estiam/dictionary"

	"github.com/gorilla/mux"
)

// Entry représente une entrée dans le dictionnaire avec un mot et sa définition.
type Entry struct {
    Word       string `json:"mot"`
    Definition string `json:"definition"`
}

func main() {
    // Créer une instance de Dictionary en utilisant la fonction New définie dans votre package
    d := dictionary.New("dictionary.json")

    // Initialiser le routeur Gorilla Mux
    r := mux.NewRouter()

    // Enregistrer les routes avec le routeur
    registerRoutes(r, d)

    // Routes existantes
    r.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
        os.Exit(0)
    }).Methods("GET")

    // Utiliser le routeur Gorilla Mux
    http.Handle("/", r)

    // Démarrer le serveur HTTP
    fmt.Println("Serveur en cours d'exécution sur le port :8080")
    http.ListenAndServe(":8080", nil)
}

// registerRoutes enregistre les routes avec le routeur Gorilla Mux.
func registerRoutes(r *mux.Router, d *dictionary.Dictionary) {
    r.HandleFunc("/add", d.HandleAdd).Methods("POST")
    r.HandleFunc("/define/{word}", d.HandleDefine).Methods("GET")
    r.HandleFunc("/remove/{word}", d.HandleRemove).Methods("DELETE")
    r.HandleFunc("/list", d.HandleList).Methods("GET")

	// Routes existantes
	r.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	}).Methods("GET")

	// Utiliser le routeur Gorilla Mux
	http.Handle("/", r)

		// Démarrer le serveur HTTP
		fmt.Println("Serveur en cours d'exécution sur le port :8080")
		http.ListenAndServe(":8080", nil)
}



