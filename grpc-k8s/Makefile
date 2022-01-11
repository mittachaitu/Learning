.PHONY: proto-gen
proto-gen:
	protoc --go_out=./protobuffers --go-grpc_out=./protobuffers proto/command.proto

.PHONY:	clean
clean:
	rm -rf ./protobuffers/go
	rm -rf ./bin/*

.PHONY: server
server:
	@PNAME=server CTLNAME=server sh -c "'$(PWD)/build/build.sh'"

.PHONY: client
client:
	@PNAME=client CTLNAME=client sh -c "'$(PWD)/build/build.sh'"
