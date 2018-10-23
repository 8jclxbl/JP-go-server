package models


type BaseResponse struct {
	Success bool	`json:"success"`
	Msg 	Message	`json:"msg"`
}

type Message struct {
	Desc string		`json:"desc,omitempty"`
	Userid string	`json:"user_id,omitempty"`
}