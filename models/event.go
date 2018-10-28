package models

type Event struct {
	EventId			string		`json:"event_id,omitempty"`
	PersonId		string		`json:"event_person_id,omitempty"`
	EventTitle 		string		`json:"event_title,omitempty"`
	EventContent	string		`json:"event_content,omitempty"`
	EventTime		string		`json:"event_time,omitempty"`
	EventFile		[]File		`json:"event_file,omitempty"`
}

type File struct {
	EventID		string	`json:"file_event_id,omitempty"`
	FileUrl		string	`json:"file_url,omitempty"`
	FileType	string	`json:"file_type,omitempty"`
}

