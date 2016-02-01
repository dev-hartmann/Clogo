package clogo

//Useful functions on sequences
func Map(data ISequence, lambda func(interface{}) bool) ISequence {
	return nil
}
func Filter(data ISequence, lambda func(interface{}) bool) ISequence {
	return nil
}

func Reduce(data ISequence, lambda func(interface{}) bool) ISequence {
	return nil
}

func IsCollection(data interface{}) bool {
	return false
}

func IsSequence(data interface{}) bool {
	return false
}

func Take(seq ISequence, n int) ISequence {
	return nil
}

func Drop(seq ISequence, n int) ISequence {
	return nil
}

func TakeWhile(seq ISequence, predicate func(interface{}) bool) ISequence {
	return nil
}
func DropWhile(seq ISequence, predicate func(interface{}) bool) ISequence {
	return nil
}

func Sort(seq ISequence) ISequence {
	return nil
}

func SortBy(predicate interface{}) ISequence {
	return nil
}
