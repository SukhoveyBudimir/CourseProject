package service

import (
	"CourseProject/Registration/internal/model"
	"CourseProject/Registration/internal/repository"
	"context"
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
func (se *Service) GetUser(ctx context.Context, id string) (*model.Player, error) {
	return se.rps.GetUserByID(ctx, id)
}

// GetAllUsers _
func (se *Service) GetAllUsers(ctx context.Context) ([]*model.Player, error) {
	return se.rps.GetAllUsers(ctx)
}

// DeleteUser _
func (se *Service) DeleteUser(ctx context.Context, id string) error {
	return se.rps.DeleteUser(ctx, id)
}

// UpdateUser _
func (se *Service) UpdateUser(ctx context.Context, id string, user *model.Player) error {
	return se.rps.UpdateUser(ctx, id, user)
}
