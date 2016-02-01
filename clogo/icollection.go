package clogo

type ICollection interface {

	//Essential collection functions
	IsEmpty() bool
	IsCollection() bool

	Count() int
	Seq() ISequence
}
