package repository

import (
	"context"
	"github.com/SukhoveyBudimir/CourseProject/Registration/internal/model"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

// Createmodel add model to db
func (p *PRepository) Createmodel(ctx context.Context, person *model.Player) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into persons(id,name,position,password) values($1,$2,$3,$4)",
		newID, &person.Name, &person.Points, &person.Password)
	if err != nil {
		log.Errorf("database error with create model: %v", err)
		return "", err
	}
	return newID, nil
}

// GetmodelByID select model by id
func (p *PRepository) GetmodelByID(ctx context.Context, idPlayer string) (*model.Player, error) {
	u := model.Player{}
	err := p.Pool.QueryRow(ctx, "select id,name,position,password from persons where id=$1", idPlayer).Scan(
		&u.ID, &u.Name, &u.Points, &u.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &model.Player{}, fmt.Errorf("model with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return &model.Player{}, err
	}
	return &u, nil
}

// GetAllmodels select all models from db
func (p *PRepository) GetAllmodels(ctx context.Context) ([]*model.Player, error) {
	var players []*model.Player
	rows, err := p.Pool.Query(ctx, "select id,name,position from persons")
	if err != nil {
		log.Errorf("database error with select all models, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		per := model.Player{}
		err = rows.Scan(&per.ID, &per.Name, &per.Points)
		if err != nil {
			log.Errorf("database error with select all models, %v", err)
			return nil, err
		}
		players = append(players, &per)
	}

	return players, nil
}

// Deletemodel delete model by id
func (p *PRepository) Deletemodel(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from persons where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("model with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("model with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete model %v", err)
		return err
	}
	return nil
}

// Updatemodel update parameters for model
func (p *PRepository) Updatemodel(ctx context.Context, id string, per *model.Player) error {
	a, err := p.Pool.Exec(ctx, "update persons set name=$1,position=$2 where id=$3", &per.Name, &per.Points, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("model with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update model %v", err)
		return err
	}
	return nil
}

// UpdateAuth logout, delete refresh token
func (p *PRepository) UpdateAuth(ctx context.Context, id, refreshToken string) error {
	a, err := p.Pool.Exec(ctx, "update persons set refreshToken=$1 where id=$2", refreshToken, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("model with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update model %v", err)
		return err
	}
	return nil
}

// SelectByIDAuth get id and refresh token by id
func (p *PRepository) SelectByIDAuth(ctx context.Context, id string) (model.Player, error) {
	per := model.Player{}
	err := p.Pool.QueryRow(ctx, "select id,refreshToken from persons where id=$1", id).Scan(&per.ID, &per.RefreshToken)

	if err != nil /*err==no-records*/ {
		if err == pgx.ErrNoRows {
			return model.Player{}, fmt.Errorf("model with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return model.Player{}, err /*p, fmt.errorf("model with this id doesn't exist")*/
	}
	return per, nil
}
