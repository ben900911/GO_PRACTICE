/*
 * for index, card := range cards {
 * index: index of this element in the array
 * card : current item iterating over
 * range *** : take the slice and loop over it
 *}
 */

package main

func main() {
	cards := newDeck()
	cards.shuffle()

	//cards.saveToFile("myCards")
	//hand, remainingCards := deal(cards, 5)
	//card1 := newDeckFromFile("myCards")
	cards.print()
}
