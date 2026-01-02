package service

import (
	"context"
	"tic-tac-toe/repositories"
)

type GameService interface {
	GetGame(ctx context.Context, req GetGameRequest) (resp GetGameResponse, err error)
}

type gameService struct {
	GameRepository repositories.GameRepository
}

func NewGameService(gameRepo repositories.GameRepository) GameService {
	return &gameService{
		GameRepository: gameRepo,
	}
}

func (s *gameService) GetGame(ctx context.Context, req GetGameRequest) (resp GetGameResponse, err error) {
	err = ValidateGetGameRequest(req)
	if err != nil {
		return resp, err
	}
	game, err := s.GameRepository.GetGameById(req.GameId)
	if err != nil {
		return resp, err
	}

	resp = GetGameResponse{
		GameId:  req.GameId,
		Players: game.Players,
	}
	return resp, nil
}

func (s *gameService) CreateGame(ctx context.Context, req CreateGameRequest) (resp CreateGameResponse, err error) {
	return
}
