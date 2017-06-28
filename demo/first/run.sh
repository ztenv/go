#!/bin/bash

#/**
#* @file run.sh
#* @brief run go file
#* @author shlian
#* @version 1.0
#* @date 2017-06-28
#*/


CURDIR=`pwd`
export GOPATH="$GOPATH:$CURDIR"

go run $1
