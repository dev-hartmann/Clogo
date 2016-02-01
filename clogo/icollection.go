package clogo

type ICollection interface {

	//Essential collection functions
	Empty() bool
	IsCollection() bool

	Count() int
	Seq() ISequence

}