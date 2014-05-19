package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	temps, _ := reader.ReadString('\n')
	tokens := strings.Split(temps, " ")

	if len(tokens) == 0 {
		fmt.Println("0")
		os.Exit(0)
	}

	var c int = 1e5
	for i := 0; i < len(tokens); i++ {
		intval, _ := strconv.Atoi(strings.TrimSpace(tokens[i]))
		if math.Abs(float64(intval)) < math.Abs(float64(c)) || (math.Abs(float64(intval)) == math.Abs(float64(c)) && intval > c) {
			c = intval
		}
	}

	fmt.Println(c)
}
