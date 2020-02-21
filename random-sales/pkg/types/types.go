package types

type RetailTransaction struct {
	Id int
	Store int
	Item string
	Amount float32
	Count int
	Tax float32
	Cost float32
}