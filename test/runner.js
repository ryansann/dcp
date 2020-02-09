const util = require('util')

// terminal color constants
const fgBlue = "\x1b[1;34m"
const fgRed = "\x1b[1;31m"
const fgGreen = "\x1b[32m"
const fgReturn = "\x1b[0m"

// functional test runner
// tests should be of the form: 
// 	{
// 		args: [
// 			// function arguments
// 		],
// 		expect: {} || "" || ...
// 	}
class TestRunner {
	constructor(name) {
		this.name = ""
	
		if (typeof name === "string") {
			this.name += (name + " ")
		}
		
		this.name += "Tests"
	}

	// run executes f against each test case and pretty prints the results
	run(f, tests = []) {
		// test output
		console.log(fgBlue + this.name + fgReturn)
		console.group()
		
		tests.forEach((test, i) => {
			let result = f(...test.args)

			let resultMsg = "TEST " + (i + 1) + ": "
			if (!deepEqual(result, test.expect)) {
				resultMsg += (fgRed + "FAIL" + fgReturn)
			} else {
				resultMsg += (fgGreen + "PASS" + fgReturn)
			}

			// test case output
			console.log(resultMsg)
			console.group()
			console.log("test: ")
			console.group()
			console.log(util.inspect(test, {showHidden: false, depth: null}))
			console.groupEnd()
			console.log("result: ")
			console.group()
			console.log(util.inspect(result, {showHidden: false, depth: null}))
			console.groupEnd()
			console.groupEnd()
		})

		console.groupEnd()
	}
}


// check if two objects are deeply equal
function deepEqual(obj1, obj2) {
	if (obj1 === obj2) { // it's just the same object, no need to compare
		return true
	}

	if (isPrimitive(obj1) && isPrimitive(obj2)) { // compare primitives
		return obj1 === obj2
	}

	if (Object.keys(obj1).length !== Object.keys(obj2).length) {
		return false
	}

	// compare objects with same number of keys
	for(let key in obj1) {
		if (!(key in obj2)) return false; // other object doesn't have this prop
		if (!deepEqual(obj1[key], obj2[key])) return false
	}

	return true
}

// check if value is primitive
function isPrimitive(obj) {
	return (obj !== Object(obj))
}

module.exports = TestRunner