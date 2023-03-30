package main

import "fmt"

func main() {

	var root *List[int] = nil
	root = root.insertAtHead(20)
	root = root.insertAtTail(10)
	root = root.insertAtTail(30)
	root = root.insertAtHead(40)
	fmt.Println("Map presentation of root : ", root.valuesAsMap())
	root = root.deleteAtHead()
	root = root.deleteAtTail()
	fmt.Println("Slice presentation of root : ", root.valuesAsSlice())
	root = root.insertAtHead(50)
	root = root.insertAtHead(60)
	root = root.deleteAtTail()
	root = root.deleteAtTail()
	fmt.Println("Slice presentation of root : ", root.valuesAsSlice())

}
