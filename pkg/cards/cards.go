package cards

import (
	"fmt"
	"math/rand"
	"sort"
)

type CardList []Card

func (cards CardList) Shuffle() CardList {
	dest := make([]Card, len(cards))
	perm := rand.Perm(len(cards))

	for i, v := range perm {
		dest[v] = cards[i]
	}

	return dest
}

func (cards CardList) Sort() CardList {
	sort.Slice(cards, func(i, j int) bool {
		cardI := cards[i]
		cardJ := cards[j]

		if cardI.Suit.Value == cardJ.Suit.Value {
			return cardI.Face.Value < cardJ.Face.Value
		}

		return cardI.Suit.Symbol < cardJ.Suit.Symbol
	})
	return cards
}

func (cards CardList) Deal(count int, opts DealOpts, cardLists ...*CardList) (CardList, error) {
	if count <= 0 {
		return cards, fmt.Errorf("card count must be greater than 0, not %d", count)
	}

	l := len(cards)
	need := count * len(cardLists)

	if l < need {
		return cards, fmt.Errorf("deck has %d cards, need %d cards to deal %d cards to %d cardLists",
			l, need, count, len(cardLists))
	}

	switch opts.Method {
	case DealingMethodAllAtOnce:
		for index := range cardLists {
			*cardLists[index] = append(
				*cardLists[index],
				cards[l-count*(index+1):l-count*index].SetFacing(opts.Facing)...)
		}
	case DealingMethodOneByOne:
		dealt := 0
		for iter := 0; iter < count; iter++ {
			for index := range cardLists {
				cards[l-dealt-1].Facing = opts.Facing
				*cardLists[index] = append(
					*cardLists[index],
					cards[l-dealt-1])
				dealt++
			}
		}
	}

	return cards[:l-need], nil
}

func (cards CardList) Peak(count int) (CardList, error) {
	if count <= 0 {
		return nil, fmt.Errorf("card count must be greater than 0, not %d", count)
	}
	return cards[:count], nil
}

func (cards CardList) SetFacing(direction FacingDirection) CardList {
	for index := range cards {
		cards[index].Facing = direction
	}
	return cards
}

func (cards *CardList) Move(index int, other *CardList, opts MoveOpts) {
	if index < 0 || index >= len(*cards) {
		return
	}

	(*cards)[index].Facing = opts.Facing

	*other = append(*other, (*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
