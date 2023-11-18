// main.go
package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
	"strings"
)

func main() {
	d := dictionary.New("dictionary.json")
	reader := bufio.NewReader(os.Stdin)

	// Channels pour les opérations d'ajout et de suppression
	addChan := make(chan dictionary.Entry)
	removeChan := make(chan string)

	// Goroutine pour traiter les opérations d'ajout
	go func() {
		for entry := range addChan {
			d.Add(entry.Word, entry.Definition)
			fmt.Println("Word added successfully!")
		}
	}()

	// Goroutine pour traiter les opérations de suppression
	go func() {
		for word := range removeChan {
			d.Remove(word)
			fmt.Println("Word removed successfully!")
		}
	}()

	for {
		fmt.Println("1. Add Word")
		fmt.Println("2. Define Word")
		fmt.Println("3. Remove Word")
		fmt.Println("4. List Words")
		fmt.Println("5. Exit")

		fmt.Print("Choose an action (1-5): ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			actionAdd(reader, addChan)
		case 2:
			actionDefine(d, reader)
		case 3:
			actionRemove(reader, removeChan)
		case 4:
			actionList(d)
		case 5:
			close(addChan)
			close(removeChan)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please choose a number between 1 and 5.")
		}
	}
}

func actionAdd(reader *bufio.Reader, addChan chan<- dictionary.Entry) {
	fmt.Print("Enter word: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = strings.TrimSpace(word)

	fmt.Print("Enter definition: ")
	definition, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	definition = strings.TrimSpace(definition)

	// Envoyer les données au channel d'ajout
	addChan <- dictionary.Entry{Word: word, Definition: definition}
}

func actionRemove(reader *bufio.Reader, removeChan chan<- string) {
	fmt.Print("Enter word to remove: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = strings.TrimSpace(word)

	// Envoyer le mot au channel de suppression
	removeChan <- word
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word to define: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1] // Remove the newline character

	// Utiliser une goroutine pour éviter le blocage du programme pendant la définition
	go func() {
		entry, err := d.Get(word)
		if err != nil {
			fmt.Println("Word not found.")
		} else {
			fmt.Println("Definition:", entry.String())
		}
	}()
}

func actionList(d *dictionary.Dictionary) {
	// Utiliser une goroutine pour éviter le blocage du programme pendant la liste
	go func() {
		d.List()
	}()
}
