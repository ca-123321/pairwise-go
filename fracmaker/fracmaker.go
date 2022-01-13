package fracmaker

import (
	"fmt"
	"math/rand"
)

type Fraction struct {
	numerator   int
	denominator int
}

func BaseFractions() [31][6]Fraction {
	fracs := [31]Fraction{
		{1, 2}, {1, 3}, {2, 3}, {1, 4}, {3, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {1, 6}, {5, 6}, {1, 7}, {2, 7},
		{3, 7}, {4, 7}, {5, 7}, {6, 7}, {1, 8}, {3, 8}, {5, 8}, {7, 8}, {1, 9}, {2, 9}, {4, 9}, {5, 9}, {7, 9}, {8, 9},
		{1, 10}, {3, 10}, {7, 10}, {9, 10},
	}

	set := [31][6]Fraction{} // each element of set is the 6 representations of the base frac

	for current := 0; current <= len(fracs)-1; current++ { // Each of our 31 base fractions
		for i := 0; i <= 5; i++ { // Each of the 6 representations
			newFrac := fracs[current]
			newFrac.numerator *= i + 1
			newFrac.denominator *= i + 1
			set[current][i] = newFrac
		}
		rand.Shuffle(len(set[current]), func(j, k int) { set[current][j], set[current][k] = set[current][k], set[current][j] })
	}

	return set
}

func Fracky(deck [][]int) [31][6]Fraction {
	order := 5

	fractions := BaseFractions()
	fractionDeck := [31][6]Fraction{} // create empty deck

	counter := 0
	cardElem := 0
	for i := 0; i <= len(deck)-1; i++ { // over every card in the deck
		for j := 0; j <= order-1; j++ { // over the elements on the card
			cardElem = deck[i][j]
			fractionDeck[cardElem-1] = AddFracToCard(fractionDeck[cardElem-1], fractions[i][counter])
			counter = (counter + 1) % 6
		}
	}

	return fractionDeck
}

// Send a card in, find the next empty spot and put a fraction there
// could generalize this later using global order
func AddFracToCard(card [6]Fraction, frac Fraction) [6]Fraction {
	for i := 0; i <= 5; i++ {
		if card[i].denominator == 0 {
			card[i] = frac
			return card
		}
	}
	fmt.Println("Something went wrong, card full?")
	return card
}

func HomogCoords(order int) [][3]int {
	triples := [][3]int{}
	max := order * order * order
	for i := 0; i < max; i++ {
		candidate := [3]int{(i / order / order) % order, (i / order) % order, i % order}
		if TestCandidate(candidate) {
			triples = append(triples, candidate)
		}
	}
	// fmt.Println(triples)
	//fmt.Println("Number remaining:", len(triples))
	//fmt.Println("n^2 + n + 1:", order*order+order+1)
	return triples
}

func TestCandidate(triple [3]int) bool {
	if triple == [3]int{0, 0, 0} {
		//	fmt.Println("Rejected:", triple)
		return false
	}
	if (triple[0] != 0) && (triple[0] != 1) { // left-most digit is not 0 or 1
		//	fmt.Println("Rejected:", triple)
		return false
	}
	if (triple[0] == 0) && (triple[1] > 1) { //left is 0, middle is 2 or more
		//	fmt.Println("Rejected:", triple)
		return false
	}
	if (triple[0] == 0) && (triple[1] == 0) && (triple[2] > 1) { // left 0, middle 0, right 2 or more
		//	fmt.Println("Rejected:", triple)
		return false
	}
	return true
}

func DotProd(v1, v2 [3]int) int {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
