#!/bin/bash

set -e

if [ "$#" -gt 0 ]; then
	pid="$1"
else
	enixProcesses=$(ps u | grep "./enix" | grep -v grep)
	if [ -z "${enixProcesses}" ]; then
		echo "no enix processes runing"
		exit 1
	fi

	processesCount=$(echo "$enixProcesses" | wc -l)

	if [ "$processesCount" -gt 1 ]; then
		echo "more than single enix process running, please specify PID explicitly"
		echo "$enixProcesses"
		exit 1
	fi

	pid=$(echo "$enixProcesses" | awk '{print $2}')
fi

gdlv attach "$pid"
