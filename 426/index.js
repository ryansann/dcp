// test cases
let tests = [
	{
		tree: {
			value: 1,
			left: {
				value: 3
			},
			right: {
				value: -3
			}
		},
		expect: 1
	},
	{
		tree: {
			value: 1,
			left: {
				value: 3,
				left: {
					value: 1,
					left: {
						value: 2
					},
					right: {
						value: 2
					}
				},
				right: {
					value: 2,
					left: {
						value: 2
					},
					right: {
						value: 2
					}
				}
			},
			right: {
				value: -3,
				left: {
					value: -3,
					left: {
						value: -4
					},
					right: {
						value: -4
					}
				},
				right: {
					value: -1,
					left: {
						value: -2
					},
					right: {
						value: -2
					}
				}
			}
		},
		expect: 3
	}
]

function minSumLevel(tree = {}) {
	let level = [ tree ]
	let curLevel = 0
	let minLevel = -1
	let minSum = Number.MAX_SAFE_INTEGER
	
	while (true) {
		let curSum = 0
		let nextLevel = []
		
		for (let node of level) {
			curSum += node.value
			if (node.left !== undefined) {
				nextLevel = [ node.left, ...nextLevel ]
			}
			if (node.right !== undefined) {
				nextLevel = [ node.right, ...nextLevel ]
			}
		}

		if (curSum < minSum) {
			minSum = curSum
			minLevel = curLevel
		}

		if (nextLevel.length === 0) {
			break
		}

		curLevel++
		level = nextLevel
	}

	return minLevel
}

function run() {
	for (test of tests) {
		let result = minSumLevel(test.tree)
		
		var resultMsg = "TEST: "
		if (result === test.expect) {
			resultMsg += "PASS"
		} else {
			resultMsg += "FAIL"
		}

		console.log(resultMsg)
		console.log("test: " + JSON.stringify(test, null, "  "))
		console.log("result: " + result)
	}
}

run()