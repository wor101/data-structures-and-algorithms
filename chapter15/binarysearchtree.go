package main

import "fmt"

type TreeNode struct {
	value      int
	leftChild  *TreeNode
	rightChild *TreeNode
}

func search(searchValue int, node *TreeNode) *TreeNode {
	fmt.Println("Current Node:", node.value)
	if node == nil {
		return nil
	}

	if searchValue == node.value {
		return node
	} else if searchValue > node.value {
		return search(searchValue, node.rightChild)
	} else {
		return search(searchValue, node.leftChild)
	}
}

func insert(value int, node *TreeNode) {
	if value < node.value {
		if node.leftChild == nil {
			node.leftChild = &TreeNode{value, nil, nil}
		} else {
			insert(value, node.leftChild)
		}
	} else if value > node.value {
		if node.rightChild == nil {
			node.rightChild = &TreeNode{value, nil, nil}
		} else {
			insert(value, node.rightChild)
		}
	}
}

func lift(node *TreeNode, nodeToDelete *TreeNode) *TreeNode {
	if node.leftChild != nil {
		node.leftChild = lift(node.leftChild, nodeToDelete)
		return node
	} else {
		nodeToDelete.value = node.value
		return node.rightChild
	}
}

func traverse_and_print(node *TreeNode) {
	if node == nil {
		return
	}
	traverse_and_print(node.leftChild)
	fmt.Println("Node:", node.value)
	traverse_and_print(node.rightChild)
}

func delete(valueToDelete int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	} else if valueToDelete < node.value {
		node.leftChild = delete(valueToDelete, node.leftChild)
		return node
	} else if valueToDelete > node.value {
		node.rightChild = delete(valueToDelete, node.rightChild)
		return node
	} else if valueToDelete == node.value {
		if node.leftChild == nil {
			return node.rightChild
		} else if node.rightChild == nil {
			return node.leftChild
		} else {
			node.rightChild = lift(node.rightChild, node)
			return node
		}
	}
	return nil
}

func greatest_value(node *TreeNode) *TreeNode {
	if node.rightChild == nil {
		return node
	}

	return greatest_value(node.rightChild)
}

func main() {

	// row 3
	node7 := TreeNode{4, nil, nil}
	node8 := TreeNode{11, nil, nil}
	node9 := TreeNode{30, nil, nil}
	node10 := TreeNode{40, nil, nil}
	node11 := TreeNode{52, nil, nil}
	node12 := TreeNode{61, nil, nil}
	node13 := TreeNode{82, nil, nil}
	node14 := TreeNode{95, nil, nil}

	// row 2
	node3 := TreeNode{10, &node7, &node8}
	node4 := TreeNode{33, &node9, &node10}
	node5 := TreeNode{56, &node11, &node12}
	node6 := TreeNode{89, &node13, &node14}

	// row 1
	node1 := TreeNode{25, &node3, &node4}
	node2 := TreeNode{75, &node5, &node6}

	//Root
	root := TreeNode{50, &node1, &node2}
	fmt.Println(root)
	fmt.Println("")

	// search
	fmt.Println("search method:")
	fmt.Println(search(40, &root))
	fmt.Println("")

	// insert
	fmt.Println("insert method:")
	insert(45, &root)
	fmt.Println(search(45, &root))
	fmt.Println("")

	// traverse
	fmt.Println("traverse_and_print method:")
	traverse_and_print(&root)
	fmt.Println("")

	// delete
	fmt.Println("delete/lift method:")
	delete(56, &root)
	traverse_and_print(&root)
	fmt.Println("")

	// greatest value
	fmt.Println("greatest_value method:")
	fmt.Println(greatest_value(&root))
	fmt.Println("")

}
