#!/bin/bash

base_tput=$1
trigger_tput=$2

./build/wlgen/wlgen_proc/wlgen_proc/wlgen_proc -outfile=stats_1.csv -tput=$base_tput -duration=30s &
sleep 30
./build/wlgen/wlgen_proc/wlgen_proc/wlgen_proc -outfile=stats_2.csv -tput=$trigger_tput -duration=10s &
sleep 10
./build/wlgen/wlgen_proc/wlgen_proc/wlgen_proc -outfile=stats_3.csv -tput=$base_tput -duration=20s