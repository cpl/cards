package cards

type CardSuit struct {
	Name   string
	Symbol string
	Value  int
}

func NewCardSuit(name, symbol string, value int) CardSuit {
	return CardSuit{
		Name:   name,
		Symbol: symbol,
		Value:  value,
	}
}

func (cs CardSuit) String() string {
	return string(cs.Symbol)
}
