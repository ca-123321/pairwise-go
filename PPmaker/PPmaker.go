// Generates a Projective Plane of given order
package PPmaker

// import "math/rand" // used for shuffling

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
	return ArrangePP(p, "")
}

func ArrangePP(p [][]int, arrangement string) [][]int {
  switch {
  case arrangement == "shuffle":
    return p // this should shuffle
  case arrangement == "solved":
    return p // this should return one elem/col using the python solver
  default:
    return p // no arrangement, just return it
  }
}
