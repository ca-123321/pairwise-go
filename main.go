package main

import (
	"encoding/json"
	"fmt"
	"pairwise/PPmaker"
	"pairwise/fracmaker"
	//	"pairwise/texmaker"
)

func main() {
	order := 3      // needs to be powers of primes
	shuffle := true // permutes each card randomly,use with fractionmaker TODO: fix the shuffle
	print := true   // print projective plane to stdout
	format := false // formats with commas for use with python/etc

	// Creates the projective plane
	p := PPmaker.MakePP(order, shuffle)
	if print {
		if format {
			printFormat(p)
		} else {
			for i := 0; i < len(p); i++ {
				fmt.Printf("%v\n", p[i])
			}
		}
	}

	// TODO: change this to an error
	// TODO: Move this stuff
	// if order != 5 {
	// 	fmt.Printf("Fractionmaker needs order=5\n")
	// } else {
	// 	fractions := fractionMaker(p)
	// 	fmt.Println(fractions)
	// }
}

// Uncomment to show fraction-equivalance cards
// TODO: add the shuffling
// fracSets := fracmaker.MakeFractions()
// for baseFrac, val := range fracSets {
// 	fmt.Println(baseFrac+1, val)
// }

// puts commas between values and arrays for using in python/etc
func printFormat(deck [][]int) {
	for i := 0; i <= len(deck)-1; i++ {
		card, _ := json.Marshal(deck[i])
		fmt.Printf("%v", string(card))
		fmt.Printf(",\n")
	}
}

func fractionMaker(p [][]int) [31][6]fracmaker.Fraction {
	f := fracmaker.Fracky(p)
	return f
}
