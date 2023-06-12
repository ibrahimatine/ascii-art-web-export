package BANNER

import (
	"os"
	"strings"
)

func Standard(Input string) string {
	text := ""
	Tab1 := strings.Split(Input, "\r\n")
	d := strings.Count(Input, "\r\n")
	File, err := os.ReadFile("Code_Art_Art/Files/standard.txt")
	Tab := strings.Split(string(File), "\n")
	if err != nil {
		return "error 500"
	}
	for index := 0; index <= d; index++ {
		for i := 0; i < 8; i++ {
			for _, word := range Tab1[index] {
				if word > 127 {
					if i > 6 {
						//log.Fatal("error :  verifiez votre argument (", Input, ") \"", string(word), "\" n'appartient pas aux caracteres man ascii attendue")
						return "error 400"
					}
				} else {
					text += Tab[(((word-31)*9)-8)+rune(i)]
				}
			}
			if Tab1[index] != "" {
				text += "\n"
			}
		}
		if Tab1[index] == "" && index != 0 {
			text += "\n"
		}
		if Tab1[index] == "" && index == 0 && float32(len(Input))/2 != float32(d) {
			text = "\n"
		}
	}
	return text
}
