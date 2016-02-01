package clogo

type ICollection interface {

	//Essential collection functions
	Empty() bool
	Count() int
	Seq() ISequence
}