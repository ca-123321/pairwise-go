package main

import (
	//"encoding/json"
  //"io/ioutil"
	"fmt"
  "flag"
	"pairwise/PPmaker"
  "pairwise/hexmaker"
	// "pairwise/fracmaker"
	//	"pairwise/texmaker"
)

func main() {
  // Flags
  // TODO: group arrangement flags: Shuffle, Solver, Default
  // TODO: Maybe move all the flags and checks out of main or something cleaner, a struct or something
  // TODO: fix shuffle
  order := flag.Int("order", 5, "Order of the projective plane")
  shuffle := flag.Bool("shuffle", false, "Permutes elements of the rows randomly")
  show := flag.Bool("show", false, "Displays to stdout (default false)")
  color := flag.Bool("color", false, "Render in color")
  hexDeck := flag.Bool("hexdeck", false, "Send the PP to hexmaker and make a set of images")
  arrange := flag.String("arrange", "", "'Shuffle', 'Solved', or default ''")
  flag.Parse()

	// Creates the projective plane
	p := PPmaker.MakePP(*order, *shuffle, *arrange)

  if *show {
     for i := 0; i < len(p); i++ {
       fmt.Printf("%v\n", p[i])
     }
  }

  if *hexDeck {
    hexmaker.MakeHexagon(p, *color)
  }
}

