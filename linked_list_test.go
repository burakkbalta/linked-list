package main

import (
	"reflect"
	"testing"
)

func TestLinkedList_InsertAtHead(t *testing.T) {
	var root *List[int] = nil
	root = root.insertAtHead(1)
	root = root.insertAtHead(2)
	root = root.insertAtHead(3)

	expected := []int{3, 2, 1}
	values := root.valuesAsSlice()

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestLinkedList_DeleteAtHead(t *testing.T) {
	var root *List[int] = nil
	root = root.insertAtHead(1)
	root = root.insertAtHead(2)
	root = root.insertAtHead(3)

	root = root.deleteAtHead()
	root = root.deleteAtHead()

	expected := []int{1}
	values := root.valuesAsSlice()

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestLinkedList_InsertAtTail(t *testing.T) {
	var root *List[int] = nil
	root = root.insertAtTail(1)
	root = root.insertAtTail(2)
	root = root.insertAtTail(3)

	expected := []int{1, 2, 3}
	values := root.valuesAsSlice()

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestLinkedList_DeleteAtTail(t *testing.T) {
	var root *List[int] = nil
	root = root.insertAtTail(1)
	root = root.insertAtTail(2)
	root = root.insertAtTail(3)

	root = root.deleteAtTail()
	root = root.deleteAtTail()

	expected := []int{1}
	values := root.valuesAsSlice()

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestLinkedLits_CreateNewNode(t *testing.T) {
	root := createNewNode(5)

	expected := &List[int]{
		val:  5,
		next: nil,
	}

	if !reflect.DeepEqual(root, expected) {
		t.Errorf("Expected %v, got %v", expected, root)
	}
}
