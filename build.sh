#!/bin/bash
# Copyright 2014 Penlook Development Team. All rights reserved.
# Use of this source code is governed by
# license that can be found in the LICENSE file.
# Author : Loi Nguyen <loint@penlook.com>

# Build all services
function build {
	SERVICE=$(echo $1 | rev | cut -d"/" -f1 | rev)
	cd $SERVICE
	go build service.go
	#service $SERVICE stop
	#./service remove
	#./service install
}

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
fi