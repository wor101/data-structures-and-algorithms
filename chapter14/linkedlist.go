package main

import "fmt"

type Node struct {
	data      string
	next_node *Node
}

type LinkedList struct {
	head *Node
}

type DoubleNode struct {
	data      string
	next_node *DoubleNode
	prev_node *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

func (dl *DoublyLinkedList) print_all_reverse() {
	current_node := dl.tail

	for {
		if current_node == nil {
			break
		}
		fmt.Println(current_node.data)
		current_node = current_node.prev_node
	}
}

func (l *LinkedList) read(index int) string {
	current_node := l.head

	for current_index := 0; current_index < index; current_index++ {
		if current_node.next_node == nil {
			return ""
		}
		current_node = current_node.next_node
	}
	return current_node.data
}

func (l *LinkedList) print_all() {
	current_node := l.head

	for {
		if current_node == nil {
			break
		}
		fmt.Println(current_node.data)
		current_node = current_node.next_node
	}
}

func (l *LinkedList) index_of(value string) int {
	current_node := l.head
	current_index := 0

	for {
		if current_node == nil {
			break
		}
		if current_node.data == value {
			return current_index
		}
		current_node = current_node.next_node
		current_index += 1
	}
	return -1
}

func (l *LinkedList) insert_at_index(index int, value string) {
	new_node := Node{value, nil}

	if index == 0 {
		new_node.next_node = l.head
		l.head = &new_node
		return
	}

	current_node := l.head

	for current_index := 0; current_index < (index - 1); current_index++ {
		current_node = current_node.next_node
	}

	new_node.next_node = current_node.next_node
	current_node.next_node = &new_node
}

func (l *LinkedList) delete_at_index(index int) {
	if index == 0 {
		l.head = l.head.next_node
	}

	current_node := l.head

	for current_index := 0; current_index < (index - 1); current_index++ {
		current_node = current_node.next_node
	}

	node_after_deleted := current_node.next_node.next_node
	current_node.next_node = node_after_deleted
}

func (l *LinkedList) last_element() *Node {
	current_node := l.head

	for {
		if current_node == nil {
			break
		}

		if current_node.next_node == nil {
			return current_node
		}
		current_node = current_node.next_node
	}
	return nil
}

func (l *LinkedList) reverse_list() {
	var nextNodes []Node

	currentNode := l.head

	for {
		if currentNode == nil {
			break
		}
		nextNodes = append(nextNodes, *currentNode)
		currentNode = currentNode.next_node
	}

	for index := (len(nextNodes) - 1); index >= 0; index-- {
		if index == 0 {
			nextNodes[index].next_node = nil
			break
		}
		nextNodes[index].next_node = &nextNodes[index-1]
	}
	l.head = &nextNodes[len(nextNodes)-1]
}

func main() {
	node1 := Node{"Once", nil}
	node2 := Node{"Upon", nil}
	node3 := Node{"A", nil}
	node4 := Node{"Time", nil}

	node1.next_node = &node2
	node2.next_node = &node3
	node3.next_node = &node4

	// fmt.Println(node1.data)
	// fmt.Println(node2.next_node.data)

	list := LinkedList{&node1}
	// fmt.Println(list.head.next_node.data)

	fmt.Println("read method:")
	fmt.Println(list.read(0))
	fmt.Println(list.read(1))
	fmt.Println(list.read(2))
	fmt.Println(list.read(3))
	fmt.Println("")

	fmt.Println("index_of method:")
	fmt.Println(list.index_of("Once"))
	fmt.Println(list.index_of("Upon"))
	fmt.Println(list.index_of("A"))
	fmt.Println(list.index_of("Time"))
	fmt.Println(list.index_of("bad value"))
	fmt.Println("")

	fmt.Println("insert method:")
	list.insert_at_index(1, "New Node")
	fmt.Println(list.read(0))
	fmt.Println(list.read(1))
	fmt.Println(list.read(2))
	fmt.Println(list.read(3))
	fmt.Println(list.read(4))
	fmt.Println("")

	fmt.Println("delete method:")
	list.delete_at_index(1)
	fmt.Println(list.read(0))
	fmt.Println(list.read(1))
	fmt.Println(list.read(2))
	fmt.Println(list.read(3))
	fmt.Println(list.read(4))
	fmt.Println(list.read(5))
	fmt.Println("")

	fmt.Println("print_all method:")
	list.print_all()
	fmt.Println("")

	dnode1 := DoubleNode{"grandma", nil, nil}
	dnode2 := DoubleNode{"got", nil, nil}
	dnode3 := DoubleNode{"run", nil, nil}
	dnode4 := DoubleNode{"over", nil, nil}

	dnode1.next_node = &dnode2
	dnode2.prev_node = &dnode1
	dnode2.next_node = &dnode3
	dnode3.prev_node = &dnode2
	dnode3.next_node = &dnode4
	dnode4.prev_node = &dnode3

	dlist := DoublyLinkedList{&dnode1, &dnode4}

	fmt.Println("print_all_reverse method:")
	dlist.print_all_reverse()
	fmt.Println("")

	fmt.Println("last_element method:")
	fmt.Println(list.last_element().data)
	fmt.Println("")

	fmt.Println("reverse_list method BEFORE:")
	list.print_all()
	fmt.Println("reverse_list method AFTER:")
	list.reverse_list()
	list.print_all()

}
