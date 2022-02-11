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
  order := flag.Int("order", 5, "Order of the projective plane")
  show := flag.Bool("show", false, "Displays to stdout (default false)")
  color := flag.Bool("color", false, "Render in color")
  hexDeck := flag.Bool("hexdeck", false, "Send the PP to hexmaker and make a set of images")
  arrange := flag.String("arrange", "", "'Shuffle', 'Solve', or default ''")
  flag.Parse()

	// Create and arrange the projective plane
	p := PPmaker.MakePP(*order, *arrange)

  if *show {
     for i := 0; i < len(p); i++ {
       fmt.Printf("%v\n", p[i])
     }
  }

  if *hexDeck {
    hexmaker.MakeHexagon(p, *color)
  }
}

