// test cases
let tests = [
	{
		graph: [ [ ] ],
		expect: [ [ 1 ] ]
	},
	{
		graph: [
			[ 0, 1, 3 ],
			[ 1, 2 ],
			[ 2 ],
			[ 3 ]
		],
		expect: [
			[ 1, 1, 1, 1 ],
			[ 0, 1, 1, 0 ],
			[ 0, 0, 1, 0 ],
			[ 0, 0, 0, 1 ]
		]
	}
]

// Graph class containing methods for finding transitive closure
class Graph {
	// constructor accepts an adjacency list representation of a graph
	constructor(graph = [[]]) {
		this.g = graph
		this.v = graph.length
	}

	// transitiveClosure is O(v^2 + e) since dfs is called for each vertex
	transitiveClosure() {
		let closure = []
		
		// find visitable nodes for each vertex in g
		// add these nodes to the closure
		for (let v in this.g) {
			closure = [ ...closure, this.dfs(v, new Array(this.v).fill(0)) ]
		}

		return closure
	}

	// dfs is O(v + e) where v is the number of verticies and e is the number of edges
	dfs(v, visited = []) {
		visited[v] = 1 // mark current vertex as visited

		for (let e of this.g[v]) {
			if (visited[e] !== 1) {
				visited = this.dfs(e, visited)
			}
		}

		return visited
	}
}

// run executes transitiveClosure against test cases
function run() {
	for (let test of tests) {
		g = new Graph(test.graph)
		c = g.transitiveClosure()

		let didFail = false
		let resultMsg = "TEST: "
		for (i in test.expect) {
			for (j in test.expect[i]) {
				let actual = c[i][j]
				let expected = test.expect[i][j]
				if (actual !== expected) {
					resultMsg += "FAIL - result[" + i + "][" + j + "] -> " + actual + " != expect[" + i + "][" + j + "] -> " + expected 
					didFail = true
					break
				}
			}
		
			if (didFail) {
				break
			}
		}

		if (!didFail) {
			resultMsg += "PASS"
		}

		console.log(resultMsg)
		console.log("test:")
		console.log(test)
		console.log("result:")
		console.log(c)

		if (didFail) {
			require("process").exit(1)
		}
	}
}

run()