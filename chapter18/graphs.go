package main

import (
	"fmt"
	"reflect"
)

type Queue struct {
	data []*Vertex
}

func (q *Queue) enqueue(vertex *Vertex) {
	q.data = append(q.data, vertex)
}

func (q *Queue) dequeue() *Vertex {
	dequeued := q.data[0]
	q.data = q.data[1:]
	return dequeued
}

func (q *Queue) read() *Vertex {
	return q.data[0]
}

type Vertex struct {
	value             string
	adjacent_vertices []*Vertex
}

func (v *Vertex) add_adjacent_vertex(vertex *Vertex) {
	for _, adjacent_vertex := range v.adjacent_vertices {
		if reflect.DeepEqual(vertex, adjacent_vertex) {
			return
		}
	}

	v.adjacent_vertices = append(v.adjacent_vertices, vertex)
	vertex.add_adjacent_vertex(v)
}

func (v *Vertex) print_all_adjacent() {
	for index, vertex := range v.adjacent_vertices {
		fmt.Printf("Index %v: %v", index, vertex)
		fmt.Println("")
	}
}

func (v *Vertex) dfs_traverse(vertex *Vertex, viseted_vertices map[string]bool) {
	viseted_vertices[vertex.value] = true

	fmt.Println(vertex.value)
	fmt.Println(viseted_vertices)

	for _, adjacent_vertex := range vertex.adjacent_vertices {
		if _, ok := viseted_vertices[adjacent_vertex.value]; ok {
			continue
		}
		v.dfs_traverse(adjacent_vertex, viseted_vertices)
	}
}

func (v *Vertex) dfs(vertex *Vertex, search_value string, viseted_vertices map[string]bool) *Vertex {
	if vertex.value == search_value {
		return vertex
	}

	viseted_vertices[vertex.value] = true

	for _, adjacent_vertex := range vertex.adjacent_vertices {
		// skip to next vertex if adjacent has already been visited
		if _, ok := viseted_vertices[adjacent_vertex.value]; ok {
			continue
		}

		// return adjacent if its value is the one we are looking for
		if adjacent_vertex.value == search_value {
			return adjacent_vertex
		}

		// trie to find it using recursion
		vertex_were_searching_for := v.dfs(adjacent_vertex, search_value, viseted_vertices)

		if vertex_were_searching_for != nil {
			return vertex_were_searching_for
		}
	}
	return nil
}

func (v *Vertex) bfs_traverse(starting_vertex *Vertex) {
	queue := Queue{make([]*Vertex, 0)}

	visited_vertices := make(map[string]bool)
	visited_vertices[starting_vertex.value] = true
	queue.enqueue(starting_vertex)

	for {
		if len(queue.data) <= 0 {
			break
		}
		current_vertex := queue.dequeue()

		fmt.Println(current_vertex)

		for _, adjacent_vertex := range current_vertex.adjacent_vertices {
			if _, ok := visited_vertices[adjacent_vertex.value]; ok {
				continue
			}
			visited_vertices[adjacent_vertex.value] = true
			queue.enqueue(adjacent_vertex)
		}
	}
}

func (v *Vertex) bfs(starting_vertex *Vertex, search_value string) *Vertex {
	if starting_vertex.value == search_value {
		return starting_vertex
	}

	queue := Queue{make([]*Vertex, 0)}

	visited_vertices := make(map[string]bool)
	visited_vertices[starting_vertex.value] = true
	queue.enqueue(starting_vertex)

	for {
		if len(queue.data) <= 0 {
			break
		}
		current_vertex := queue.dequeue()

		if current_vertex.value == search_value {
			fmt.Println("Search Result:", current_vertex)
			return current_vertex
		}

		for _, adjacent_vertex := range current_vertex.adjacent_vertices {
			if _, ok := visited_vertices[adjacent_vertex.value]; ok {
				continue
			}
			visited_vertices[adjacent_vertex.value] = true
			queue.enqueue(adjacent_vertex)
		}
	}
	return nil
}

func main() {
	fmt.Println("Welcome to Graphs!")

	will := Vertex{"Will", make([]*Vertex, 0)}
	theon := Vertex{"Theon", make([]*Vertex, 0)}
	morgar := Vertex{"Morgar", make([]*Vertex, 0)}

	fmt.Println("add_adjacent_vertex method:")
	will.add_adjacent_vertex(&theon)
	will.add_adjacent_vertex(&theon)
	will.add_adjacent_vertex(&morgar)
	fmt.Println("WillAdjacent:")
	will.print_all_adjacent()
	fmt.Println("")
	fmt.Println("TheonAdjacent:")
	theon.print_all_adjacent()
	fmt.Println("")
	fmt.Println("MorgarAdjacent:")
	morgar.print_all_adjacent()
	fmt.Println("")

	fmt.Println("dfs_traverse method:")
	will.dfs_traverse(&will, make(map[string]bool))
	fmt.Println("")

	fmt.Println("dfs search method:")
	searchNode := will.dfs(&will, "Morgar", make(map[string]bool))
	fmt.Println(searchNode)
	fmt.Println("")

	fmt.Println("bfs_traverse method:")
	will.bfs_traverse(&will)
	fmt.Println("")

	fmt.Println("bfs search method:")
	will.bfs(&will, "Morgar")
}
