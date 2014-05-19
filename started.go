package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	buffer := bufio.NewReader(os.Stdin)
	s, _ := buffer.ReadString('\n')
	words, _ := buffer.ReadString('\n')
	tokens := strings.Split(words, " ")
	w1, w2 := tokens[0], tokens[1]

	if a+b == len(strings.TrimSpace(s)) {
		fmt.Println(w1)
	} else {
		fmt.Println(w2)
	}
}
