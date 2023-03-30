package main

type List[T any] struct {
	next *List[T]
	val  T
}

func createNewNode[T any](value T) *List[T] {
	return &List[T]{
		next: nil,
		val:  value,
	}
}

func (root *List[T]) insertAtHead(value T) *List[T] {
	newNode := createNewNode(value)
	newNode.next = root
	return newNode
}

func (root *List[T]) insertAtTail(value T) *List[T] {
	if root == nil {
		return root.insertAtHead(value)
	}

	curr := root
	for curr.next != nil {
		curr = curr.next
	}

	newNode := createNewNode(value)
	curr.next = newNode
	return root
}

func (root *List[T]) deleteAtHead() *List[T] {
	if root != nil {
		root = root.next
	}

	return root
}

func (root *List[T]) deleteAtTail() *List[T] {
	if root != nil {
		if root.next == nil {
			root = root.deleteAtHead()
			return root
		}

		curr := root
		for curr.next.next != nil {
			curr = curr.next
		}

		curr.next = nil
	}

	return root
}

func (root *List[T]) valuesAsMap() map[int]T {
	values := make(map[int]T)
	curr := root
	index := 0
	for curr != nil {
		values[index] = curr.val
		curr = curr.next
		index++
	}

	return values
}

func (root *List[T]) valuesAsSlice() []T {
	values := []T{}
	curr := root

	for curr != nil {
		values = append(values, curr.val)
		curr = curr.next
	}

	return values
}
