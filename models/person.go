package models

type Person struct {
	PersonId		string	`json:"person_id,omitempty"`
	PersonName 		string 	`json:"person_name,omitempty"`
	PersonDescribe	string	`json:"person_describe,omitempty"`
	PersonSex		string	`json:"person_sex,omitempty"`
	PersonBirthday 	string 	`json:"person_birthday,omitempty"`
	PersonHomeplace string	`json:"person_homeplace,omitempty"`
	PersonAddress 	string 	`json:"person_address,omitempty"`
	PersonImgurl	string	`json:"person_imgurl,omitempty"`
	ParentId		string	`json:"parent_id,omitempty"`
	UserId 			string	`json:"person_user_id,omitempty"`
}

//personlist的条件集合
type PersonSelect struct {
	ConPageNum 			int		`json:"con_page_num"`
	ConPageSize			int		`json:"con_page_size"`
	ConUserId			string	`json:"con_user_id"`
	ConPersonName		string	`json:"con_person_name"`
	ConPersonSex 		string	`json:"con_person_sex"`
	ConPersonBirthday	string	`json:"con_person_birthday"`
	ConPersonHomePlace	string	`json:"con_person_homeplace"`
}
