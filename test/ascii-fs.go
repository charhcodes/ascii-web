package ascii

import (
	"fmt"
	"os"
	"strings"
)

var r = "\n"

func readFile(font string) string {
	if font == "thinkertoy" {
		text, _ := os.ReadFile("thinkertoy.txt")
		r = "\r\n"
		return string(text)
	} else if font == "shadow" {
		text, _ := os.ReadFile("shadow.txt")
		return string(text)
	} else {
		text, _ := os.ReadFile("standard.txt")
		return string(text)
	}
}

func AsciiFS(str string, font string) string {
	lines := strings.Split(string(readFile(font)), r)
	s := ""
	split := strings.Split(str, `\n`)
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
					s += lines[int(((rune(split[i][k])-32)*9+1))+j]
				}
				fmt.Print(r)
				s += r
			}
		}
	}
	return s
}
