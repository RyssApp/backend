go--genproto:
	protoc --go_out=plugins=grpc:src/go/common/pb -I proto/ proto/*.proto

go--deps:
	go mod tidy
	go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

user-service--push:
	bazel run //:user-service -- --norun
