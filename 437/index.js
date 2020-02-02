// test cases
let tests = [
	{
		args: {
			s: "figehaeci",
			set: {
				"a": true,
				"e": true,
				"i": true
			}
		},
		expect: "aeci"
	},
	{
		args: {
			s: "abc",
			set: {
				"l": true
			}
		},
		expect: null
	}
]

// shortestSubstring returns the shortest substring in s containing all characters in set
// it returns null if no such substring exists
// naive approach, worstcase runtime: n * n - 1 -> n^2 - n -> O(n^2)
function shortestSubstring(s = "", set = {}) {
	if (s.length === 0) {
		return null
	}

	let exists = false
	let shortest = s

	for (let i = 0; i < s.length; i++) {
		if (set[s[i]] === true) {
			let cur = ""
			let seen = {}
			let found = false

			for (let j = i; j < s.length; j++) {
				cur += s[j]
				if (set[s[j]] === true) {
					seen[s[j]] = true
				}
				if (Object.keys(seen).length === Object.keys(set).length) { // seen all required for substring match
					found = true
					break
				}
			}
			
			if (found && cur.length < shortest.length) {
				exists = true
				shortest = cur
			}
		}
	}

	if (exists) {
		return shortest
	}

	return null
}

// run test cases against shortestSubstring function
function run() {
	for (let test of tests) {
		let result = shortestSubstring(test.args.s, test.args.set)

		let resultMsg = "TEST: "
		if (result !== test.expect) {
			resultMsg += "FAIL"
		} else {
			resultMsg += "PASS"
		}

		console.log(resultMsg)
		console.log("test: ")
		console.log(test)
		console.log("result: " + result)
	}
}

run()