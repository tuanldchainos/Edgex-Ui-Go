#!/bin/bash
#
# Copyright (c) 2018
# Tencent
#
# SPDX-License-Identifier: Apache-2.0
#

DIR=$PWD
CMD=../

# Kill all edgex-ui-go* stuff
function cleanup {
	pkill Edgex-Ui-Go
}

cd $CMD
exec -a Edgex-Ui-Go ./Edgex-Ui-Go &
cd $DIR

trap cleanup EXIT

while : ; do sleep 1 ; done
