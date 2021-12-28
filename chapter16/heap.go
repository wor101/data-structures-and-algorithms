package main

import "fmt"

type Heap struct {
	data []int
}

func (h *Heap) root_node() int {
	return h.data[0]
}

func (h *Heap) last_node() int {
	return h.data[len(h.data)-1]
}

func (h *Heap) left_child_index(index int) int {
	newIndex := (index * 2) + 1

	if newIndex > (len(h.data) - 1) {
		return -1
	}
	return newIndex
}

func (h *Heap) right_child_index(index int) int {
	newIndex := (index * 2) + 2

	if newIndex > (len(h.data) - 1) {
		return -1
	}
	return newIndex
}

func (h *Heap) parent_index(index int) int {
	if index <= 0 {
		return 0
	}
	return (index - 1) / 2
}

func (h *Heap) insert(value int) {
	h.data = append(h.data, value)
	new_node_index := len(h.data) - 1

	for {
		// continue as long as new node is not root and its avlue is greater than its parent node value
		if new_node_index > 0 && h.data[new_node_index] > h.data[h.parent_index(new_node_index)] {
			//swap new node with parent node
			h.data[h.parent_index(new_node_index)], h.data[new_node_index] = h.data[new_node_index], h.data[h.parent_index(new_node_index)]

			// update new node index
			new_node_index = h.parent_index(new_node_index)
		} else {
			break
		}
	}
}

func (h *Heap) has_greater_child(index int) bool {
	var left_child_exists bool
	var right_child_exists bool
	var left_child_value int
	var right_child_value int

	if h.left_child_index(index) != -1 {
		left_child_exists = true
		left_child_value = h.data[h.left_child_index(index)]
	}

	if h.right_child_index(index) != -1 {
		right_child_exists = true
		right_child_value = h.data[h.right_child_index(index)]
	}

	trickle_node_value := h.data[index]

	if (left_child_exists && left_child_value > trickle_node_value) || (right_child_exists && right_child_value > trickle_node_value) {
		fmt.Println("Left Value:", left_child_value)
		fmt.Println("Right Value:", right_child_value)
		fmt.Println("Trickle Value:", trickle_node_value)
		return true
	} else {
		return false
	}
}

func (h *Heap) calculate_larger_child_index(index int) int {
	// if no right child
	if h.right_child_index(index) > len(h.data) {
		return h.left_child_index(index)
	}

	// if right child value is greater than left child value
	if h.data[h.right_child_index(index)] > h.data[h.left_child_index(index)] {
		return h.right_child_index(index)
	} else {
		return h.left_child_index(index)
	}
}

func (h *Heap) delete() {
	// delete root and replace with last node
	h.data = append([]int{h.data[len(h.data)-1]}, h.data[1:]...)

	// track index of trickle node
	trickle_node_index := 0

	// trickle node down until no longer has a child greater than it
	for {
		if h.has_greater_child(trickle_node_index) {
			// save larger child index in variable
			larger_child_index := h.calculate_larger_child_index(trickle_node_index)

			// swap trickle node with larger child node
			h.data[trickle_node_index], h.data[larger_child_index] = h.data[larger_child_index], h.data[trickle_node_index]

			// update trickle node index
			trickle_node_index = larger_child_index
		} else {
			break
		}
	}
}

func main() {
	heap := Heap{[]int{100, 88, 25, 87, 16, 8, 12, 86, 50, 2, 15, 3}}

	// root node
	fmt.Println("root_node method:")
	fmt.Println(heap.root_node())
	fmt.Println("")

	// last node
	fmt.Println("last_node method:")
	fmt.Println(heap.last_node())
	fmt.Println("")

	//insert node
	fmt.Println("insert method:")
	fmt.Println("Before:", heap.data)
	heap.insert(40)
	fmt.Println("After:", heap.data)
	fmt.Println("")

	// right child index
	fmt.Println("right_child_index method:")
	fmt.Printf("RightChildIndex: %v", heap.right_child_index(2))
	fmt.Println("")

	// left child index
	fmt.Println("left_child_index method:")
	fmt.Printf("LeftChildIndex: %v", heap.left_child_index(2))
	fmt.Println("")

	// has greater child
	fmt.Println("has_greater_child method:")
	heap.has_greater_child(4)
	fmt.Println("")

	// delete node
	fmt.Println("delete method:")
	fmt.Println("Before: ", heap.data)
	heap.delete()
	fmt.Println("After:", heap.data)
	fmt.Println("")

}
