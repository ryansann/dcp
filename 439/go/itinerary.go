package test

type flight [2]string // origin, dest

// nodes are airports an edge represents a flight between them
type flightGraph map[string][]string

// findItinerary takes a list of flights and a starting point and returns the lexicographically smallest itinerary
func findItinerary(flights []flight, start string) []string {
	g := createFlightGraph(flights)

	// find path from start
	q := []string{start}
	var itinerary []string
	for len(q) > 0 {
		// grab current node, remove from q, visit it, and add it to the itinerary
		cur := q[0]
		q = q[1:]
		itinerary = append(itinerary, cur)

		// add lexicographically lowest to q
		if len(g[cur]) != 0 {
			next := g[cur][0]
			g[cur] = g[cur][1:]
			q = append(q, next)
		}
	}

	// we didn't visit all the airports
	if len(itinerary) < len(flights)+1 { // if you had n flights, you visited n+1 airports
		return nil
	}

	return itinerary
}

// createFlightGraph returns the flights represented as an adjacency list
func createFlightGraph(flights []flight) flightGraph {
	g := make(flightGraph, 0)
	for _, f := range flights {
		origin, dest := f[0], f[1]

		// origin doesn't in g
		if _, ok := g[origin]; !ok {
			g[origin] = []string{dest}
			continue
		}

		// g needs a new node
		add := dest
		for i, n := range g[origin] { // add in lexicographical order
			if dest < n {
				g[origin][i], add = dest, n
			}
		}

		g[origin] = append(g[origin], add)
	}

	return g
}
