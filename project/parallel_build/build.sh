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
if [ $# == 0 ] 
then
    echo "useage:$0 go_file"
else
    mkdir ./bin
    cd ./bin
    go build "../$1"
fi
