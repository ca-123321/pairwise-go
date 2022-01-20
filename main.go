package main

import (
	"encoding/json"
	"fmt"
  "os"
  "strconv"
	"pairwise/PPmaker"
	// "pairwise/fracmaker"
	//	"pairwise/texmaker"
)

func main() {
  order, _ := strconv.Atoi(os.Args[1])

	shuffle := true // permutes each card randomly,use with fractionmaker TODO: fix the shuffle
	format := false// formats with commas for use with python/etc

	// Creates the projective plane
	p := PPmaker.MakePP(order, shuffle)
  display(p, format)

}

// puts commas between values and arrays for using in python/etc
func display(deck [][]int, format bool) {
	if format {
    for i := 0; i <= len(deck)-1; i++ {
		  card, _ := json.Marshal(deck[i])
		  fmt.Printf("%v", string(card))
		  fmt.Printf(",\n")
    }
  } else {
    for i := 0; i < len(deck); i++ {
      fmt.Printf("%v\n", deck[i])
    }
	}
}
