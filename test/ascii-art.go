// package ascii

// import (
// 	"fmt"
// 	"strings"
// )

// // var j = len(os.Args) - 1 //j = font
// // var r = "\n"

// // var flag = strings.Split(os.Args[1], "=") //split flag into two (output and file name)
// // // readFile checks os.Args for the font, also changes r value for thinkertoy

// // func readFile() string {
// // 	if string(os.Args[j]) == "thinkertoy" {
// // 		text, _ := os.ReadFile("thinkertoy.txt")
// // 		r = "\r\n"
// // 		return string(text)
// // 	} else if string(os.Args[j]) == "shadow" {
// // 		text, _ := os.ReadFile("shadow.txt")
// // 		return string(text)
// // 	} else {
// // 		text, _ := os.ReadFile("standard.txt")
// // 		return string(text)
// // 	}
// // }

// func Ascii(str string) string {
// 	//length := len(os.Args)
// 	lines := strings.Split(string("standard.txt"), `\n`)
// 	// switch length {
// 	// case 4: // args + 1
// 	// 	s := ""
// 	// 	split := strings.Split(os.Args[2], `\n`)
// 	// 	for i := 0; i < len(split); i++ {
// 	// 		if string(split[i]) == "" {
// 	// 			fmt.Println()
// 	// 		} else {
// 	// 			for j := 0; j < 8; j++ {
// 	// 				for k := 0; k < len(string(split[i])); k++ {
// 	// 					fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
// 	// 					s += lines[int(((rune(split[i][k])-32)*9+1))+j]
// 	// 				}
// 	// 				fmt.Print(r)
// 	// 				s += r
// 	// 			}
// 	// 			err := os.WriteFile(flag[1], []byte(s), 0644)
// 	// 			if err != nil {
// 	// 				panic(err)
// 	// 			}
// 	// 		}
// 	// 	}
// 	// default:
// 	split := strings.Split(str, `\n`)
// 	r := ""
// 	for i := 0; i < len(split); i++ {
// 		if string(split[i]) == "" {
// 			fmt.Println()
// 		} else {
// 			for j := 0; j < 8; j++ {
// 				for k := 0; k < len(string(split[i])); k++ {
// 					r += lines[int(((rune(split[i][k])-32)*9+1))+j]
// 				}
// 				r += "\n"
// 			}
// 		}
// 	}
// 	return r
// }

// // }
