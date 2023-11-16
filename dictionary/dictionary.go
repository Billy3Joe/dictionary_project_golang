// dictionary/dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	Definition string
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
	d.loadFromFile()
	return d
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
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

func (d *Dictionary) List() ([]string, map[string]Entry) {
	words := make([]string, 0, len(d.entries))
	for word, entry := range d.entries {
		words = append(words, word)
		fmt.Println("Word:", word, "Definition:", entry.String()) // For debugging
	}
	return words, d.entries
}

func (d *Dictionary) saveToFile() error {
	file, err := os.Create(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
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
