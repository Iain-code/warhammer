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
	Id             uuid.UUID      `json:"id"`
	CreatedAt      sql.NullTime   `json:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
	Email          string         `json:"email"`
	HashedPassword sql.NullString `json:"-"`
	IsAdmin        bool           `json:"is_admin"`
}

type Model struct {
	OldID       sql.NullInt32  `json:"old_id"`
	DatasheetID string         `json:"datasheet_id"`
	Name        string         `json:"name"`
	M           string         `json:"M"`
	T           int32          `json:"T"`
	Sv          sql.NullString `json:"Sv"`
	InvSv       sql.NullString `json:"inv_sv"`
	W           sql.NullInt32  `json:"W"`
	Ld          sql.NullString `json:"Ld"`
	Oc          sql.NullString `json:"OC"`
}

type Faction struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FactionID string `json:"faction_id"`
}
type Wargear struct {
	DatasheetID int32          `json:"datasheet_id"`
	Field2      sql.NullInt32  `json:"field2"`
	Name        string         `json:"name"`
	Range       sql.NullString `json:"range"`
	Type        sql.NullString `json:"type"`
	A           sql.NullString `json:"attacks"`
	BsWs        sql.NullString `json:"BS_WS"`
	Strength    sql.NullString `json:"strength"`
	Ap          sql.NullString `json:"AP"`
	Damage      sql.NullString `json:"damage"`
}
