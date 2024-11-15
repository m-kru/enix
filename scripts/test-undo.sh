#!/bin/bash

update=false

help_msg="Script for managing undo command regression tests.
Must be run from the project's root.

Usage:
  scripts/test-cmd.sh <command>

Commands:
  help    Display help message.
  run     Run tests.
  update  Run tests discarding errors and update stderr.golden files using stderr files.

If no command is provided the run is assumed.
"

while true ; do
	case "$1" in
		help) printf "$help_msg" ; exit 0 ;;
		run) shift ;;
		"") shift ; break ;;
		*) echo "invalid argument '$1'" ; exit 1 ;;
	esac
done

set -e

cd tests/undo/

echo -e "\nRunning undo regression tests\n"

for dir in $(find . -maxdepth 2 -mindepth 2 -type d | sort);
do
	testname=`basename $dir`
	# Ignore tests starting with '_' character.
	if [ ${testname::1} = "_" ]; then
		continue
	fi

	echo "  $dir"
	cd "$dir"
	../../../../enix -config ../../../config.json -script script file || true
	diff --color got want
	rm got
	cd ../..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
