# ============= Golang ============= #

go--genproto:
	protoc --go_out=plugins=grpc:src/go/common/pb -I proto/ proto/*.proto

go--deps:
	go mod tidy
	go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

# ============= Store service ============= #

store-service--build:
	make build PROJECT=store-service

store-service--run:
	make run PROJECT=store-service

store-service--docker-build:
	make docker-build PROJECT=store-service

store-service--docker-push:
	make docker-push PROJECT=store-service

# ============= User service ============= #

user-service--build:
	make build PROJECT=user-service

user-service--run:
	make run PROJECT=user-service

user-service--docker-build:
	make docker-build PROJECT=user-service

user-service--docker-push:
	make docker-push PROJECT=user-service

# ============= Gateway service ============= #

gateway-service--build:
	make build PROJECT=gateway-service

gateway-service--run:
	make run PROJECT=gateway-service

gateway-service--docker-build:
	make docker-build PROJECT=gateway-service

gateway-service--docker-push:
	make docker-push PROJECT=gateway-service

# ============= Product service ============= #

product-service--build:
	make build PROJECT=product-service

product-service--run:
	make run PROJECT=product-service

product-service--docker-build:
	make docker-build PROJECT=product-service

product-service--docker-push:
	make docker-push PROJECT=product-service

# ============= Product service ============= #

store-product-service--build:
	make build PROJECT=store-product-service

store-product-service--run:
	make run PROJECT=store-product-service

store-product-service--docker-build:
	make docker-build PROJECT=store-product-service

store-product-service--docker-push:
	make docker-push PROJECT=store-product-service

# ============= Sesssion service ============= #

session-service--build:
	make build PROJECT=session-service

session-service--run:
	make run PROJECT=session-service

session-service--docker-build:
	make docker-build PROJECT=session-service

session-service--docker-push:
	make docker-push PROJECT=session-service

# ============= Generic commands ============= #

docker-build:
	bazel run //:$(PROJECT) -- --norun

docker-push:
	bazel run //:$(PROJECT)-push

build:
	bazel build //src/go/$(PROJECT):$(PROJECT)

run:
	bazel run //src/go/$(PROJECT):$(PROJECT)
