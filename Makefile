go--genproto:
	protoc --go_out=plugins=grpc:src/go/common/pb -I proto/ proto/votebot.proto

go--deps:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle