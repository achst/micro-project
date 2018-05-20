#!/bin/bash

# modify here
name="service_order"

################################################################################

path="./py/$name"

export GOPATH=/Users/tanshuai/lab/go

cd $GOPATH/src/github.com/hopehook/micro-project/proto

mkdir -p $path

touch ./__init__.py
touch ./py/__init__.py
touch $path/__init__.py

python -m grpc_tools.protoc -I$GOPATH/src:. --python_out=$path --grpc_python_out=$path ./$name.proto