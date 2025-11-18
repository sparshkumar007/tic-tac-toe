package service

type GetGameRequest struct {
	GameId uint32 `json:"game_id"`
}

type GetGameResponse struct {
	GameId  uint32   `json:"game_id"`
	Players []string `json:"players"`
}
