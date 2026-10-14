package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func modifierTexte(texte string) string {
	re := regexp.MustCompile(`(\d+)\s*\((hex|bin)\)`)

	resultat := re.ReplaceAllStringFunc(texte, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		nombreStr := submatches[1]
		typeConv := submatches[2]

		_, err := strconv.Atoi(nombreStr)
		if err != nil {
			return match
		}

		switch typeConv {
		case "hex":
			nombre, err := strconv.ParseInt(nombreStr, 16, 64)
			if err != nil {
				return match
			}
			return strconv.FormatInt(nombre, 10)

		case "bin":
			nombre, err := strconv.ParseInt(nombreStr, 2, 64)
			if err != nil {
				return match
			}
			return strconv.FormatInt(nombre, 10)

		default:
			return match
		}
	})
	return resultat
}

func ReplaceHexBin() {
	texteInitial := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."

	texteModifie := modifierTexte(texteInitial)

	fmt.Println("Texte original :", texteInitial)
	fmt.Println("Texte modifié  :", texteModifie)
}
