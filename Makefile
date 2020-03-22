go--genproto:
	protoc --go_out=plugins=grpc:src/go/common/pb -I proto/ proto/*.proto

go--deps:
	go mod tidy
	go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

store-service--build:
	bazel run //:store-service -- --norun

store-service--push:
	bazel run //:store-service-push

user-service--build:
	bazel run //:user-service -- --norun

user-service--push:
	bazel run //:user-service-push

gateway-service--build:
	bazel run //:gateway-service -- --norun

gateway-service--push:
	bazel run //:gateway-service-push

product-service--build:
	bazel run //:product-service -- --norun

product-service--push:
	bazel run //:product-service-push

session-service--build:
	bazel run //:session-service -- --norun

session-service--push:
	bazel run //:session-service-push