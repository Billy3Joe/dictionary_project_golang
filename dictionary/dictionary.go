// dictionary/dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

// Entry représente une entrée dans le dictionnaire avec un mot et sa définition.
type Entry struct {
	Word       string `json:"mot"`       // Le mot associé à l'entrée, utilisé comme clé dans la map
	Definition string `json:"definition"` // La définition du mot
}

// String retourne une représentation textuelle de l'entrée (la définition du mot).
func (e Entry) String() string {
	return e.Definition
}

// Dictionary représente un dictionnaire qui associe des mots à leurs définitions.
type Dictionary struct {
	filename string             // Nom du fichier dans lequel le dictionnaire est stocké
	entries  map[string]Entry  // Map pour stocker les entrées du dictionnaire avec des mots comme clés
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
    // Utilisez le mot comme clé dans la map pour assurer l'unicité des mots
    d.entries[word] = Entry{Word: word, Definition: definition}
    d.syncToFile() // Synchronise le dictionnaire avec le fichier après l'ajout
}

// Get récupère l'entrée associée à un mot dans le dictionnaire.
func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, fmt.Errorf("mot introuvable")
	}
	return entry, nil
}

// Remove supprime un mot et sa définition du dictionnaire.
func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	d.syncToFile() // Synchronise le dictionnaire avec le fichier après la suppression
}

// List affiche tous les mots et leurs définitions dans le dictionnaire.
func (d *Dictionary) List() {
	fmt.Println("Mots dans le dictionnaire:")
	for word, entry := range d.entries {
		fmt.Printf("Mot: %s\nDéfinition: %s\n\n", word, entry.String())
	}
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
	if err := d.saveToFile(); err != nil {
		fmt.Println("Erreur lors de la synchronisation avec le fichier:", err)
	}
}
