#! /bin/bash

set -e # exit if any command returns non-zero exit code

ds=$(ls -d */ | tr -d '/')
for d in ${ds[@]}
do
	if [[ $d =~ ^-?[0-9]+$ ]]; then # is the directory name an integer
		[ -d "$d/js" ] && ls $d/js | grep index.js | 2> /dev/null
		if [[ $? -eq 0 ]]; then
			node $d/js/index.js
		fi
	fi
done