#!/bin/bash

command=$1

if [ -z $command ]; then
    echo "There must be a command"
    exit 0
fi

echo " start $command:" $(date -u)
$command
sleep_time=$(((RANDOM % 4)+4)) # 4-7 seconds
echo " sleeping for $sleep_time seconds..."
sleep $sleep_time # simulate work
echo " end $command:" $(date -u)