package datafile

import (
	"fmt"
	"strings"
)

const (
	StandardHash   = "ac85e83127e49ec42487f272d9b9db8b"
	ShadowHash     = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	ThinkertoyHash = "86d9947457f6a41a18cb98427e314ff8"
)

func Asciiart(text, banner string) (string, int) {
	for _, alter := range text {
		if (rune(alter) < rune(32) || rune(alter) > rune(127)) && alter != rune(13) && alter != rune(10) {
			fmt.Println("ERROR: non printable character")
			return "", 400
		} else if alter == rune(13) || alter == rune(10) {
			text = text + string(alter)
		}
	}
	filename := ""

	if banner == "shadow" {
		filename = "shadow"
		if filename != "shadow" {
			return "", 500
		}
	} else if banner == "thinkertoy" {
		filename = "thinkertoy"
		if filename != "thinkertoy" {
			return "", 500
		}
	} else if banner == "standard" {
		filename = "standard"
		if filename != "standard" {
			return "", 500
		}
	}

	filename = "assets/fonts/" + filename + ".txt"

	// Check for hash
	if GetHash(filename) == StandardHash || GetHash(filename) == ShadowHash || GetHash(filename) == ThinkertoyHash {
		asciiLines, err := GetStrings(filename)
		if err != nil {
			fmt.Println("ERROR: can't read file")
			return "", 500
		}
		asciiMap := make(map[rune][]string)
		x := 1
		y := 9
		for key := 32; key < 127; key++ {
			asciiMap[rune(key)] = asciiLines[x:y]
			x = x + 9
			y = y + 9
		}
		res := ""
		textstr := strings.ReplaceAll(text, "\n", "\n")
		arg := strings.Split(textstr, "\r\n")
		for i, v := range arg {
			if v == "" {
				arg[i] = ""
			}
		}
		newline := forNewLines(arg)
		for w := 0; w < len(arg); w++ {
			if newline && w == len(arg)-1 {
				break
			}
			if arg[w] != "" {
				for i := 0; i < 8; i++ {
					for _, ch := range arg[w] {
						res = res + asciiMap[ch][i]
					}
					res = res + string(rune(10))
				}
			} else if arg[w] == "" {
				res = res + string(rune(10))
			}
		}
		return res, 200
	} else {
		fmt.Println("Error: Wrong hash!")
		return "", 500
	}
}

func forNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}
