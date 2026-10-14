package main

import (
	"fmt"
	"os"
)

func main() {

	contenu, err := os.ReadFile(os.Args[1])

	texteInitial := string(contenu)
	test := modifierTexte(texteInitial)

	err = os.WriteFile(os.Args[2], []byte(test), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println("Transformation réussite !")
	ReplaceHexBin()

}
