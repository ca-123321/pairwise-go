// Generates a Projective Plane of given order
package PPmaker

// import "math/rand" // used for shuffling
import (
  "encoding/json"
  "io/ioutil"
  //"os/exec"
  "fmt"
  "math/rand"
)

// Takes (power of prime TODO: check & error) order n, returns PP: n^2+n+1 arrays of n+1
func MakePP(order int, arrange string) [][]int {
	p := [][]int{}
	coords := HomogCoords(order)

	for _, val := range coords {
		card := []int{}
		for idx2, val2 := range coords {
			if DotProd(val, val2)%order == 0 {
				card = append(card, idx2+1)
			}
		}
		p = append(p, card)
	}
  // Write json - needed for arranging by the solver
  // just do it every time because it's cheap -- could move it out later
  file, _ := json.Marshal(p)
  _ = ioutil.WriteFile("PPmaker/PP.json", file, 0644)

  // Return after arranging
  p = ArrangePP(p, arrange)
	return p
}

func ArrangePP(p [][]int, arrangement string) [][]int {
  switch {

  // Shuffles rows of p randomly
  case arrangement == "shuffle":
    for _, row := range p {
      rand.Shuffle(len(row), func(j, k int) { row[j], row[k] = row[k], row[j] })
    }
    return p

  // Returns p arranged by the python solver, one elem/col
  // TODO: Broken currently, fix!
  case arrangement == "solve":
    // run solver.py
    // read PP.json, output arrangedPP.json
		fmt.Println("Calling solver is currently broken")

    // after the above works, should read the json then return p
    return p

  // No arrangement, just return it
  default:
    return p
  }
}
