package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func main() {

// 	for i := 0; i < 1000; i++ {
// 		fmt.Println(string(i))
// 		fmt.Println("////////////////////")
// 	}

// }

func main() {
	fmt.Println(strconv.Itoa(3))
}
