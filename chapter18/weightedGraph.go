package main

import "fmt"

type WeightedGraphVertex struct {
	value             string
	adjacent_vertices map[*WeightedGraphVertex]int
}

func (wgv *WeightedGraphVertex) add_adjacent_vertex(vertex *WeightedGraphVertex, weight int) {
	wgv.adjacent_vertices[vertex] = weight
}

func newWeightedGraphVertex(name string) WeightedGraphVertex {
	return WeightedGraphVertex{name, make(map[*WeightedGraphVertex]int)}
}

func filterSlice(unvisited_cities []*WeightedGraphVertex, current_city *WeightedGraphVertex) []*WeightedGraphVertex {
	filtered_cities := make([]*WeightedGraphVertex, 0)

	for _, city := range unvisited_cities {
		if city.value != current_city.value {
			filtered_cities = append(filtered_cities, city)
		}
	}
	return filtered_cities
}

func cheapestUnvisited(cities []*WeightedGraphVertex, cheapest_prices_table map[string]int) *WeightedGraphVertex {
	if len(cities) <= 0 {
		return nil
	}
	cheapestCity := cities[0]

	for _, city := range cities {
		currentLow := cheapest_prices_table[cheapestCity.value]
		cityCost := cheapest_prices_table[city.value]

		if cityCost < currentLow {
			cheapestCity = city
		}
	}
	return cheapestCity
}

func dijkstra_shortest_path(starting_city *WeightedGraphVertex, final_destination *WeightedGraphVertex) []string {
	cheapest_prices_table := make(map[string]int)
	cheapest_previous_stopover_city_table := make(map[string]string)

	unvisited_cities := make([]*WeightedGraphVertex, 0)

	visited_cities := make(map[string]bool)

	cheapest_prices_table[starting_city.value] = 0

	current_city := starting_city

	for {
		if current_city == nil {
			break
		}

		visited_cities[current_city.value] = true
		unvisited_cities = filterSlice(unvisited_cities, current_city)

		//fmt.Println(current_city.adjacent_vertices)
		for adjacent_city, price := range current_city.adjacent_vertices {
			// add to list of unvisited cities if it is not in the list of visited
			if _, ok := visited_cities[adjacent_city.value]; !ok {
				unvisited_cities = append(unvisited_cities, adjacent_city)
			}

			price_through_current_city := cheapest_prices_table[current_city.value] + price

			// if price from starting city to adjacent city is cheapest so far
			if _, ok := cheapest_prices_table[adjacent_city.value]; !ok || price_through_current_city < cheapest_prices_table[adjacent_city.value] {

				//update two tables
				cheapest_prices_table[adjacent_city.value] = price_through_current_city
				cheapest_previous_stopover_city_table[adjacent_city.value] = current_city.value
			}

		}
		//vist next unvistied city, choosing the cheapest to get to from STARTING CITY
		current_city = cheapestUnvisited(unvisited_cities, cheapest_prices_table)
	}

	fmt.Println("Cheapest Price Table:", cheapest_prices_table)
	fmt.Println("Cheapest Stopover Table:", cheapest_previous_stopover_city_table)

	// create route path

	shortest_path := make([]string, 0)

	current_city_name := final_destination.value

	for {
		if current_city_name == starting_city.value {
			break
		}
		shortest_path = append(shortest_path, current_city_name)
		current_city_name = cheapest_previous_stopover_city_table[current_city_name]
	}
	shortest_path = append(shortest_path, starting_city.value)

	return shortest_path
}

func main() {

	denver := newWeightedGraphVertex("Denver")
	elpaso := newWeightedGraphVertex("El Paso")
	chicago := newWeightedGraphVertex("Chicago")
	atlanta := newWeightedGraphVertex("Atlanta")
	boston := newWeightedGraphVertex("Boston")

	// add adjacent
	//fmt.Println("add_adjacent_vertex method:")
	atlanta.add_adjacent_vertex(&denver, 160)
	atlanta.add_adjacent_vertex(&boston, 100)
	boston.add_adjacent_vertex(&denver, 180)
	boston.add_adjacent_vertex(&chicago, 120)
	chicago.add_adjacent_vertex(&elpaso, 80)
	denver.add_adjacent_vertex(&chicago, 40)
	denver.add_adjacent_vertex(&elpaso, 140)
	elpaso.add_adjacent_vertex(&boston, 100)

	// fmt.Println(denver)
	// fmt.Println(elpaso)
	// fmt.Println(chicago)
	// fmt.Println(atlanta)
	// fmt.Println(boston)
	// fmt.Println("")

	fmt.Println(dijkstra_shortest_path(&atlanta, &elpaso))

	// exercise 5
	idris := newWeightedGraphVertex("Idris")
	kamil := newWeightedGraphVertex("Kamil")
	lina := newWeightedGraphVertex("Lina")
	sasha := newWeightedGraphVertex("Sasha")
	marco := newWeightedGraphVertex("Marco")
	ken := newWeightedGraphVertex("Ken")
	talia := newWeightedGraphVertex("Talia")
	idris.add_adjacent_vertex(&kamil, 1)
	idris.add_adjacent_vertex(&talia, 1)
	kamil.add_adjacent_vertex(&lina, 1)
	lina.add_adjacent_vertex(&sasha, 1)
	sasha.add_adjacent_vertex(&marco, 1)
	ken.add_adjacent_vertex(&ken, 1)
	talia.add_adjacent_vertex(&talia, 1)

	fmt.Println(dijkstra_shortest_path(&idris, &lina))

}
