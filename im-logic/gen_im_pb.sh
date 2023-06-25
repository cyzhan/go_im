#!/bin/bash

service_name="imlogic"
expect_array=("common.proto")

service_path="internal/rpc"
go_folder="$(pwd)/internal"
proto_folder="../../im-pb"
rpc_folder="${go_folder}/rpc"

echo "-----delete pb.go-----"
echo "cd $go_folder"
cd "$go_folder"
rm -rf "$rpc_folder"

echo "-----gen proto-----"
echo "cd $proto_folder"
cd "$proto_folder"

for proto_file in *.proto; do
  scriptStr="--proto_path=. --go_out=. ${proto_file} --go-grpc_out=require_unimplemented_servers=false:."

  for expect in "${expect_array[@]}";do
    if ! [[ "${proto_file}" == $expect ]];then
       scriptStr="${scriptStr} --go_opt=Mcommon.proto=${service_name}/${service_path}/common --go-grpc_opt=Mcommon.proto=${service_name}/${service_path}/common"
      fi
  done

  protoc ${scriptStr}
done

mv -fn ./internal/rpc "$rpc_folder"
