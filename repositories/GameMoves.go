package repositories

import "gorm.io/gorm"

type GameMoves struct {
	Id             uint32 `gorm:"primaryKey;autoIncrement" json:"id" column:"id"`
	GameId         uint32 `gorm:"type:int" json:"game_id" column:"game_id"`
	PlayerId       string `gorm:"type:varchar(255)" json:"player_id" column:"player_id"`
	PositionX      uint32 `gorm:"type:int" json:"position_x" column:"position_x"`
	PositionY      uint32 `gorm:"type:int" json:"position_y" column:"position_y"`
	SequenceNumber uint32 `gorm:"type:int" json:"sequence_number" column:"sequence_number"`
}

func (GameMoves) TableName() string {
	return "game_moves"
}

func (g *gameRepository) AddGameMove(tx *gorm.DB, move GameMoves) (resp GameMoves, err error) {
	result := tx.Create(&move)
	if result.Error != nil {
		return resp, result.Error
	}
	return move, nil
}

func (g *gameRepository) GetMovesByGameId(gameId uint32) (resp []GameMoves, err error) {
	result := g.db.Where("game_id = ?", gameId).Find(&resp)
	if result.Error != nil {
		return resp, result.Error
	}
	return resp, nil
}
