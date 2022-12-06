package server

import (
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/model"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/service"
	pb "github.com/SukhoveyBudimir/CourseProject/Registration/proto"

	"context"
)

// Server struct
type Server struct {
	pb.UnimplementedGuessTheNumberServer
	se *service.Service
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

// GetPlayer get user by id from db
func (s *Server) GetPlayer(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	accessToken := request.GetAccessToken()
	if err := s.se.Verify(accessToken); err != nil {
		return nil, err
	}
	idPlayer := request.GetId()
	PlayerDB, err := s.se.GetUser(ctx, idPlayer)
	if err != nil {
		return nil, err
	}
	PlayerProto := &pb.GetUserResponse{
		Player: &pb.Player{
			Id:       PlayerDB.ID,
			Name:     PlayerDB.Name,
			Points:   PlayerDB.Points,
			Password: PlayerDB.Password,
		},
	}
	return PlayerProto, nil
}

// GetAllPlayers get all users from db
func (s *Server) GetAllPlayers(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Players, err := s.se.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.Player
	for _, Player := range Players {
		PlayerProto := new(pb.Player)
		PlayerProto.Id = Player.ID
		PlayerProto.Name = Player.Name
		PlayerProto.Points = Player.Points
		list = append(list, PlayerProto)
	}
	return &pb.GetAllResponse{Player: list}, nil
}

// DeletePlayer delete user by id
func (s *Server) DeletePlayer(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Response, error) {
	idUser := request.GetId()
	err := s.se.DeleteUser(ctx, idUser)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// UpdatePlayer update user with new parameters
func (s *Server) UpdatePlayer(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Response, error) {
	accessToken := request.GetAccessToken()
	if err := s.se.Verify(accessToken); err != nil {
		return nil, err
	}
	user := &model.Player{
		Name:   request.Player.Name,
		Points: request.Player.Points,
	}
	idUser := request.GetId()
	err := s.se.UpdateUser(ctx, idUser, user)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
