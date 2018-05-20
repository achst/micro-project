#!/bin/bash

# modify here
name="service_subscriber"

################################################################################

path="./go/$name"

export GOPATH=/Users/tanshuai/lab/go

cd $GOPATH/src/github.com/hopehook/micro-project/proto

mkdir -p $path

protoc --proto_path=$GOPATH/src:. --micro_out=$path --go_out=$path ./$name.proto
