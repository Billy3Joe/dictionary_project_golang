package main

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func NewDictionary() Dictionary {
	return make(Dictionary)
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) (string, bool) {
	definition, exists := d[word]
	return definition, exists
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() []string {
	var wordList []string
	for word := range d {
		wordList = append(wordList, word)
	}
	sort.Strings(wordList)
	return wordList
}

func main() {
	dict := NewDictionary()

	dict.Add("Go", "A statically typed compiled programming language designed at Google.")
	dict.Add("Map", "A collection type that holds key-value pairs")

	word := "Go"
	definition, exists := dict.Get(word)
	if exists {
		fmt.Printf("Definition of '%s': %s\n", word, definition)
	} else {
		fmt.Println("Word not found:", word)
	}

	dict.Remove("Map")

	fmt.Println("Words in dictionary:")
	for _, w := range dict.List() {
		def, _ := dict.Get(w)
		fmt.Printf("%s: %s\n", w, def)
	}
}
