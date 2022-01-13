// Generates a Projective Plane of given order
package PPmaker

import "math/rand"

// Takes (power of prime TODO: check & error) order n, returns PP: n^2+n+1 arrays of n+1
func MakePP(order int, shuffle bool) [][]int {
	PP := [][]int{}
	coords := HomogCoords(order)
	for _, val := range coords {
		card := []int{}
		for idx2, val2 := range coords {
			if DotProd(val, val2)%order == 0 {
				card = append(card, idx2+1)
			}
		}
		if shuffle {
			rand.Shuffle(len(card), func(j, k int) { card[j], card[k] = card[k], card[j] })
		}
		PP = append(PP, card)
	}
	return PP
}

//
func HomogCoords(order int) [][3]int {
	triples := [][3]int{}
	max := order * order * order
	for i := 0; i < max; i++ {
		candidate := [3]int{(i / order / order) % order, (i / order) % order, i % order}
		if TestCandidate(candidate) {
			triples = append(triples, candidate)
		}
	}
	return triples
}

// determines if a certain triple should be included in the homogeneous coords or not
func TestCandidate(triple [3]int) bool {
	switch {
	case triple == [3]int{0, 0, 0}:
		return false
	case (triple[0] != 0) && (triple[0] != 1): // left-most digit is not 0 or 1
		return false
	case (triple[0] == 0) && (triple[1] > 1): // left is 0, middle is 2 or more
		return false
	case (triple[0] == 0) && (triple[1] == 0) && (triple[2] > 1): // left 0, middle 0, right 2 or more
		return false
	default:
		return true
	}
}

// computes dot product
func DotProd(v1, v2 [3]int) int {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
