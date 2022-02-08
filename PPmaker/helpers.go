package PPmaker

// Gives homogeneous coordinates for specified order
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

// Determines if a certain triple should be included in the homogeneous coords or not
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
