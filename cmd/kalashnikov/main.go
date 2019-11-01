package main

import (
	"fmt"
	"math/rand"
	"time"

	"cpl.li/go/cards/pkg/sets"

	"cpl.li/go/cards/pkg/cards"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var discard, player0, player1 cards.CardList

	deck := sets.StandardFrench.Deck()
	deck = deck.Shuffle().SetFacing(cards.FacingDirectionDown)

	deck, _ = deck.Deal(4, cards.DealOpts{
		Method: cards.DealingMethodOneByOne,
		Facing: cards.FacingDirectionDown,
	}, &player0, &player1)

	fmt.Println(player0.SetFacing(cards.FacingDirectionUp))
	player0.Move(1, &discard, cards.MoveOpts{Facing: cards.FacingDirectionDown})

	fmt.Println(player0.SetFacing(cards.FacingDirectionUp))
	fmt.Println(discard.SetFacing(cards.FacingDirectionUp))
}
