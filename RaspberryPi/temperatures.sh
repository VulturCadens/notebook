#!/bin/bash

cpu=$(</sys/class/thermal/thermal_zone0/temp)
gpu=$(vcgencmd measure_temp | awk -F"=|'" '{print $2}')

echo "GPU: $gpu °C"
echo "CPU: $((cpu/1000)) °C"
