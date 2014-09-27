package main

import "fmt"
import "os"

type Pos struct {
	x int
	y int
}

type Proposition struct {
	found  bool
	killed int
	pos    *Pos
	newmap []string
}

const (
	NODE         = '@'
	PASSIVE_NODE = '#'
	EMPTY        = '.'
	KILLED       = 'X'
	KILLED3      = '3'
	KILLED2      = '2'
	KILLED1      = '1'
	BOMB         = 'B'
)

var fmap []string
var width, height int
var total int

func main() {
	fmt.Fprintln(os.Stderr, "Debug messages...")

	fmt.Scan(&width, &height)
	for i := 0; i < height; i++ {
		var mapRow string

		fmt.Scan(&mapRow)
		fmap = append(fmap, mapRow)

		for x := 0; x < len(mapRow); x++ {

			if mapRow[x] == NODE {
				total += 1
			}
		}
	}

	for {
		var rounds, bombs int
		fmt.Scan(&rounds, &bombs)

		if total != 0 {
			total -= putBomb(bombs, total)

			// Updates map
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if fmap[y][x] == KILLED {
						fmap[y] = replace(fmap[y], KILLED3, x)
					} else if fmap[y][x] == KILLED3 {
						fmap[y] = replace(fmap[y], KILLED2, x)
					} else if fmap[y][x] == KILLED2 {
						fmap[y] = replace(fmap[y], KILLED1, x)
					} else if fmap[y][x] == KILLED1 {
						fmap[y] = replace(fmap[y], EMPTY, x)
					}
				}
			}

		} else {
			fmt.Println("WAIT")
		}
	}
}

func putBomb(bombs, max int) int {
	bestProposition := &Proposition{false, 0, &Pos{0, 0}, fmap}

	for i := 0; i < height; i++ {
		c := make(chan *Proposition)

		for x := 0; x < width; x++ {
			go func(pos int) {
				strline := fmap[i]
				if strline[pos] != EMPTY {
					c <- &Proposition{found: false, killed: 0, newmap: fmap}
				} else {
					var nb int
					nmap := make([]string, len(fmap))
					copy(nmap, fmap)

					// left
					var t int = 0
					for z := pos - 1; z >= 0 && t < 3; z-- {
						if strline[z] == PASSIVE_NODE {
							break
						} else if strline[z] == NODE {
							nb += 1
							nmap[i] = replace(nmap[i], KILLED, z)
						}

						t++
					}

					// right
					t = 0
					for z := pos + 1; z < width && t < 3; z++ {
						if strline[z] == PASSIVE_NODE {
							break
						} else if strline[z] == NODE {
							nb += 1
							nmap[i] = replace(nmap[i], KILLED, z)
						}

						t++
					}

					// up
					t = 0
					for z := i - 1; z >= 0 && t < 3; z-- {
						if fmap[z][pos] == PASSIVE_NODE {
							break
						} else if fmap[z][pos] == NODE {
							nb += 1
							nmap[z] = replace(nmap[z], KILLED, pos)
						}

						t++
					}

					// down
					t = 0
					for z := i + 1; z < height && t < 3; z++ {
						if fmap[z][pos] == PASSIVE_NODE {
							break
						} else if fmap[z][pos] == NODE {
							nb += 1
							nmap[z] = replace(nmap[z], KILLED, pos)
						}

						t++
					}

					nmap[i] = replace(nmap[i], BOMB, pos)
					c <- &Proposition{true, nb, &Pos{pos, i}, nmap}
				}
			}(x)
		}

		for x := 0; x < width; x++ {
			select {
			case res := <-c:
				if res.killed > bestProposition.killed {
					if bombs == 1 && res.killed != max {
						continue
					}

					bestProposition = res
				}
			}
		}
	}

	if bestProposition.found == true {
		fmt.Println(fmt.Sprintf("%d %d", bestProposition.pos.x, bestProposition.pos.y))
		fmap = bestProposition.newmap
		return bestProposition.killed
	} else {
		fmt.Fprintln(os.Stderr, fmap)
		fmt.Println("WAIT")
		return 0
	}
}

func replace(str string, r rune, i int) string {
	out := []rune(str)
	out[i] = r
	return string(out)
}
