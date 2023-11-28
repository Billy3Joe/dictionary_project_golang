// dictionary/handlers.go
package dictionary

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

// HandleAdd gère la requête d'ajout d'une entrée au dictionnaire.
func HandleAdd(w http.ResponseWriter, r *http.Request) {
    // Créer une instance de Dictionary
    d := New("dictionary.json")

    // Utilisez la méthode Add de la struct Dictionary
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
func HandleDefine(w http.ResponseWriter, r *http.Request) {
    // Créer une instance de Dictionary
    d := New("dictionary.json")

    // Utilisez la méthode Get de la struct Dictionary
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
func HandleRemove(w http.ResponseWriter, r *http.Request) {
    // Créer une instance de Dictionary
    d := New("dictionary.json")

    // Utilisez la méthode Remove de la struct Dictionary
    vars := mux.Vars(r)
    word := vars["word"]
    d.Remove(word)
    fmt.Println("Word removed:", word)
    fmt.Fprintf(w, "Word removed successfully!")
}

// HandleList gère la requête pour obtenir la liste de tous les mots dans le dictionnaire.
func HandleList(w http.ResponseWriter, r *http.Request) {
    // Créer une instance de Dictionary
    d := New("dictionary.json")

    // Utilisez la méthode List de la struct Dictionary
    d.List()
}
