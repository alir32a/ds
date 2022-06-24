package linkedlist

type Node[T any] struct {
	value T
	next *Node[T]
	prev *Node[T]
}

func (n Node[T]) Value() T {
	return n.value
}