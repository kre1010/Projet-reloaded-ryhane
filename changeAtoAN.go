package main

import (
	"errors"
	"fmt"
	"strings"
)

func ChangeAtoAN(original string) (string, error) {
	if !strings.Contains(original, "a") {
		return "", errors.New("erreur: la phrase ne contient pas de a")
	}
	return strings.ReplaceAll(original, "a", "an"), nil
}

func Affichagechangement() {
	phraseOriginal := "There is no greater agony than bearing a untold story inside you"
	phraseModifier, err := ChangeAtoAN(phraseOriginal)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(phraseModifier)
}
