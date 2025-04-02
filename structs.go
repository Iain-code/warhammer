package main

import (
	"database/sql"
	"warhammer/internal/db"

	"github.com/google/uuid"
)

type ApiConfig struct {
	db db.Queries
}

type User struct {
	Id             uuid.UUID      `json:"id"`
	CreatedAt      sql.NullTime   `json:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
	Email          string         `json:"email"`
	HashedPassword sql.NullString `json:"-"`
}

type Model struct {
	OldID       int     `json:"old_id"`
	DatasheetID float64 `json:"datasheet_id"`
	Name        string  `json:"name"`
	M           string  `json:"M"`
	T           int     `json:"T"`
	Sv          string  `json:"Sv"`
	InvSv       string  `json:"inv_sv"`
	W           int     `json:"W"`
	Ld          string  `json:"Ld"`
	Oc          int     `json:"OC"`
}

type Faction struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FactionID string `json:"faction_id"`
}

type Wargear struct {
	DatasheetID int    `json:"datasheet_id"`
	Field2      int    `json:"field2"`
	Name        string `json:"name"`
	Range       string `json:"range"`
	Type        string `json:"type"`
	A           string `json:"attacks"`
	BS_WS       string `json:"BS_WS"`
	S           string `json:"strength"`
	AP          int    `json:"AP"`
	D           string `json:"damage"`
}
