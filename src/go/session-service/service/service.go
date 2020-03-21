package service

import (
	redis "github.com/go-redis/redis/v7"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/session-service/config"
	"github.com/ryssapp/backend/src/go/session-service/delivery"
	"github.com/ryssapp/backend/src/go/session-service/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

const ExpirationLength time.Duration = time.Day

func Start() {
	c := config.Load()
	db := initRedis(c.RedisAddress, c.RedisPassword, c.RedisDatabase)
	defer db.Close()

	r := repository.New(db, ExpirationLength)
	srv := delivery.NewServer(r, c.Secret, ExpirationLength)

	lis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		zap.L().Fatal("Failed to start tcp server.", zap.Error(err))
	}
	zap.L().Info("Serving grpc service.", zap.String("address", c.BindAddress))

	grpcServer := grpc.NewServer()
	pb.RegisterSessionServiceServer(grpcServer, srv)
	zap.L().Fatal("Error while serving.", zap.Error(grpcServer.Serve(lis)))
}

func initRedis(addr string, pw string, db int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})
	val, err := c.Ping().Result()
	if err != nil {
		zap.L().Fatal("Failed to ping redis server.", zap.Error(err))
	}
	zap.L().Info("Received ping from redis", zap.String("ping", val))
	return c
}
