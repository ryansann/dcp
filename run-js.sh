#! /bin/bash

set -e # exit if any command returns non-zero exit code

ds=$(ls -d */ | tr -d '/')
for d in ${ds[@]}
do
	ls $d | grep index.js | 2> /dev/null
	if [[ $? -eq 0 ]]; then
		node $d/index.js
	fi
done