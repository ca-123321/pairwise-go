package main

import (
	"encoding/json"
	"fmt"
  "flag"
	"pairwise/PPmaker"
  "pairwise/hexmaker"
	// "pairwise/fracmaker"
	//	"pairwise/texmaker"
)

func main() {
  order := flag.Int("order", 5, "Order of the projective plane")
  format := flag.Bool("format", false, "Format with commas for python/etc")
  shuffle := flag.Bool("shuffle", false, "Permutes cards randomly, use with fracmaker")
  show := flag.Bool("show", false, "Displays to stdout")
  color := flag.Bool("color", true, "Render in color")

  // TODO: fix shuffle
  flag.Parse()

	// Creates the projective plane
	p := PPmaker.MakePP(*order, *shuffle)
  if *show {
    display(p, *format)
  }
  // TODO: Integrate solver at this point to get one elem/column
  hexmaker.MakeHexagon(p, *color)
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
