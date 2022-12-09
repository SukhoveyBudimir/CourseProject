package server

import (
	"context"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/player"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/service"
	pb "github.com/SukhoveyBudimir/CourseProject/Registration/proto"
)

type Server struct {
	pb.UnimplementedGuessTheNumberServer
	se *service.Service
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

// CreatePlayer create new Player
func (s *Server) CreatePlayer(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	m := player.Player{
		Name:   request.Name,
		Points: request.Points,
	}
	newID, err := s.se.CreatePlayer(ctx, &m)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: newID}, nil
}

// GetPlayer get Player by id from db
func (s *Server) GetPlayer(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	idPlayer := request.GetId()
	PlayerDB, err := s.se.GetPlayer(ctx, idPlayer)
	if err != nil {
		return nil, err
	}
	PlayerProto := &pb.GetUserResponse{
		Player: &pb.Player{
			Id:     PlayerDB.Id,
			Name:   PlayerDB.Name,
			Points: PlayerDB.Points,
		},
	}
	return PlayerProto, nil
}

// GetAllPlayer get all Player from db
func (s *Server) GetAllPlayer(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Players, err := s.se.GetAllPlayer(ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.Player
	for _, Player := range Players {
		PlayerProto := new(pb.Player)
		PlayerProto.Id = Player.Id
		PlayerProto.Name = Player.Name
		PlayerProto.Points = Player.Points
		list = append(list, PlayerProto)
	}
	return &pb.GetAllResponse{Player: list}, nil
}

// DeletePlayer delete Player by id
func (s *Server) DeletePlayer(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Response, error) {
	idMed := request.GetId()
	err := s.se.DeletePlayer(ctx, idMed)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// ChangePlayer update Player with new parameters
func (s *Server) ChangePlayer(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Response, error) {
	pl := &player.Player{
		Name:   request.Player.Name,
		Points: request.Player.Points,
	}
	idMed := request.GetId()
	err := s.se.UpdatePlayer(ctx, idMed, pl)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
