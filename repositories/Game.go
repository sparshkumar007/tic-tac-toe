package repositories

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	Id         uint32     `gorm:"primaryKey;autoIncrement" json:"id" column:"id"`
	Status     GameStatus `gorm:"type:varchar(50)" json:"status" column:"status"`
	Winner     string     `gorm:"type:varchar(255)" json:"winner" column:"winner"`
	LastMoveBy string     `gorm:"type:varchar(255)" json:"last_move_by" column:"last_move_by"`
	CreatedAt  int64      `gorm:"autoCreateTime" json:"created_at" column:"created_at"`
	UpdatedAt  int64      `gorm:"autoUpdateTime" json:"updated_at" column:"updated_at"`
}

type GameStatus string

const (
	GameStatusCreated   GameStatus = "CREATED"
	GameStatusInProcess GameStatus = "IN_PROCESS"
	GameStatusCompleted GameStatus = "COMPLETED"
)

type GamePlayers struct {
	GameId  uint32 `gorm:"primaryKey;autoIncrement:false" json:"game_id" column:"game_id"`
	EmailId string `gorm:"primaryKey;autoIncrement:false" json:"email_id" column:"email_id"`
}

type GameRepository interface {
	NewGame(Game) (Game, error)
	NewGameWithTx(*gorm.DB, Game) (Game, error)
	GetGameById(uint32) (Game, error)
	GetPlayersByGameId(uint32) ([]GamePlayers, error)
	AddPlayersToGame(*gorm.DB, []GamePlayers) error
	AddGameMove(*gorm.DB, GameMoves) (GameMoves, error)
	GetMovesByGameId(uint32) ([]GameMoves, error)
	UpdateGameWithTx(*gorm.DB, map[string]interface{}) error
	GetGamesByPlayerEmail(string) ([]Game, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(col *gorm.DB) GameRepository {
	return &gameRepository{
		db: col,
	}
}

func (Game) TableName() string {
	return "game"
}

func (GamePlayers) TableName() string {
	return "game_players"
}

func (g *gameRepository) NewGame(game Game) (resp Game, err error) {
	tx := g.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return g.NewGameWithTx(tx, game)
}

func (g *gameRepository) NewGameWithTx(tx *gorm.DB, game Game) (resp Game, err error) {
	createdAt := time.Now().Unix()
	game.CreatedAt = createdAt
	game.UpdatedAt = createdAt
	result := tx.Create(&game)
	if result.Error != nil {
		return resp, result.Error
	}
	return game, nil
}

func (g *gameRepository) AddPlayersToGame(tx *gorm.DB, gamePlayers []GamePlayers) error {
	result := tx.Create(&gamePlayers)
	return result.Error
}

func (g *gameRepository) GetGameById(id uint32) (Game, error) {
	var game Game
	result := g.db.First(&game, id)
	return game, result.Error
}

func (g *gameRepository) GetPlayersByGameId(gameId uint32) ([]GamePlayers, error) {
	var players []GamePlayers
	result := g.db.Where("game_id = ?", gameId).Find(&players)
	return players, result.Error
}

func (g *gameRepository) UpdateGameWithTx(tx *gorm.DB, updateMap map[string]interface{}) error {
	updateMap["updated_at"] = time.Now().Unix()
	result := tx.Model(&Game{}).Where("id = ?", updateMap["id"]).Updates(updateMap)
	return result.Error
}

func (g *gameRepository) GetGamesByPlayerEmail(emailId string) (resp []Game, err error) {
	var gamePlayers []GamePlayers
	result := g.db.Where("email_id = ?", emailId).Find(&gamePlayers)
	if result.Error != nil {
		return resp, result.Error
	}
	gameIds := make([]uint32, 0)
	for _, gp := range gamePlayers {
		gameIds = append(gameIds, gp.GameId)
	}
	result = g.db.Where("id IN ?", gameIds).Find(&resp)
	if result.Error != nil {
		return resp, result.Error
	}
	return resp, nil
}
