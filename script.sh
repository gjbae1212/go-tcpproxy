#!/bin/bash
set -e -o pipefail
trap '[ "$?" -eq 0 ] || echo "Error Line:<$LINENO> Error Function:<${FUNCNAME}>"' EXIT
cd `dirname $0`
CURRENT=`pwd`

function linux_build
{
   GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o tcpproxy
}

function build
{
   go build -ldflags "-s -w" -o tcpproxy
}

CMD=$1
shift
$CMD $*
