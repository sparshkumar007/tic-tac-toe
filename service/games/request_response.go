package service

import "tic-tac-toe/repositories"

type GetGameRequest struct {
	GameId uint32 `json:"game_id"`
}

type GetGameResponse struct {
	GameId  uint32                  `json:"game_id"`
	Players []string                `json:"players"`
	Winner  string                  `json:"winner"`
	Status  repositories.GameStatus `json:"status"`
}

type CreateGameRequest struct {
	EmailIds []string `json:"email_ids"`
}

type CreateGameResponse struct {
	GameId uint32 `json:"game_id"`
}

type AddGameMoveRequest struct {
	GameId    uint32 `json:"game_id"`
	PlayerId  string `json:"player_id"`
	PositionX uint32 `json:"position_x"`
	PositionY uint32 `json:"position_y"`
}

type AddGameMoveResponse struct {
	MoveId uint32 `json:"move_id"`
}

type GetAllGamesForPlayerResponse struct {
	Games []GetGameResponse `json:"games"`
}

type GetGameBoardRequest struct {
	GameId           uint32            `json:"game_id"`
	PlayerToTokenMap map[string]string `json:"player_to_token_map"`
}

type GetGameBoardResponse struct {
	Board  [3][3]string            `json:"board"`
	Status repositories.GameStatus `json:"status"`
	Winner string                  `json:"winner"`
}
