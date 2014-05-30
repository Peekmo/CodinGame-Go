package main

import (
	"fmt"
	"os"
)

type Piece struct {
	allowed map[string]([]string)
}

type Direction struct {
	modx int
	mody int
	out  string
}

var pieces = make(map[int]Piece)
var directions = make(map[string]Direction)

func main() {
	var x, y, ex int
	fmt.Scanf("%d %d", &x, &y)

	var cmap = make([][]int, y)
	for i := 0; i < y; i++ {
		cmap[i] = make([]int, x)

		for j := 0; j < x; j++ {
			fmt.Scanf("%d", &cmap[i][j])
		}
	}

	fmt.Scanf("%d", &ex)

	fmt.Fprintln(os.Stderr, cmap) // Map
	fill()

	for {
		var posx, posy int
		var in string
		fmt.Scanf("%d %d %s", &posx, &posy, &in)

		p := pieces[cmap[posy][posx]]

		// Iterate through outs
		for _, e := range p.allowed[in] {
			nposx, nposy := posx+directions[e].modx, posy+directions[e].mody

			// Checks if in range
			if nposx >= 0 && nposx < x && nposy >= 0 && nposy < y {
				np := pieces[cmap[nposy][nposx]]

				for k, _ := range np.allowed {
					if k == directions[e].out {
						fmt.Printf("%d %d \n", nposx, nposy)
						break
					}
				}
			}
		}
	}
}

func fill() {
	directions["TOP"] = Direction{0, -1, "DOWN"}
	directions["DOWN"] = Direction{0, 1, "TOP"}
	directions["LEFT"] = Direction{-1, 0, "RIGHT"}
	directions["RIGHT"] = Direction{1, 0, "LEFT"}

	pieces[0] = Piece{map[string]([]string){}}
	pieces[1] = Piece{map[string]([]string){"TOP": []string{"DOWN"}, "RIGHT": []string{"DOWN"}, "LEFT": []string{"DOWN"}}}
	pieces[2] = Piece{map[string]([]string){"RIGHT": []string{"LEFT"}, "LEFT": []string{"RIGHT"}}}
	pieces[3] = Piece{map[string]([]string){"TOP": []string{"DOWN"}}}
	pieces[4] = Piece{map[string]([]string){"TOP": []string{"LEFT"}, "RIGHT": []string{"DOWN"}}}
	pieces[5] = Piece{map[string]([]string){"TOP": []string{"RIGHT"}, "LEFT": []string{"DOWN"}}}
	pieces[6] = Piece{map[string]([]string){"LEFT": []string{"RIGHT"}, "RIGHT": []string{"LEFT"}}}
	pieces[7] = Piece{map[string]([]string){"TOP": []string{"DOWN"}, "RIGHT": []string{"DOWN"}}}
	pieces[8] = Piece{map[string]([]string){"LEFT": []string{"DOWN"}, "RIGHT": []string{"DOWN"}}}
	pieces[9] = Piece{map[string]([]string){"LEFT": []string{"DOWN"}, "TOP": []string{"DOWN"}}}
	pieces[10] = Piece{map[string]([]string){"TOP": []string{"LEFT"}}}
	pieces[11] = Piece{map[string]([]string){"TOP": []string{"RIGHT"}}}
	pieces[12] = Piece{map[string]([]string){"RIGHT": []string{"DOWN"}}}
	pieces[13] = Piece{map[string]([]string){"LEFT": []string{"DOWN"}}}
}
