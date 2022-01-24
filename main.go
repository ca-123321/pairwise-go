package main

import (
	"encoding/json"
  "io/ioutil"
	"fmt"
  "flag"
	"pairwise/PPmaker"
  "pairwise/hexmaker"
	// "pairwise/fracmaker"
	//	"pairwise/texmaker"
)

func main() {
  // Go inserts some defaults for -help but only in some cases. 
  // TODO: Weird, look into, but not too hard. Only true-valued bools seem to
  // be automatically given a default.
  order := flag.Int("order", 5, "Order of the projective plane")
  format := flag.Bool("format", false, "Format with commas for python/etc (default false)")
  shuffle := flag.Bool("shuffle", false, "Permutes cards randomly, use with fracmaker (default false)")
  show := flag.Bool("show", false, "Displays to stdout (default false)")
  color := flag.Bool("color", false, "Render in color")
  solver := flag.Bool("solver", false, "Use the solver to ensure every element is in exactly one column")
  hexDeck := flag.Bool("hexdeck", false, "Send the PP to hexmaker and make a set of images")
  makejson := flag.Bool("json", false, "Marshal the PP and write it to a json file")
  // TODO: Various checks for incompatible flags
  // TODO: Maybe move all the flags and checks out of main or something cleaner
  // a struct or something
  // TODO: fix shuffle
  flag.Parse()

  if *solver {
    fmt.Println("The solver isn't implemented yet")
  }
	// Creates the projective plane
	p := PPmaker.MakePP(*order, *shuffle)

  if *makejson {
    file, _ := json.Marshal(p)
    _ = ioutil.WriteFile("solver/PP.json", file, 0644) 
    // TODO: dynamic filename for order, separate dir? maybe not for general
    // solver use
  }

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
