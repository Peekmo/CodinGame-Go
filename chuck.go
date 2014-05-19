package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	s, _ := buffer.ReadString('\n')

	last, f := "-1", ""
	for i := 0; i < len(s)-1; i++ {
		b := fmt.Sprintf("%b", s[i])
		b = strings.Repeat("0", 7-len(b)) + b

		for _, c := range b {
			if ch := string(rune(c)); last != ch {
				switch last = ch; ch {
				case "0":
					f += " 00 "
				case "1":
					f += " 0 "
				}
			}

			f += "0"
		}
	}

	fmt.Println(strings.TrimSpace(f))
}
