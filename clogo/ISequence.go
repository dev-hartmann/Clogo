package clogo

type ISequence interface {

	//Essential seq functions
	First() interface{}
	Rest() ISequence
	Cons() ISequence

	Concat() ISequence

	Take(n int) ISequence
	Drop(n int) ISequence

	TakeWhile(predicate func(interface{}) bool) ISequence
	DropWhile(predicate func(interface{}) bool) ISequence

	Sort() ISequence
	SortBy(predicate interface{}) ISequence
}
