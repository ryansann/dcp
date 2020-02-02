// test cases
let tests = [ 
	{
		args: [
			undefined,
			5
		],
		expect: {
			floor: null, 
			ceiling: null
		}
	},
	{
		args: [
			{
				value: 5,
				left: {
					value: 3
				},
				right: {
					value: 7
				}
			},
			2
		],
		expect: {
			floor: null, 
			ceiling: 7
		}
	},
	{
		args: [
			{
				value: 5,
				left: {
					value: 3
				},
				right: {
					value: 7
				}
			}, 
			8
		],
		expect: {
			floor: 3, 
			ceiling: null
		}
	},
	{
		args: [
			{
				value: 5,
				left: {
					value: 3
				},
				right: {
					value: 7
				}
			},
			5
		],
		expect: {
			floor: 3, 
			ceiling: 7
		}
	},
	{
		args: [
			{
				value: 5,
				left: {
					value: 4
				},
				right: {
					value: 7
				}
			},
			4
		],
		expect: {
			floor: 4, 
			ceiling: 7
		}
	},
	{
		args: [
			{
				value: 5,
				left: {
					value: 4
				},
				right: {
					value: 7
				}
			},
			7
		],
		expect: {
			floor: 4, 
			ceiling: 7
		}
	},
	{
		args: [
			{
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
			7
		],
		expect: {
			floor: 1, 
			ceiling: 9
		}
	},
]

// find returns an object with the floor and ceiling for target in tree.
function find(tree = { }, target = 0) {
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

const Runner = new require("../runner")
let r = new Runner("BST Floor & Ceiling")
r.run(find, tests)