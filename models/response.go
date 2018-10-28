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
	Success bool			`json:"success"`
	Msg 	PersonMessage	`json:"msg"`
}

type PersonMessage struct {
	Desc 		string		`json:"desc,omitempty"`
	PersonInfo 	Person		`json:"person_info,omitempty"`
	PeopleList 	[]Person	`json:"people_list,omitempty"`
}

type EventResponse struct {
	Success 	bool			`json:"success"`
	Msg 		EventMessage	`json:"msg"`
}

type EventMessage struct {
	Desc		string	`json:"desc"`
	EventInfo	Event	`json:"event_info,omitempty"`
	EventList 	[]Event	`json:"event_list"`
}


type FileResponse struct {
	Success bool		`json:"success"`
	Msg		FileMessage	`json:"msg"`
}

type FileMessage struct {
	Desc 		string	`json:"desc"`
	FileUrl		string	`json:"file_url,omitempty"`
}