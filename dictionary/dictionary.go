// dictionary/dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

// Entry représente une entrée dans le dictionnaire avec un mot et sa définition.
type Entry struct {
	Word       string `json:"mot"`
	Definition string `json:"definition"`
}

// String retourne une représentation textuelle de l'entrée (la définition du mot).
func (e Entry) String() string {
	return e.Definition
}

// Dictionary représente un dictionnaire qui associe des mots à leurs définitions.
type Dictionary struct {
	filename string            // Nom du fichier dans lequel le dictionnaire est stocké
	entries  map[string]Entry // Map pour stocker les entrées du dictionnaire avec des mots comme clés
	mutex    sync.Mutex        // Mutex pour synchroniser l'accès concurrent au dictionnaire
}

// New crée une nouvelle instance de Dictionary en chargeant les données à partir d'un fichier.
func New(filename string) *Dictionary {
	d := &Dictionary{
		filename: filename,
		entries:  make(map[string]Entry),
	}
	if err := d.loadFromFile(); err != nil {
		fmt.Println("Erreur lors du chargement depuis le fichier:", err)
	}
	return d
}

// Add ajoute un mot et sa définition au dictionnaire.
func (d *Dictionary) Add(word string, definition string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Utilisez le mot comme clé dans la map pour assurer l'unicité des mots
	d.entries[word] = Entry{Word: word, Definition: definition}
	d.syncToFile() // Synchronise le dictionnaire avec le fichier après l'ajout
}

// Get récupère l'entrée associée à un mot dans le dictionnaire.
func (d *Dictionary) Get(word string) (Entry, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	entry, found := d.entries[word]
	if !found {
		return Entry{}, fmt.Errorf("mot introuvable")
	}
	return entry, nil
}

// Remove supprime un mot et sa définition du dictionnaire.
func (d *Dictionary) Remove(word string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	delete(d.entries, word)
	d.syncToFile() // Synchronise le dictionnaire avec le fichier après la suppression
}

// List affiche tous les mots et leurs définitions dans le dictionnaire.
func (d *Dictionary) List() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	fmt.Println("Mots dans le dictionnaire:")
	for word, entry := range d.entries {
		fmt.Printf("Mot: %s\nDéfinition: %s\n\n", word, entry.String())
	}
}

// RegisterHandlers enregistre les gestionnaires de routes avec le routeur Gorilla Mux.
func (d *Dictionary) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/add", d.HandleAdd).Methods("POST")
	r.HandleFunc("/define/{word}", d.HandleDefine).Methods("GET")
	r.HandleFunc("/remove/{word}", d.HandleRemove).Methods("DELETE")
	r.HandleFunc("/list", d.HandleList).Methods("GET")
}

// HandleAdd gère la requête d'ajout d'une entrée au dictionnaire.
func (d *Dictionary) HandleAdd(w http.ResponseWriter, r *http.Request) {
	var entry Entry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&entry)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	d.Add(entry.Word, entry.Definition)
	fmt.Println("Word added:", entry.Word)
	fmt.Fprintf(w, "Word added successfully!")
}

// HandleDefine gère la requête pour récupérer la définition d'un mot.
func (d *Dictionary) HandleDefine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	entry, err := d.Get(word)
	if err != nil {
		http.Error(w, "Word not found", http.StatusNotFound)
	} else {
		response, _ := json.Marshal(entry)
		fmt.Printf("Definition of %s: %s\n", word, entry.Definition)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// HandleRemove gère la requête de suppression d'une entrée du dictionnaire.
func (d *Dictionary) HandleRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	d.Remove(word)
	fmt.Println("Word removed:", word)
	fmt.Fprintf(w, "Word removed successfully!")
}

// HandleList gère la requête pour obtenir la liste de tous les mots dans le dictionnaire.
func (d *Dictionary) HandleList(w http.ResponseWriter, r *http.Request) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Utiliser json.Marshal pour encoder la réponse JSON
	response, err := json.Marshal(d.entries)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// saveToFile sauvegarde les données du dictionnaire dans un fichier au format JSON.
func (d *Dictionary) saveToFile() error {
	file, err := os.Create(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Utilise l'option Indent pour formater le JSON avec une indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // Utilise quatre espaces pour l'indentation

	if err := encoder.Encode(d.entries); err != nil {
		return err
	}

	return nil
}

// loadFromFile charge les données du dictionnaire à partir d'un fichier JSON.
func (d *Dictionary) loadFromFile() error {
	file, err := os.Open(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&d.entries); err != nil {
		return err
	}

	return nil
}

// syncToFile synchronise les données du dictionnaire avec le fichier.
func (d *Dictionary) syncToFile() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := d.saveToFile(); err != nil {
		fmt.Println("Erreur lors de la synchronisation avec le fichier:", err)
	}
}
