package handlers

import (
	"database/sql"
	"warhammer/internal/db"

	"github.com/google/uuid"
)


type ApiConfig struct {
	Db          *db.Queries
	TokenSecret string
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
	Id          int32  `json:"id"`
	DatasheetID int32  `json:"datasheet_id"`
	Line        int32  `json:"line"`
	Description string `json:"description"`
	Cost        int32  `json:"cost"`
}

type Keyword struct {
	Id          int32  `json:"id"`
	DatasheetID int32  `json:"datasheet_id"`
	Keyword     string `json:"keyword"`
}

type Enhancement struct {
	ID          int32  `json:"id"`
	FactionID   string `json:"faction_id"`
	Name        string `json:"name"`
	Cost        int32  `json:"cost"`
	Detachment  string `json:"detachment"`
	Legend      string `json:"legend"`
	Description string `json:"description"`
	Field8      string `json:"field8"`
}

type Abilities struct {
	DatasheetID int32  `json:"datasheet_id"`
	Line        int32  `json:"line"`
	AbilityID   int32  `json:"ability_id"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Parameter   string `json:"parameter"`
	Field9      string `json:"FIELD9"`
}

type Roster struct {
	Id          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ArmyList    ArmyList  `json:"army_list"`
	Enhancement []string  `json:"enhancements"`
	Name        string    `json:"name"`
	Faction     string    `json:"faction"`
}

type ArmyList struct {
	Character  []Unit `json:"character"`
	Battleline []Unit `json:"battleline"`
	Transport  []Unit `json:"transport"`
	Mounted    []Unit `json:"mounted"`
	Aircraft   []Unit `json:"aircraft"`
	Infantry   []Unit `json:"infantry"`
	Monster    []Unit `json:"monster"`
	Vehicle    []Unit `json:"vehicle"`
}

type UnitPoints struct {
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

type Unit struct {
	OldID       int        `json:"old_id"`
	DatasheetID int        `json:"datasheet_id"`
	Name        string     `json:"name"`
	M           string     `json:"M"`
	T           string     `json:"T"`
	W           int        `json:"W"`
	Sv          string     `json:"Sv"`
	InvSv       string     `json:"inv_sv"`
	Ld          string     `json:"Ld"`
	OC          int        `json:"OC"`
	Instance    int64      `json:"instance"`
	Keywords    []string   `json:"keywords"`
	UnitPoints  UnitPoints `json:"unitPoints"`
}

type KeywordsModel struct {
	Id          int32    `json:"id"`
	DatasheetID int32    `json:"datasheet_id"`
	Keywords    []string `json:"keyword"`
}

type AbilityUpdate struct {
	DatasheetID int    `json:"datasheet_id"`
	Name        string `json:"name"`
	Line        int32  `json:"line"`
	Description string `json:"description"`
}
