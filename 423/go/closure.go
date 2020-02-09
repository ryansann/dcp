package test

type graph [][]int

type closure [][]bool

// transitiveClosure returns the transitive closure for the passed in graph g,
// where g is represented as an adjacency list.
// runtime: O(v^2) where v is # of verticies
func transitiveClosure(g graph) closure {
	if g == nil {
		return nil
	}

	var c closure
	for i := range g { // perform dfs for each vertex
		c = append(c, dfs(g, i, make([]bool, len(g)))) // v * (v + e) = v^2 + ve = O(v^2)
	}

	return c
}

// dfs performs a depth first search on g from vertex.
// it returns a slice that indicates which verticies are reachable from vertex.
// runtime: O(v + e) where v is # of verticies and e is # of edges.
func dfs(g graph, vertex int, visited []bool) []bool {
	visited[vertex] = true

	for _, v := range g[vertex] {
		if !visited[v] {
			visited = dfs(g, v, visited)
		}
	}

	return visited
}
