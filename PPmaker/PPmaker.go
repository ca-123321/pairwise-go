// Generates a Projective Plane of given order
package PPmaker

// import "math/rand" // used for shuffling
import (
  "encoding/json"
  "io/ioutil"
)

// Takes (power of prime TODO: check & error) order n, returns PP: n^2+n+1 arrays of n+1
func MakePP(order int, shuffle bool) [][]int {
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
  // write json - needed for arranging by the solver
  file, _ := json.Marshal(p)
  _ = ioutil.WriteFile("PPmaker/PP.json", file, 0644)

  // return after arranging
  // TODO: implement arrangements
	return ArrangePP(p, "")
}

func ArrangePP(p [][]int, arrangement string) [][]int {
  switch {
  // shuffles rows of p randomly
  case arrangement == "shuffle":
    return p 
  // returns p arranged by the python solver, one elem/col
  case arrangement == "solved":
    return p 
  // no arrangement, just return it
  default:
    return p 
  }
}
