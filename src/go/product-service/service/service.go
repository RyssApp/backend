package service

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/product-service/config"
	"github.com/ryssapp/backend/src/go/product-service/delivery"
	"github.com/ryssapp/backend/src/go/product-service/repository"
	"github.com/ryssapp/backend/src/go/product-service/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func Start() {
	cfg := config.Load()
	db := initPostgres(cfg.PostgresConnection)
	r := repository.NewPostgresRepository(db)
	u := usecase.New(r)
	srv := delivery.NewServer(u)
	lis, err := net.Listen("tcp", cfg.BindAddress)
	if err != nil {
		zap.L().Fatal("Error while listening.", zap.String("address", cfg.BindAddress), zap.Error(err))
	}
	zap.L().Info("Listening")
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, srv)
	zap.L().Fatal("Error while serving.", zap.Error(grpcServer.Serve(lis)))
}

func initPostgres(addr string) *pg.DB {
	db := connectPostgres(addr)
	err := db.CreateTable(&types.Product{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		zap.L().Error("Error while creating postgres scheme.", zap.Error(err))
	}
	zap.L().Info("Connected to database")
	return db
}

func connectPostgres(addr string) *pg.DB {
	opts, err := pg.ParseURL(addr)
	if err != nil {
		zap.L().Fatal("Error while connecting to postgres.", zap.Error(err))
	}
	return pg.Connect(opts)
}
