package clogo

type ISequence interface {

	First() interface{}
	Rest() ISequence
	Cons() ISequence

	//Useful functions on sequences
	Map() ISequence
	Filter() ISequence
	Reduce() ISequence
}

type Sequence struct {
}
