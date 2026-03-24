package iterator

type Iterator interface {
	Next() *Video
	HasNext() bool
}
