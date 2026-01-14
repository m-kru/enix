#!/bin/bash

update=false

help_msg="Script for managing command line arguments parsing tests.
Must be run from the project's root.

Usage:
  scripts/test-arg.sh <command>

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
    update) update=true ; shift ;;
    "") shift ; break ;;
    *) echo "invalid argument '$1'" ; exit 1 ;;
  esac
done

if ! $update; then
  set -e
fi

cd tests/arg/

echo -e "\nRunning command line argument parsing tests\n"

for dir in $(find . -maxdepth 1 -mindepth 1 -type d | sort);
do
  testname=`basename $dir`
  # Ignore tests starting with '_' character.
  if [ ${testname::1} = "_" ]; then
    continue
  fi

  echo "  $dir"
  cd "$dir"
  ./run.sh || true
  diff --color got want
  if $update; then
    cp got want
  fi
  rm got
  cd ..
done

if $update; then
  echo -e "\nwant files updated\n"
else
  echo -e "\nAll \e[1;32mPASSED\e[0m!"
fi
