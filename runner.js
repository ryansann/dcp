// terminal color constants
var fgBlue = "\x1b[1;34m"
var fgRed = "\x1b[1;31m"
var fgGreen = "\x1b[32m"
var fgReturn = "\x1b[0m"

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

	run(f, tests = []) {
		console.log(fgBlue + this.name + fgReturn)
		console.group()
		
		for (let i in tests) {
			let test = tests[i]

			let result = f(...test.args)

			let resultMsg = "TEST " + i + ": "
			if (!deepEqual(result, test.expect)) {
				resultMsg += (fgRed + "FAIL" + fgReturn)
			} else {
				resultMsg += (fgGreen + "PASS" + fgReturn)
			}

			console.log(resultMsg)
			console.group()
			console.log("test: ")
			console.group()
			console.log(test)
			console.groupEnd()
			console.log("result: ")
			console.group()
			console.log(result)
			console.groupEnd()
			console.groupEnd()
		}

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
		if (!deepEqual(obj1[key], obj2[key])) return false;
	}

	return true;
}

// check if value is primitive
function isPrimitive(obj) {
	return (obj !== Object(obj));
}

module.exports = TestRunner