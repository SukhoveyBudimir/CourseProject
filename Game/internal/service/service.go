package service

import (
	"context"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/player"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/repository"
)

type Service struct {
	jwtKey []byte
	rps    repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{rps: pool, jwtKey: jwtKey}
}

// GetUser _
func (se *Service) GetPlayer(ctx context.Context, id string) (*player.Player, error) {
	return se.rps.GetUserByID(ctx, id)
}

// GetAllUsers _
func (se *Service) GetAllPlayer(ctx context.Context) ([]*player.Player, error) {
	return se.rps.GetAllUsers(ctx)
}

// DeleteUser _
func (se *Service) DeletePlayer(ctx context.Context, id string) error {
	return se.rps.DeleteUser(ctx, id)
}

// UpdateUser _
func (se *Service) UpdatePlayer(ctx context.Context, id string, user *player.Player) error {
	return se.rps.UpdateUser(ctx, id, user)
}

func (se *Service) CreatePlayer(ctx context.Context, user *player.Player) (string, error) {
	return se.rps.CreateUser(ctx, user)
}
