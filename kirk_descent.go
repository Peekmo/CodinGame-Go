package main

import (
	"fmt"
)

func main() {
	var sx, sy int
	var mounts [8]int

	for {
		fmt.Scanf("%d %d", &sx, &sy)

		high, pos := 0, 0
		for i := 0; i < 8; i++ {
			fmt.Scanf("%d", &mounts[i])
			if mounts[i] > high {
				high = mounts[i]
				pos = i
			}
		}

		if sx == pos {
			fmt.Println("FIRE")
		} else {
			fmt.Println("HOLD")
		}
	}
}
