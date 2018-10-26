package models


type UserResponse struct {
	Success bool	`json:"success"`
	Msg 	UserMessage	`json:"msg"`
}

type UserMessage struct {
	Desc 		string	`json:"desc,omitempty"`
	Userid 		string	`json:"user_id,omitempty"`

}

type PersonResponse struct {
	Success bool	`json:"success"`
	Msg 	PersonMessage	`json:"msg"`
}

type PersonMessage struct {
	Desc 		string	`json:"desc,omitempty"`
	PersonInfo 	Person	`json:"person_info,omitempty"`
}