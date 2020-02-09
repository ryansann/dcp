package test

type graph [][]int

type closure [][]bool

// transitiveClosure returns the transitive closure for the passed in graph g,
// where g is represented as an adjacency list.
func transitiveClosure(g graph) closure {
	if g == nil {
		return nil
	}

	var c closure
	for i := range g {
		c = append(c, dfs(g, i, make([]bool, len(g))))
	}

	return c
}

// dfs performs a depth first search on g from vertex.
// it returns a slice that indicates which vertexes are reachable from vertex.
func dfs(g graph, vertex int, visited []bool) []bool {
	visited[vertex] = true

	for _, v := range g[vertex] {
		if !visited[v] {
			visited = dfs(g, v, visited)
		}
	}

	return visited
}
