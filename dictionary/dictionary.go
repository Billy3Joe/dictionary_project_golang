// dictionary/dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	Definition string `json:"definition"`
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	filename string
	entries  map[string]Entry
}

func New(filename string) *Dictionary {
	d := &Dictionary{
		filename: filename,
		entries:  make(map[string]Entry),
	}
	if err := d.loadFromFile(); err != nil {
		fmt.Println("Error loading from file:", err)
	}
	return d
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{Definition: definition}
	d.syncToFile()
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, fmt.Errorf("word not found")
	}
	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	d.syncToFile()
}

func (d *Dictionary) List() {
	fmt.Println("Words in the dictionary:")
	for word, entry := range d.entries {
		fmt.Printf("Word: %s\nDefinition: %s\n\n", word, entry.String())
	}
}

func (d *Dictionary) saveToFile() error {
	file, err := os.Create(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Utiliser l'option Indent pour formater le JSON avec une indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // Utiliser quatre espaces pour l'indentation

	if err := encoder.Encode(d.entries); err != nil {
		return err
	}

	return nil
}

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

func (d *Dictionary) syncToFile() {
	if err := d.saveToFile(); err != nil {
		fmt.Println("Error syncing to file:", err)
	}
}