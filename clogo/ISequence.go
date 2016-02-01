package clogo

type ISequence interface {

	//Essential seq functions
	First() interface{}
	Rest() ISequence
	Cons() ISequence
}
