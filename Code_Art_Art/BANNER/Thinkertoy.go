package BANNER

import (
	"os"
	"strings"
)

func Thinkertoy(Input string) string {
	text := ""
	Tab1 := strings.Split(Input, "\r\n")
	d := strings.Count(Input, "\r\n")
	File, err := os.ReadFile("Code_Art_Art/Files/thinkertoy.txt")
	Tab := strings.Split(string(File), "\r\n")
	if err != nil {
		return "error 500"
	}
	for index := 0; index <= d; index++ {
		for i := 0; i < 8; i++ {
			for _, word := range Tab1[index] {
				if word > 127 {
					if i > 6 {
						return "error 400"
					}
				} else {
					text += Tab[(((word-31)*9)-8)+rune(i)]
				}
			}
			if Tab1[index] != "" {
				text += "\r\n"
			}
		}
		if Tab1[index] == "" && index != 0 {
			text += "\r\n"
		}
		if Tab1[index] == "" && index == 0 && float32(len(Input))/2 != float32(d) {
			text = "\r\n"
		}
	}
	return text
}
