#!/bin/bash
# Copyright 2014 Penlook Development Team. All rights reserved.
# Use of this source code is governed by
# license that can be found in the LICENSE file.
# Loi Nguyen <loint@penlook.com>
mvn package
mv ./target/service.jar ./service.jar
echo
echo
java -jar ./service.jar $@
echo