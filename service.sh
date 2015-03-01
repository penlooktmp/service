#!/bin/bash
# Copyright 2014 Penlook Development Team. All rights reserved.
# Use of this source code is governed by
# license that can be found in the LICENSE file.
# Author : Loi Nguyen <loint@penlook.com>

# Build daemon
function daemon {
	cd component/daemon
	chmod +x ./build.sh && ./build.sh
	cd ../../
}

# Get service name
function service_name {
	echo $1 | rev | cut -d"/" -f1 | rev
}

# Build service
function build {
	
	cd $1
	go build
	sudo ./$SERVICE remove
	sudo ./$SERVICE install

	LOGFILE="/var/log/$1.log"
	ERRFILE="/var/log/$1.err"

	[[ -e $LOGFILE ]] || sudo touch $LOGFILE
	[[ -e $ERRFILE ]] || sudo touch $ERRFILE

	sudo chmod a+w $LOGFILE
	sudo chmod a+w $ERRFILE

	echo "" > $LOGFILE
	echo "" > $ERRFILE

	cd ../
}

# Watch service
function watch {
	sudo netstat -tulpn
}

# Cleanup service
function clean {

	stop $1

	LOGFILE="/var/log/$1.log"
	ERRFILE="/var/log/$1.err"
	PIDFILE="/var/run/$1.pid"

	sudo rm -rf $LOGFILE
	sudo rm -rf $ERRFILE
	sudo rm -rf $PIDFILE
}

# Test service
function test {
	LOGFILE="/var/log/$1.log"
	ERRFILE="/var/log/$1.err"

	status $1
	cd $1
	go test -v
	cat $LOGFILE
	cat $ERRFILE
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

function client {
	SERVICE=$(service_name $1)
	CLIENT="./client/$SERVICE/main"
	sudo chmod +x $CLIENT*

	for cmd in $@
	do
		if [ $cmd != $SERVICE ] &&  [ $cmd != $2 ]
		then
			[ -e $CLIENT.$cmd ] && eval $CLIENT.$cmd
		fi
	done
}

# Main function
function main {

	daemon
	cd service

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

		if [ $1 == "watch" ] || [ $1 == "clean" ]
		then
			$1
			exit
		fi

		if [ $2 == "client" ]
		then
			cd ..
			$2 $@
			exit
		fi

		# Enter debug mode
		debug $@
	fi
}

main $@