package cards

type Set struct {
	Faces []CardFace
	Suits []CardSuit
}

func (s Set) Deck() CardList {
	deck := make([]Card, len(s.Suits)*len(s.Faces))
	count := 0
	for _, suit := range s.Suits {
		for _, face := range s.Faces {
			deck[count] = NewCard(face, suit)
			count++
		}
	}

	return deck
}
