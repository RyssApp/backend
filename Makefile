go--genproto:
	protoc --go_out=plugins=grpc:src/go/common/pb -I proto/ proto/*.proto

go--deps:
	go mod tidy
	go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

store-service--build:
	bazel run //:store-service

store-service--push:
	bazel run //:store-service-push
