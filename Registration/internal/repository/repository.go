package repository

import (
	"context"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, p *model.Player) (string, error)
	GetUserByID(ctx context.Context, idPerson string) (*model.Player, error)
	GetAllUsers(ctx context.Context) ([]*model.Player, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, per *model.Player) error
	SelectByIDAuth(ctx context.Context, id string) (model.Player, error)
	UpdateAuth(ctx context.Context, id string, refreshToken string) error
}
