package clogo

type ISequence interface {
	First() interface{}
	Rest() ISequence
	Cons() ISequence
	Concat() ISequence

	Empty() bool

	Take(n int) ISequence
	Drop(n int) ISequence

	TakeWhile(func(interface{}) bool) ISequence
	DropWhile(func(interface{}) bool) ISequence

	Sort() ISequence
	SortBy(predicate interface{}) ISequence

	//Useful functions on sequences
	Map(func(interface{}) bool) ISequence
	Filter(func(interface{}) bool) ISequence
	Reduce(func(interface{}) bool) ISequence
}

type Sequence struct {
}
