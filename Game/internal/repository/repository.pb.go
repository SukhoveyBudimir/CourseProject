package repository

import (
	"context"
	"fmt"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/player"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type PRepository struct {
	Pool *pgxpool.Pool
}

// CreatePlayer add Player to db
func (p *PRepository) CreatePlayer(ctx context.Context, pl *player.Player) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into Player(id,name,count,price) values($1,$2,$3,$4)",
		newID, &pl.Name, &pl.Points)
	if err != nil {
		log.Errorf("database error with create Player: %v", err)
		return "", err
	}
	return newID, nil
}

// GetPlayerByID select Player by id
func (p *PRepository) GetPlayerByID(ctx context.Context, idPlayer string) (*player.Player, error) {
	u := player.Player{}
	err := p.Pool.QueryRow(ctx, "select id,name,count,price from Player where id=$1", idPlayer).Scan(
		&u.Id, &u.Name, &u.Points)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &player.Player{}, fmt.Errorf("Player with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return &player.Player{}, err
	}
	return &u, nil
}

// GetAllPlayer select all Players from db
func (p *PRepository) GetAllPlayer(ctx context.Context) ([]*player.Player, error) {
	var Players []*player.Player
	rows, err := p.Pool.Query(ctx, "select id,name,count,price from Players")
	if err != nil {
		log.Errorf("database error with select all Players, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		pl := player.Player{}
		err = rows.Scan(&pl.Id, &pl.Name, &pl.Points)
		if err != nil {
			log.Errorf("database error with select all Players, %v", err)
			return nil, err
		}
		Players = append(Players, &pl)
	}

	return Players, nil
}

// DeletePlayer delete Player by id
func (p *PRepository) DeletePlayer(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from Players where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("Player with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("Player with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete Player %v", err)
		return err
	}
	return nil
}

// ChangePlayer update parameters for Player
func (p *PRepository) ChangePlayer(ctx context.Context, id string, pl *player.Player) error {
	a, err := p.Pool.Exec(ctx, "update Players set name=$1,count=$2,price=$3 where id=$4", &pl.Name, &pl.Points, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("Player with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update Playerr %v", err)
		return err
	}
	return nil
}
