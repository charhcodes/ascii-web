package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Ascii(str string) string {
	split := strings.Split(str, `\n`)
	text, err := os.ReadFile("ascii-art/standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(text), "\n")
	output := ""
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					output += lines[int(((rune(split[i][k])-32)*9+1))+j]
					fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
				}
				output += "\n"
				fmt.Println()
			}
		}
	}
	return output
}
