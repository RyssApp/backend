package main

import (
	"github.com/go-pg/pg/v9/orm"
	"github.com/ryssapp/backend/src/go/store-service/config"
	"github.com/ryssapp/backend/src/go/store-service/store"
	"log"
	"net"

	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/repository"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	opts, err := pg.ParseURL(cfg.PostgresConnection)
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(opts)
	defer db.Close()
	err = db.CreateTable(&store.Store{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	r := repository.NewPostgresRepository(db)

	lis, err := net.Listen("tcp", cfg.BindAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening")
	grpcServer := grpc.NewServer()
	pb.RegisterStoreServiceServer(grpcServer, NewServer(r))
	log.Fatal(grpcServer.Serve(lis))
}
