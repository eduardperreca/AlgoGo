package main

import (
	"fmt"
	"unicode"
)

func main() {
	var stringa string
	fmt.Scan(&stringa)

	checkUp := 0
	cehckLow := 0
	for _, k := range stringa {
		if unicode.IsUpper(k) {
			checkUp++
		} else {
			cehckLow++
		}
	}

	if checkUp > cehckLow {
		for _, k := range stringa {
			fmt.Print(string(unicode.ToUpper(k)))
		}
	} else {
		for _, k := range stringa {
			fmt.Print(string(unicode.ToLower(k)))
		}
	}

}
