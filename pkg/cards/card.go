package cards

import "fmt"

type Card struct {
	Face CardFace
	Suit CardSuit

	Facing FacingDirection
}

func NewCard(face CardFace, suit CardSuit) Card {
	return Card{
		Face: face,
		Suit: suit,
	}
}

func (c Card) String() string {
	switch c.Facing {
	case FacingDirectionDown:
		return "facing down"
	}

	return fmt.Sprintf("%s of %s", c.Face.String(), c.Suit.String())
}
