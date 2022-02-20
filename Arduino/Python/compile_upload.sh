#!/bin/bash

if ! arduino-cli compile --fqbn arduino:avr:micro Upload ; then
    echo -e "\n[ Compiling failed ]"
    exit 1
fi

if ! arduino-cli upload -p /dev/ttyACM0 --fqbn arduino:avr:micro Upload ; then
    echo -e "\n[ Uploading failed ]"
    exit 1
fi

echo -e "\nCompiling and uploading OK"
exit 0

