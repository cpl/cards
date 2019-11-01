package sets

import "cpl.li/go/cards/pkg/cards"

var StandardSwissGerman = cards.Set{
	Faces: []cards.CardFace{
		cards.NewCardFace(6, " 6"),
		cards.NewCardFace(7, " 7"),
		cards.NewCardFace(8, " 8"),
		cards.NewCardFace(9, " 9"),
		cards.NewCardFace(10, " B"),
		cards.NewCardFace(11, " U"),
		cards.NewCardFace(12, " O"),
		cards.NewCardFace(13, " K"),
	},
	Suits: []cards.CardSuit{
		cards.NewCardSuit("Rosen", "🏵", 0),
		cards.NewCardSuit("Schellen", "🔔", 1),
		cards.NewCardSuit("Eicheln", "🌰", 2),
		cards.NewCardSuit("Schilten", "🛡", 3),
	},
}
