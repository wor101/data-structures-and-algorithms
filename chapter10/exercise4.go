package main

type Numbers struct {
	number int
}

type NumberArrays struct {
	array []int
}

type multiTypeSlice interface {
	Read()
}

func (n Numbers) Read() int {
	return n.number
}

func (a NumberArrays) Read() []int {
	return a.array
}

func main() {

	arrays := []multiTypeSlice{}

}
