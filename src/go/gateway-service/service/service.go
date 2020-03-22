package service

import (
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/gateway-service/config"
	"github.com/ryssapp/backend/src/go/gateway-service/delivery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Start() {
	cfg := config.Load()

	// Connect to Store Service
	storeServiceConnection, err := grpc.Dial(cfg.StoreServiceAddress, grpc.WithInsecure())
	if err != nil {
		zap.L().Fatal("Error while connecting to store service.", zap.Error(err))
	}
	zap.L().Info("Connected to store service.")
	defer storeServiceConnection.Close()
	storeService := pb.NewStoreServiceClient(storeServiceConnection)

	// Connect to User Service
	userServiceConnection, err := grpc.Dial(cfg.UserServiceAddress, grpc.WithInsecure())
	if err != nil {
		zap.L().Fatal("Error while connecting to user service.", zap.Error(err))
	}
	zap.L().Info("Connected to user service.")
	defer userServiceConnection.Close()
	userService := pb.NewUserServiceClient(userServiceConnection)

	svr := delivery.NewHTTPServer(cfg.BindAddress)
	svr.StoreService = storeService
	svr.UserService = userService
	svr.Start()
}
