package service

type GetGameRequest struct {
	GameId uint32 `json:"game_id"`
}

type GetGameResponse struct {
	GameId  uint32   `json:"game_id"`
	Players []string `json:"players"`
}

type CreateGameRequest struct {
	Player1 uint32
}

type CreateGameResponse struct {
	GameId uint32
}
