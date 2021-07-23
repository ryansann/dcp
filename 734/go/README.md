# Time Based Map

Write a map implementation with a get function that lets you retrieve the value of a key at a particular time.

It should contain the following methods:

set(key, value, time): sets key to value for t = time.
get(key, time): gets the key at t = time.
The map should work like this. If we set a key at a particular time, it will maintain that value forever or until it gets set at a later time. In other words, when we get a key at a time, it should return the value that was set for that key set at the most recent time.

set("a", "val", 1)  
set("a", "val2", 3)  
get("a", 2) => "val"  
get("a", 3+) => "val2"  