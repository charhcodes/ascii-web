package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Ascii(fnt, str string) string {
	//fmt.Printf("value %s", fnt)
	Nstring := strings.ReplaceAll(str, "\r", "")
	Nstring = strings.ReplaceAll(Nstring, "\n", `\n`)
	split := strings.Split(Nstring, `\n`)

	text, err := os.ReadFile("ascii-art/" + fnt + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	Nlines := strings.ReplaceAll(string(text), "\r", "")
	lines := strings.Split(string(Nlines), "\n")
	output := ""
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					output += lines[int(((rune(split[i][k])-32)*9+1))+j]
					//fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
				}
				output += "\n"
				//fmt.Println()
			}
		}
	}
	return output
}
