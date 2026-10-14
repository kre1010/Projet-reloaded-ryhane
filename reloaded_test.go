package main
 
import "testing"
 
func TestModifierTexte(t *testing.T) {
	tests := []struct {
		name, input, want string
	}{
		{"exemple du sujet",
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"hex avec lettres", "1E (hex) files", "30 files"},
		{"bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"sans marqueur", "bonjour", "bonjour"},
		{"chaîne vide", "", ""},
		{"bin invalide", "12 (bin)", "12 (bin)"},
	}
 
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := modifierTexte(test.input)
			if got != test.want {
				t.Errorf("modifierTexte(%q) = %q; want %q", test.input, got, test.want)
			}
		})
	}
}
