package service

import (
	"context"
	"tic-tac-toe/repositories"

	"gorm.io/gorm"
)

type GameService interface {
	GetGame(ctx context.Context, req GetGameRequest) (resp GetGameResponse, err error)
	CreateGame(ctx context.Context, req CreateGameRequest) (resp CreateGameResponse, err error)
	AddGameMove(ctx context.Context, req AddGameMoveRequest) (resp AddGameMoveResponse, err error)
	GetAllGamesForPlayer(ctx context.Context, emailId string) (resp GetAllGamesForPlayerResponse, err error)
	GetGameBoard(ctx context.Context, req GetGameBoardRequest) (resp GetGameBoardResponse, err error)
}

type gameService struct {
	GameRepository repositories.GameRepository
	db             *gorm.DB
}

func NewGameService(gameRepo repositories.GameRepository, db *gorm.DB) GameService {
	return &gameService{
		GameRepository: gameRepo,
		db:             db,
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
	gamePlayersResp, err := s.GameRepository.GetPlayersByGameId(req.GameId)
	if err != nil {
		return resp, err
	}
	gamePlayers := make([]string, 0)
	for _, player := range gamePlayersResp {
		gamePlayers = append(gamePlayers, player.EmailId)
	}
	resp = GetGameResponse{
		GameId:  req.GameId,
		Players: gamePlayers,
		Winner:  game.Winner,
		Status:  repositories.GameStatusCreated,
	}
	return resp, nil
}

func (s *gameService) CreateGame(ctx context.Context, req CreateGameRequest) (resp CreateGameResponse, err error) {
	err = ValidateCreateGameRequest(req)
	if err != nil {
		return resp, err
	}
	tx := s.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	game, err := s.GameRepository.NewGameWithTx(tx, repositories.Game{
		Status: repositories.GameStatusCreated,
	})
	if err != nil {
		return resp, err
	}
	gamePlayers := []repositories.GamePlayers{}
	for _, email := range req.EmailIds {
		gamePlayers = append(gamePlayers, repositories.GamePlayers{
			GameId:  game.Id,
			EmailId: email,
		})
	}
	err = s.GameRepository.AddPlayersToGame(tx, gamePlayers)
	if err != nil {
		return resp, err
	}
	resp = CreateGameResponse{
		GameId: game.Id,
	}
	return
}

func (s *gameService) AddGameMove(ctx context.Context, req AddGameMoveRequest) (resp AddGameMoveResponse, err error) {
	err = ValidateAddGameMoveRequest(req)
	if err != nil {
		return resp, err
	}
	tx := s.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	game, err := s.GameRepository.GetGameById(req.GameId)
	if err != nil {
		return resp, err
	}

	moves, err := s.GameRepository.GetMovesByGameId(req.GameId)
	if err != nil {
		return resp, err
	}
	err = valiateMove(game, moves, repositories.GameMoves{
		PlayerId:  req.PlayerId,
		PositionX: req.PositionX,
		PositionY: req.PositionY,
	})
	if err != nil {
		return resp, err
	}

	seqNumber := uint32(len(moves) + 1)
	addMoveResp, err := s.GameRepository.AddGameMove(tx, repositories.GameMoves{
		GameId:         req.GameId,
		PlayerId:       req.PlayerId,
		PositionX:      req.PositionX,
		PositionY:      req.PositionY,
		SequenceNumber: seqNumber,
	})
	if err != nil {
		return resp, err
	}
	moves = append(moves, repositories.GameMoves{
		GameId:         req.GameId,
		PlayerId:       req.PlayerId,
		PositionX:      req.PositionX,
		PositionY:      req.PositionY,
		SequenceNumber: seqNumber,
	})
	updateMap := map[string]interface{}{
		"id":           req.GameId,
		"last_move_by": req.PlayerId,
		"status":       repositories.GameStatusInProcess,
	}
	winner, isCompleted := checkWinner(moves)
	if isCompleted {
		updateMap["winner"] = winner
		updateMap["status"] = repositories.GameStatusCompleted
	}
	err = s.GameRepository.UpdateGameWithTx(tx, updateMap)
	if err != nil {
		return resp, err
	}

	return AddGameMoveResponse{
		MoveId: addMoveResp.Id,
	}, nil
}

func (s *gameService) GetAllGamesForPlayer(ctx context.Context, emailId string) (resp GetAllGamesForPlayerResponse, err error) {
	games, err := s.GameRepository.GetGamesByPlayerEmail(emailId)
	if err != nil {
		return resp, err
	}
	gameResponses := make([]GetGameResponse, 0)
	for _, game := range games {
		gamePlayersResp, err := s.GameRepository.GetPlayersByGameId(game.Id)
		if err != nil {
			return resp, err
		}
		gamePlayers := make([]string, 0)
		for _, player := range gamePlayersResp {
			gamePlayers = append(gamePlayers, player.EmailId)
		}
		gameResponses = append(gameResponses, GetGameResponse{
			GameId:  game.Id,
			Players: gamePlayers,
			Winner:  game.Winner,
			Status:  game.Status,
		})
	}
	resp = GetAllGamesForPlayerResponse{
		Games: gameResponses,
	}
	return resp, nil
}

func (s *gameService) GetGameBoard(ctx context.Context, req GetGameBoardRequest) (resp GetGameBoardResponse, err error) {
	moves, err := s.GameRepository.GetMovesByGameId(req.GameId)
	if err != nil {
		return resp, err
	}
	for _, move := range moves {
		resp.Board[move.PositionX][move.PositionY] = req.PlayerToTokenMap[move.PlayerId]
	}
	game, err := s.GameRepository.GetGameById(req.GameId)
	if err != nil {
		return resp, err
	}
	resp.Status = game.Status
	resp.Winner = game.Winner
	return resp, nil
}
