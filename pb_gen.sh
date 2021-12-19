#!/bin/bash
# protoc --go_out=. *.proto
pakcage_name=${1}
idl_in_path=./pkg/idl/message
idl_out_path=./pkg/idl/pb_file

protoc --go_out=${idl_out_path}/${pakcage_name} ${idl_in_path}/${pakcage_name}/*.proto
echo "create pb files finished"
