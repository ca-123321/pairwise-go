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
  order := flag.Int("order", 2, "Order of the projective plane")
  show := flag.String("show", "", "'text', or 'array', (default no output)")
  color := flag.Bool("color", false, "Render in color")
  hexDeck := flag.Bool("hexdeck", false, "Send the PP to hexmaker and make a set of images")
  arrange := flag.String("arrange", "", "'shuffle', 'solve', 'test', 'order', or default ''")
  flag.Parse()

	// Create and arrange the projective plane
	p := PPmaker.MakePP(*order, *arrange)

  switch *show {
    case "array":
      for i := 0; i < len(p); i++ {
        fmt.Printf("%v\n", p[i])
      }

    case "text":
      for i := 0; i < len(p); i++ {
        fmt.Printf("Card %d: ", i+1)
        for j := 0; j < len(p[i]); j++ {
          fmt.Printf("%v ", p[i][j])
        }
        fmt.Println()
      }
    }

  if *hexDeck {
    hexmaker.MakeHexagon(p, *color)
  }
}

