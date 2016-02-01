package clogo

type Sequence struct {
	head SequenceElement
}

func NewSequence(first SequenceElement) *Sequence {
	return &Sequence{
		head: first,
	}
}

func (s *Sequence) First() interface{} {
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
