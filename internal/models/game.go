package models

import (
	"time"
)

type Game struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ESPNID       string    `gorm:"type:varchar(20);not null;uniqueIndex"`
	StartTime    time.Time `gorm:"type:date;not null"`
	SeasonYear   string    `gorm:"type:varchar(10)"`
	Type         string    `gorm:"type:varchar(20)"`
	HomeTeamID   string    `gorm:"type:string"`
	AwayTeamID   string    `gorm:"type:string"`
	HomeScore    int       `gorm:"type:integer"`
	HomeQ1Score  int       `gorm:"type:integer"`
	HomeQ2Score  int       `gorm:"type:integer"`
	HomeQ3Score  int       `gorm:"type:integer"`
	HomeQ4Score  int       `gorm:"type:integer"`
	AwayScore    int       `gorm:"type:integer"`
	AwayQ1Score  int       `gorm:"type:integer"`
	AwayQ2Score  int       `gorm:"type:integer"`
	AwayQ3Score  int       `gorm:"type:integer"`
	AwayQ4Score  int       `gorm:"type:integer"`
	Arena        string    `gorm:"type:varchar(100)"`
	Referees     string    `gorm:"type:text"`
	WinnerTeamID string    `gorm:"type:string"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
