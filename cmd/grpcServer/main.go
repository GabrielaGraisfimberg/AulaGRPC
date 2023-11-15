package main

import (
	"database/sql"
	"net"

	"github.com/gargraisfimberg/Projetos/fullcycle/AulaGRPC/internal/database"
	"github.com/gargraisfimberg/Projetos/fullcycle/AulaGRPC/internal/pb"
	"github.com/gargraisfimberg/Projetos/fullcycle/AulaGRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}