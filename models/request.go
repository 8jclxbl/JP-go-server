package models


type JsonRequest struct {
	Params JsonData	`json:"params"`
}

type JsonData struct {
	Action string	`json:"action"`
	User
}
