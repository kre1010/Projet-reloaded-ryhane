package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	contenu, err := os.ReadFile(os.Args[1])

	texteInitial := string(contenu)
	texteTransform := strings.ToUpper(texteInitial)

	err = os.WriteFile(os.Args[2], []byte(texteTransform), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println("Transformation réussite !")

}
