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
  solver := flag.Bool("solver", false, "Use the solver to ensure every element is in exactly one column")
  hexDeck := flag.Bool("hexdeck", true, "Send the PP to hexmaker and make a set of images")

  // TODO: fix shuffle
  flag.Parse()

  if *solver {
    fmt.Println("The solver isn't implemented yet")
  }
	// Creates the projective plane
	p := PPmaker.MakePP(*order, *shuffle)
  if *show {
    display(p, *format)
  }
  
  if *hexDeck {
    hexmaker.MakeHexagon(p, *color)
  }
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
