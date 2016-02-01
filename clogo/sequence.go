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
}

type Sequence struct {
	head SequenceElement

}

func NewSequence(first SequenceElement) *Sequence {
	return &Sequence{
		head: first,
	}
}

func(s *Sequence)First() interface{} {
	return s.head
}


type SequenceElement struct {
	Value interface{}
	Next  SequenceElement
}

func NewSequenceElement(value interface{}, next interface{}) *SequenceElement {
	return &SequenceElement{
		Value: value,
	}
}

