package clogo

import "container/list"

type Sequence struct {
	head SequenceElement
	rest Sequence

	count int
}

func NewSequence(first SequenceElement) *Sequence {
	return &Sequence{
		head:  first,
		rest:  nil,
		count: 1,
	}
}

func NewSequenceFromList(data list.List) Sequence {
	count := data.Len()

	if count > 1 {
		//get first element of list
		first := data.Front()
		//remove first element from list to get 'rest' list
		_ := data.Remove(first)

		//create array from list
		rest := make([]interface{}, count)
		for element := data.Front(); element != nil; element = element.Next() {
			rest = append(rest, element)
		}

		return createSequenceFromCollection(first, rest)
	} else if count == 1 {
		return NewSequence(data.Front())
	}

	return nil
}

func NewSequenceFromArray(data []interface{}) Sequence {
	count := len(data)

	if count > 1 {
		first := data[0]
		rest := data[1:]
		return createSequenceFromCollection(first, rest)
	} else if count == 1 {
		return NewSequence(data[0])
	}

	return nil
}

func createSequenceFromCollection(first interface{}, rest []interface{}) Sequence {
	count := 1 + len(rest)

	return &Sequence{
		first: first,
		rest:  rest,
		count: count,
	}
}

func (s *Sequence) First() interface{} {
	return s.head
}

func (s *Sequence) Next() Sequence {
	if s.count == 1 {
		return nil
	}

	return s.rest
}

func (s *Sequence) Count() int {
	return s.count
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
