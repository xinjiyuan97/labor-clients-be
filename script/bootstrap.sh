#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
env=$1
BinaryName=hertz_service
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName} server $env