#!/bin/sh

update=false

help_msg="Script for managing command regression tests.
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
    help) printf '%s\n' "$help_msg" ; exit 0 ;;
    run) shift ;;
    "") break ;;
    *) printf "invalid argument '%s'\n" "$1"; exit 1 ;;
  esac
done

set -e

cd tests/cmd/

printf "\nRunning command regression tests\n"

for dir in $(find . -maxdepth 2 -mindepth 2 -type d | sort);
do
  testname=$(basename "$dir")
  # Ignore tests starting with '_' character.
  if [ "${testname#_}" != "$testname" ]; then
    continue
  fi

  printf "  %s\n" "$dir"
  cd "$dir"
  ../../../../enix -config ../../../config.json -script script file || true
  diff --color got want
  rm got
  cd ../..
done

printf "\n\033[32mAll PASSED\033[0m\n";