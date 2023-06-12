package asciiart

import (
	"os"
	"package/Code_Art_Art/BANNER"
)

func Affichage(Input, Banner string) string {
	text := ""
	saisie := Input
	if Banner == "Standard" || Banner == "" {
		text = BANNER.Standard(saisie)
	}
	if Banner == "Shadow" {
		text = BANNER.Shadow(saisie)
	}
	if Banner == "Thinkertoy" {
		text = BANNER.Thinkertoy(saisie)
	}
	f := os.WriteFile("result.txt", []byte(text), 0777)
	if f != nil {
		return ""
	}
	return text
}
