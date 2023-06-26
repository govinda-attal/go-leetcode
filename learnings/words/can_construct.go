package words

import (
	"fmt"
	"strings"
)

func CanConstruct(target string, wordBank []string) bool {
	table := make([]bool, len(target)+1)

	table[0] = true
	for i := range table {
		if !table[i] {
			continue
		}
		next := target[i:]
		for _, word := range wordBank {
			if strings.HasPrefix(next, word) {
				fmt.Println(i, next, word)
				table[i+len(word)] = true
				// next = next[i+len(word):]
			}
		}
		fmt.Println(i, table)
	}
	return table[len(target)]
}

// func CanConstruct(target string, wordBank []string) bool {
// 	variations := make([]bool, len(wordBank)+1)

// 	// variations[0] = true
// 	possible := false
// 	for i := range wordBank {
// 		curr, ok := strings.CutPrefix(target, wordBank[i])
// 		if !ok {
// 			continue
// 		}
// 		// var variation []string
// 		var j int
// 		for len(curr) != 0 && j < len(wordBank) {

// 			rem, ok := strings.CutPrefix(curr, wordBank[j])
// 			if !ok {
// 				j++
// 				continue
// 			}
// 			j = 0
// 			// variation = append(variation, curr)
// 			curr = rem
// 		}
// 		if len(curr) != 0 {
// 			continue
// 		}
// 		possible = true
// 		variations[i] = len(curr) == 0
// 	}
// 	fmt.Println(variations)
// 	return possible
// }
