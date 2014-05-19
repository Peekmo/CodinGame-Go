package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, q int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &q)

	var ext, mime string
	mimes := make(map[string]string)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &ext, &mime)
		mimes[strings.ToLower(ext)] = mime
	}

	var file string
	for i := 0; i < q; i++ {
		fmt.Scanf("%s", &file)
		if v := strings.Split(strings.ToLower(file), "."); mimes[v[len(v)-1]] != "" && len(v) > 1 {
			fmt.Println(mimes[v[len(v)-1]])
		} else {
			fmt.Println("UNKNOWN")
		}
	}
}
