package cards

type CardFace struct {
	Value  int
	Symbol string
}

func NewCardFace(value int, symbol string) CardFace {
	return CardFace{
		Value:  value,
		Symbol: symbol,
	}
}

func (cf CardFace) String() string {
	return string(cf.Symbol)
}
