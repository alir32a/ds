package linkedlist

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		len: 0,
	}
}

func From[T any](src ...T) *LinkedList[T] {
	ll := New[T]()

	for _, val := range src {
		ll.Push(val)
	}

	return ll
}

func (ll *LinkedList[T]) Len() int {
	return ll.len
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList[T]) Push(value T) {
	if ll.IsEmpty() {
		ll.head = &Node[T]{value: value}
		ll.tail = ll.head
		ll.len++
		return
	}

	node := &Node[T]{value: value, prev: ll.tail}
	ll.tail.next = node
	ll.tail = node
	ll.len++
}

func (ll *LinkedList[T]) Pop() (value T) {
	if ll.IsEmpty() {
		return value
	}

	value = ll.tail.value
	ll.tail.prev.next = nil
	ll.tail = ll.tail.prev
	ll.len--

	return value
}

func (ll *LinkedList[T]) Dequeue() (value T) {
	if ll.IsEmpty() {
		return value
	}

	value = ll.head.value
	ll.head = ll.head.next
	ll.len--

	return value
}

func (ll *LinkedList[T]) Enqueue(value T) {
	if ll.IsEmpty() {
		ll.Push(value)
		return
	}

	node := &Node[T]{value: value, next: ll.head}
	ll.head.prev = node
	ll.head = node
	ll.len++
}

func (ll *LinkedList[T]) Insert(index int, value T) {
	if ll.IsEmpty() {
		ll.head = &Node[T]{value: value}
		ll.tail = ll.head
		ll.len++
		return
	}

	if index < 0 || index > ll.len {
		return
	}

	if index == ll.len {
		ll.Push(value)
		return
	}

	node := &Node[T]{value: value}
	if index == 0 {
		ll.Enqueue(value)
		return
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	node.prev = current.prev
	node.next = current
	current.prev.next = node
	current.prev = node
	ll.len++
}

func (ll *LinkedList[T]) Remove(i int) (value T) {
	if ll.IsEmpty() && i >= ll.len {
		return value
	}

	if i == 0 {
		return ll.Dequeue()
	}

	current := ll.head.next
	for k := 1; k <= i; k++ {
		if current == nil {
			return value
		}
		if k == i {
			prev := current.prev
			next := current.next

			prev.next = next
			if next != nil {
				next.prev = prev
			}
			value = current.value
		}
		current = current.next
	}

	ll.len--

	return value
}

func (ll *LinkedList[T]) Has(value T, eqFunc func(T,T) bool) bool {
	current := ll.head
	for current != nil {
		if eqFunc(current.value, value) {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList[T]) Head() T {
	return ll.head.value
}

func (ll *LinkedList[T]) Tail() T {
	return ll.tail.value
}

func (ll *LinkedList[T]) ToSlice() (res []T) {
	if ll.IsEmpty() {
		return nil
	}

	current := ll.head
	for current != nil {
		res = append(res, current.value)
		current = current.next
	}	

	return res;
}