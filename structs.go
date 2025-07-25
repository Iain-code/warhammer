package main

import (
	"database/sql"
	"warhammer/internal/db"

	"github.com/google/uuid"
)

type ApiConfig struct {
	db          db.Queries
	tokenSecret string
}

type User struct {
	Id             uuid.UUID    `json:"id"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
	Username       string       `json:"username"`
	HashedPassword string       `json:"-"`
	IsAdmin        bool         `json:"is_admin"`
}

type Model struct {
	OldID       int32  `json:"old_id"`
	DatasheetID int32  `json:"datasheet_id"`
	Name        string `json:"name"`
	M           string `json:"M"`
	T           string `json:"T"`
	Sv          string `json:"Sv"`
	InvSv       string `json:"inv_sv"`
	W           int32  `json:"W"`
	Ld          string `json:"Ld"`
	Oc          int32  `json:"OC"`
}

type Faction struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FactionID string `json:"faction_id"`
}
type Wargear struct {
	DatasheetID int32         `json:"datasheet_id"`
	Id          int32         `json:"id"`
	Name        string        `json:"name"`
	Range       string        `json:"range"`
	Type        string        `json:"type"`
	A           string        `json:"attacks"`
	BsWs        string        `json:"BS_WS"`
	Strength    string        `json:"strength"`
	Ap          sql.NullInt32 `json:"AP"`
	Damage      string        `json:"damage"`
}

type Points struct {
	Id           int32  `json:"id"`
	Datasheet_id int32  `json:"datasheet_id"`
	Line         int32  `json:"line"`
	Description  string `json:"description"`
	Cost         int32  `json:"cost"`
}

type Keyword struct {
	Id          int32  `json:"id"`
	DatasheetID int32  `json:"datasheet_id"`
	Keyword     string `json:"keyword"`
}
