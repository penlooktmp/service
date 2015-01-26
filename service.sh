#!/bin/bash
# Copyright 2014 Penlook Development Team. All rights reserved.
# Use of this source code is governed by
# license that can be found in the LICENSE file.
# Author : Loi Nguyen <loint@penlook.com>

# Build daemon
function daemon {
	cd modules/daemon
	chmod +x ./build.sh && ./build.sh
	cd ../../
}

# Get service name
function service_name {
	echo $1 | rev | cut -d"/" -f1 | rev
}

# Build services
function build {
	cd $1
	go build
	sudo ./$SERVICE remove
	sudo ./$SERVICE install
	cd ../
}

# Watch services
function watch {
	sudo netstat -tulpn
}

# Test services
function test {
	cd $1
	go test
	cd ..
}

# Start service
function start {
	sudo service $1 start
}

# Stop service
function stop {
	sudo service $1 stop
}

# Satus service
function status {
	sudo service $1 status
}

# Restart service
function restart {
	stop $1
	start $1
}

# Debug mode
function debug {
	SERVICE=$(service_name $1)

	for cmd in $@
	do
		if [ $cmd != $SERVICE ]
		then
			eval $cmd $SERVICE
		fi
	done
}

# Main function
function main {
	daemon
	cd services

	if [ -z "$1" ]
	then
		# Build all services
		for D in `find . -type d`
		do
			if [ $D != "." ]
			then
				build $D
			fi
		done
	else

		if [ -z "$2" ]
		then
			SERVICE=$(service $1)
			test $SERVICE
			build $SERVICE
			exit
		fi

		if [ $1 == "watch" ]
		then
			watch
			exit
		fi

		# Enter debug mode
		debug $@
	fi
}

main $@