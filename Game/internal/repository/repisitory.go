package repository

import (
	"context"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/player"
)

type Repository interface {
	CreateUser(ctx context.Context, p *player.Player) (string, error)
	GetUserByID(ctx context.Context, idPerson string) (*player.Player, error)
	GetAllUsers(ctx context.Context) ([]*player.Player, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, per *player.Player) error
	SelectByIDAuth(ctx context.Context, id string) (player.Player, error)
	UpdateAuth(ctx context.Context, id string, refreshToken string) error
}
