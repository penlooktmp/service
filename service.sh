#!/bin/bash
#
# Penlook Project
#
# Copyright (c) 2015 Penlook Development Team
#
# --------------------------------------------------------------------
#
# This program is free software: you can redistribute it and/or
# modify it under the terms of the GNU Affero General Public License
# as published by the Free Software Foundation, either version 3
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public
# License along with this program.
# If not, see <http://www.gnu.org/licenses/>.
#
# --------------------------------------------------------------------
#
# Authors:
#     Loi Nguyen       <loint@penlook.com>
#     Tin Nguyen      <tinntt@penlook.com>

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

	# if executable file is exist, service maybe installed
	# we need to remove it before installing a new one
	[[ -f ./$SERVICE ]] && sudo ./$SERVICE remove

	go build
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
	true
}

# Stop service
function stop {
	sudo service $1 stop
	true
}

# Satus service
function status {
	sudo service $1 status
	true
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