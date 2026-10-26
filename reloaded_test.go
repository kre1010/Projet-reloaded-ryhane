package main

import "testing"

func TestChangeAtoAN(t *testing.T) {
	tests := []struct {
		name, input, want string
	}{
		{"exemple du sujet",
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"hex avec lettres majuscules", "1E (hex) files", "30 files"},
		{"hex avec lettres minuscules", "ff (hex)", "255"},
		{"hex casse mixte", "aF (hex)", "175"},
		{"hex zéro", "0 (hex)", "0"},
		{"hex grand nombre", "FFFF (hex)", "65535"},

		{"bin simple", "It has been 10 (bin) years", "It has been 2 years"},
		{"bin plus long", "1111 (bin)", "15"},
		{"bin zéro", "0 (bin)", "0"},
		{"bin invalide", "12 (bin)", "12 (bin)"},
		{"plusieurs conversions",
			"A (hex) + 101 (bin) = 15",
			"10 + 5 = 15"},
		{"deux hex consécutifs", "10 (hex) 20 (hex)", "16 32"},
		{"sans espace avant le marqueur", "10(bin)", "2"},
		{"plusieurs espaces", "10   (hex)", "16"},
		{"sans marqueur", "bonjour", "bonjour"},
		{"chaîne vide", "", ""},
		{"marqueur seul", "(hex)", "(hex)"},
		{"fin de mot non hex", "file (hex)", "file (hex)"},
		{"marqueur inconnu", "10 (oct)", "10 (oct)"},
		{"marqueur en majuscules", "10 (BIN)", "10 (BIN)"},
		{"nombre sans marqueur", "42 apples", "42 apples"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ChangeAtoAN(test.input)
			got := (test.input)
			if got != test.want {
				t.Errorf("modifierTexte(%q) = %q; want %q", test.input, got, test.want)
			}
		})
	}
}
