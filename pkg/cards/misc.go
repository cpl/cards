package cards

type FacingDirection int

const (
	FacingDirectionDown = iota
	FacingDirectionUp
)

type DealingMethod int

const (
	DealingMethodOneByOne = iota
	DealingMethodAllAtOnce
)

type DealOpts struct {
	Facing FacingDirection
	Method DealingMethod
}

type MoveOpts struct {
	Facing FacingDirection
}
