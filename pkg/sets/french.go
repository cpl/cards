package sets

import "cpl.li/go/cards/pkg/cards"

var StandardFrench = cards.Set{
	Faces: []cards.CardFace{
		cards.NewCardFace(2, " 2"),
		cards.NewCardFace(3, " 3"),
		cards.NewCardFace(4, " 4"),
		cards.NewCardFace(5, " 5"),
		cards.NewCardFace(6, " 6"),
		cards.NewCardFace(7, " 7"),
		cards.NewCardFace(8, " 8"),
		cards.NewCardFace(9, " 9"),
		cards.NewCardFace(10, "10"),
		cards.NewCardFace(11, " J"),
		cards.NewCardFace(12, " Q"),
		cards.NewCardFace(13, " K"),
		cards.NewCardFace(14, " A"),
	},
	Suits: []cards.CardSuit{
		cards.NewCardSuit("Hearts", "♥", 0),
		cards.NewCardSuit("Diamonds", "♦", 1),
		cards.NewCardSuit("Clubs", "♣", 2),
		cards.NewCardSuit("Spades", "♠", 3),
	},
}
