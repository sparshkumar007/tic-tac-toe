package repositories

import "gorm.io/gorm"

type Game struct {
	Id       uint32   `gorm:"primaryKey;autoIncrement" json:"id"`
	Players  []string `gorm:"type:text[]" json:"players"`
	MetaData string   `gorm:"type:jsonb" json:"meta_data"`
}

type GameRepository interface {
	NewGame(Game) (Game, error)
	GetGameById(uint32) (Game, error)
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

func (g *gameRepository) NewGame(game Game) (resp Game, err error) {
	tx := g.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return NewGameWithTx(tx, game)
}

func NewGameWithTx(tx *gorm.DB, game Game) (resp Game, err error) {
	result := tx.Create(&game)
	if result.Error != nil {
		return resp, result.Error
	}
	return game, nil
}

func (g *gameRepository) GetGameById(id uint32) (Game, error) {
	var game Game
	result := g.db.First(&game, id)
	return game, result.Error
}
