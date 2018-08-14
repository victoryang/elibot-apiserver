#!/bin/sh
#
# Start api-server
#

case "$1" in
	start)
		echo "starting api-server..."
		nohub ./rbctrl/api-server > /dev/null &

	stop)
		echo -n "stopping api-server"
		killall api-server
esac