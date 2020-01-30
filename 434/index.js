// test cases
let tests = [ 
	{
		target: 5,
		expect: {
			floor: null, 
			ceiling: null
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 3
			},
			right: {
				value: 7
			}
		},
		target: 2,
		expect: {
			floor: null, 
			ceiling: 7
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 3
			},
			right: {
				value: 7
			}
		},
		target: 8,
		expect: {
			floor: 3, 
			ceiling: null
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 3
			},
			right: {
				value: 7
			}
		},
		target: 5,
		expect: {
			floor: 3, 
			ceiling: 7
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 4
			},
			right: {
				value: 7
			}
		},
		target: 4,
		expect: {
			floor: 4, 
			ceiling: 7
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 4
			},
			right: {
				value: 7
			}
		},
		target: 7,
		expect: {
			floor: 4, 
			ceiling: 7
		}
	},
	{
		tree: {
			value: 5,
			left: {
				value: 3,
				left: {
					value: 1
				},
				right: {
					value: 4
				}
			},
			right: {
				value: 7,
				left: {
					value: 6
				},
				right: {
					value: 9
				}
			}
		},
		target: 7,
		expect: {
			floor: 1, 
			ceiling: 9
		}
	},
]

// find returns an object with the floor and ceiling for target in tree.
function find(target = 0, tree = { }) {
	return { floor: findFloor(target, tree), ceiling: findCeiling(target, tree) }
}

// O(n) where n is number of nodes in tree, if balanced O(log_2(n))
function findFloor(target, tree) {
	let node = tree
	while (true) {
		if (node.left !== undefined) { // traverse to minimum element in tree
			node = node.left
			continue
		}

		if (node.value <= target) {
			return node.value
		} else {
			return null
		}
	}
}

// O(n) where n is number of nodes in tree, if balanced O(log_2(n))
function findCeiling(target, tree) {
	let node = tree
	while (true) {
		if (node.right !== undefined) { // traverse to maximum element in tree
			node = node.right
			continue
		}

		if (node.value >= target) {
			return node.value
		} else {
			return null
		}
	}
}

// run runs the tests against find
function run() {
	// run function against each test
	for (let test of tests) {
		// find the floor and ceiling
		let result = find(test.target, test.tree)

		// verify the result is what is expected
		let resultMsg = "TEST: "
		if (result.ceiling === test.expect.ceiling && result.floor === test.expect.floor) {
			resultMsg += "PASS"
		} else {
			resultMsg += "FAIL"
		}

		// log the test result
		console.log(resultMsg)
		console.log("test: " + JSON.stringify(test, null, "  "))
		console.log("result: " + JSON.stringify(result, null, "  "))
	}
}

run()