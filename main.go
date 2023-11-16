// main.go
package main

import (
	"fmt"
	"sort"
)

// Dictionary représente une carte de mots et de définitions.
type Dictionary map[string]string

// Add ajoute un mot et sa définition au dictionnaire.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get renvoie la définition d'un mot du dictionnaire.
func (d Dictionary) Get(word string) (string, bool) {
	definition, found := d[word]
	return definition, found
}

// Remove supprime un mot du dictionnaire.
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List renvoie une liste triée des mots et de leurs définitions.
func (d Dictionary) List() []string {
	var wordList []string
	for word := range d {
		wordList = append(wordList, word)
	}
	sort.Strings(wordList)
	return wordList
}

func main() {
	// Créez une nouvelle instance de Dictionary.
	myDictionary := make(Dictionary)

	// Ajoutez des mots et des définitions.
	myDictionary.Add("gopher", "Un animal mignon")
	myDictionary.Add("go", "Un langage de programmation")
	myDictionary.Add("map", "Un magasin clé-valeur")

	// Affichez la définition d'un mot spécifique.
	definition, found := myDictionary.Get("gopher")
	if found {
		fmt.Printf("Définition de 'gopher' : %s\n", definition)
	} else {
		fmt.Println("Mot non trouvé dans le dictionnaire.")
	}

	// Supprimez un mot de la map.
	myDictionary.Remove("go")

	// Obtenez la liste triée des mots et de leurs définitions.
	wordList := myDictionary.List()
	fmt.Println("Liste triée des mots :")
	for _, word := range wordList {
		fmt.Printf("%s : %s\n", word, myDictionary[word])
	}
}
