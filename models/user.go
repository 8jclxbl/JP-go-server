package models

import (
	"encoding/json"
	"fmt"
	"bytes"
)

type User struct {
	UserNickname 	string	`json:"user_nickname"`
	UserName		string	`json:"user_name"`
	UserPass		string	`json:"user_pass"`
	UserSex			string	`json:"user_sex"`
	UserBirthday	string	`json:"user_birthday"`
	UserPhone		string	`json:"user_phone"`
	UserEmail		string	`json:"user_email"`
	UserHomeplace	string	`json:"user_homeplace"`
	UserAddress		string	`json:"user_address"`
	UserImgurl		string	`json:"user_imgurl"`

	UserId 			string	`json:"user_id"`
	PersonId		string	`json:"person_id"`
}

func (u *User) String() string {
	b, err := json.Marshal(*u)
	if err != nil {
		return fmt.Sprintf("%+v", *u)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *u)
	}
	return out.String()
}