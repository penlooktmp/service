#!/bin/bash
# Copyright 2014 Penlook Development Team. All rights reserved.
# Use of this source code is governed by
# license that can be found in the LICENSE file.
# Author : Loi Nguyen <loint@penlook.com>

# Build all services
function build {
	SERVICE=$(echo $1 | rev | cut -d"/" -f1 | rev)
	cd $SERVICE
	go test
	go build
	sudo service $SERVICE stop
	sudo rm -rf /var/run/api.pid
	sudo ./$SERVICE remove
	sudo ./$SERVICE install
}

# Watch all services
function watch {
	SERVICE=$(echo $1 | rev | cut -d"/" -f1 | rev)
	sudo service $SERVICE start
	sudo netstat -tulpn
}

# Main function
function main {
	cd services

	# Build command
	if [ -z "$1" ]
	then
		for D in `find . -type d`
		do
			if [ $D != "." ]
			then
				build $D
			fi
		done
	else
		build $1
		watch $1
	fi
}

main $1