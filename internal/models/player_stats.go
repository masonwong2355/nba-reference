package models

import (
	"time"
)

type PlayerStats struct {
	ID           string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	GameEspnID   string `gorm:"type:varchar(20);"`
	PlayerEspnID string `gorm:"type:varchar(20);"`
	TeamEspnID   string `gorm:"type:varchar(20);"`

	Min         int `gorm:"type:integer"`
	FGMade      int `gorm:"type:integer"`
	FGAtt       int `gorm:"type:integer"`
	ThreeptMade int `gorm:"type:integer"`
	ThreeptAtt  int `gorm:"type:integer"`
	FtMade      int `gorm:"type:integer"`
	FtAtt       int `gorm:"type:integer"`
	Oreb        int `gorm:"type:integer"`
	Dreb        int `gorm:"type:integer"`
	Reb         int `gorm:"type:integer"`
	Ast         int `gorm:"type:integer"`
	Stl         int `gorm:"type:integer"`
	Blk         int `gorm:"type:integer"`
	Turnover    int `gorm:"type:integer"`
	Pf          int `gorm:"type:integer"`
	Pts         int `gorm:"type:integer"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
