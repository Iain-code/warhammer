package main

type Model struct {
	DatasheetID int    `json:"datasheet_id"`
	Name        string `json:"name"`
	M           string `json:"M"`
	T           string `json:"T"`
	Sv          string `json:"Sv"`
	InvSv       string `json:"inv_sv"`
	W           int    `json:"W"`
	Ld          string `json:"Ld"`
	Oc          int    `json:"OC"`
}
