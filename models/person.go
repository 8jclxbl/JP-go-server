package models

type Person struct {
	PersonId		string	`json:"person_id"`
	PersonName 		string 	`json:"person_name,omitempty"`
	PersonSex		string	`json:"person_sex,omitempty"`
	PersonBirthday 	string 	`json:"person_birthday,omitempty"`
	PersonHomeplace string	`json:"person_homeplace,omitempty"`
	PersonAddress 	string 	`json:"person_address,omitempty"`
	PersonImgurl	string	`json:"person_imgurl,omitempty"`
	ParentId		string	`json:"parent_id,omitempty"`
	UserId 			string	`json:"user_id,omitempty"`
}
