package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage : programme <fichier source> <fichier destination>")
		return
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v", err)
		return
	}

	texte := string(contenu)
	mots := strings.Fields(texte)
	mots = Appliquer(mots)
	texte = strings.Join(mots, " ")
	texte = clearponctuation(texte)
	texte, _ = ChangeAtoAN(texte)
	err = os.WriteFile(os.Args[2], []byte(texte), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}
	fmt.Println(texte)
	fmt.Println("Transformation réussite !")
}
