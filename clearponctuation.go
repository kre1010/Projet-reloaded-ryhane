package main

import (
	"strings"
)

func clearponctuation(contenu string) string {
	caracteres := []rune(contenu)
	var resultat strings.Builder

	for index, caractere := range caracteres {
		if caractere == ' ' && index+1 < len(caracteres) {
			prochaincaractere := caracteres[index+1]
			if strings.ContainsRune(".,!?:;", prochaincaractere) {
				continue
			}
		}
		resultat.WriteRune(caractere)
	}
	return resultat.String()
}
