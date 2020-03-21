package service

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/user-service/config"
	"github.com/ryssapp/backend/src/go/user-service/delivery"
	"github.com/ryssapp/backend/src/go/user-service/repository"
	"github.com/ryssapp/backend/src/go/user-service/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func Start() {
	c := config.Load()

	db := initPostgres(c.PostgresConnection)
	r := repository.NewPostgresRepository(db)
	u := usecase.New(r)

	conn, err := grpc.Dial(c.SessionServiceAddress, grpc.WithInsecure())
	if err != nil {
		zap.L().Fatal("Failed to connect to session service.", zap.Error(err))
	}
	defer conn.Close()
	client := pb.NewSessionServiceClient(conn)

	srv := delivery.NewServer(u, c.Cost, client)

	lis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		zap.L().Fatal("Failed to start tcp server.", zap.Error(err))
	}
	zap.L().Info("Serving grpc service.", zap.String("address", c.BindAddress))

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, srv)
	zap.L().Fatal("Error while serving.", zap.Error(grpcServer.Serve(lis)))
}

func initPostgres(addr string) *pg.DB {
	db := connectPostgres(addr)
	err := db.CreateTable(&types.User{}, &orm.CreateTableOptions{
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
