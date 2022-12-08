package Registration

import (
	"context"
	"fmt"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/model"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/repository"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/server"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/service"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/SukhoveyBudimir/CourseProject/Registration/proto"
)

var (
	poolP pgxpool.Pool
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :50051...")
	key := []byte("super-key")
	cfg := model.Config{JwtKey: key}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	conn := DBConnection(&cfg)
	fmt.Println("DB successfully connect...")
	ns := grpc.NewServer()
	newService := service.NewService(conn, cfg.JwtKey)
	srv := server.NewServer(newService)
	pb.RegisterGuessTheNumberServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

// DBConnection create connection with db
func DBConnection(cfg *model.Config) repository.Repository {
	log.Info(cfg.PostgresDBURL)
	poolP, err := pgxpool.Connect(context.Background(), cfg.PostgresDBURL) // "postgresql://postgres:123@localhost:5432/person"
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
		return nil
	}
	return &repository.PRepository{Pool: poolP}
}
