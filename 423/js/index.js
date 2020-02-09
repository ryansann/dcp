// test cases
let tests = [
	{
		args: [
			[ [ ] ]
		],
		expect: [ [ 1 ] ]
	},
	{
		args: [
			[
				[ 0, 1, 3 ],
				[ 1, 2 ],
				[ 2 ],
				[ 3 ]
			]
		],
		expect: [
			[ 1, 1, 1, 1 ],
			[ 0, 1, 1, 0 ],
			[ 0, 0, 1, 0 ],
			[ 0, 0, 0, 1 ]
		]
	}
]

// transitiveClosure is O(v^2 + e) since dfs is called for each vertex
function transitiveClosure(g = [[]]) {
	let closure = []
	
	// find visitable nodes for each vertex in g
	// add these nodes to the closure
	for (let v in g) {
		closure = [ ...closure, dfs(g, v, new Array(g.length).fill(0)) ]
	}

	return closure
}

// dfs is O(v + e) where v is the number of verticies and e is the number of edges
function dfs(g, v, visited = []) {
	visited[v] = 1 // mark current vertex as visited

	for (let e of g[v]) {
		if (visited[e] !== 1) {
			visited = dfs(g, e, visited)
		}
	}

	return visited
}

const Runner = new require("../../test/runner")
let r = new Runner("Graph Transitive Closure")
r.run(transitiveClosure, tests)