package db

import (
	"JP-go-server/models"
	"JP-go-server/util"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func CreatPerson(person models.Person) (string, error){

	stmt, err := dbConn.Prepare("INSERT INTO person " +
		"(personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid) " +
		"VALUES (?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return "",err
	}
	userId, _ := strconv.Atoi(person.UserId)
	PersonId := util.GenerateId()

	_, err = stmt.Exec(PersonId,person.PersonName,person.PersonSex,person.PersonBirthday,
		person.PersonHomeplace,person.PersonAddress,person.PersonImgurl,
		person.ParentId,userId)

	if err != nil {
		return "",err
	}

	defer stmt.Close()
	return PersonId,nil
}

func GetPersonById(personId string) (*models.Person, error) {
	stmt, err := dbConn.Prepare("SELECT " +
		"personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid " +
		"FROM person Where personid = ? ")

	if err != nil {
		return nil, err
	}
	var personname,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	err = stmt.QueryRow(personId).Scan(
		&personname,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid,)

	if err != nil && err != sql.ErrNoRows {
		//fmt.Println("content err")
		return nil, err
	}

	if err == sql.ErrNoRows {
		//fmt.Println("sql err")
		return  nil, nil
	}

	personTemp := &models.Person{
		PersonId:personId,
		PersonName:personname,
		PersonSex:personsex,
		PersonBirthday:personbirthday,
		PersonHomeplace:personhomeplace,
		PersonAddress:personaddress,
		PersonImgurl:personimgurl,
		UserId:userid,
	}

	defer stmt.Close()
	return personTemp, nil
}


func GetPersonByName(userName string) (*models.Person, error) {
	stmt, err := dbConn.Prepare("SELECT " +
		"personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid " +
		"FROM person Where personname = ? ")

	if err != nil {
		return nil, err
	}
	fmt.Println(err.Error())
	var personid,personname,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	err = stmt.QueryRow(userName).Scan(
		&personid,&personname,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid,)

	if err != nil && err != sql.ErrNoRows {
		//fmt.Println("content err")
		return nil, err
	}

	if err == sql.ErrNoRows {
		//fmt.Println("sql err")
		return  nil, nil
	}

	personTemp := &models.Person{
		PersonId:personid,
		PersonName:personname,
		PersonSex:personsex,
		PersonBirthday:personbirthday,
		PersonHomeplace:personhomeplace,
		PersonAddress:personaddress,
		PersonImgurl:personimgurl,
		UserId:userid,
	}


	defer stmt.Close()
	return personTemp, nil
}

func UpdatePerson(person models.Person) error{
	personTemp, err := GetPersonByName(person.PersonName)
	if err != nil {
		return err
	}


	if person.PersonName == "" {
		person.PersonName = personTemp.PersonName
	}
	if person.PersonSex == "" {
		person.PersonSex = personTemp.PersonSex
	}
	if person.PersonBirthday == "" {
		person.PersonBirthday = personTemp.PersonBirthday
	}
	if person.PersonHomeplace == "" {
		person.PersonHomeplace = personTemp.PersonHomeplace
	}
	if person.PersonAddress == "" {
		person.PersonAddress = personTemp.PersonAddress
	}
	if person.PersonImgurl == "" {
		person.PersonImgurl = personTemp.PersonImgurl
	}
	if person.PersonSex == "" {
		person.PersonSex = personTemp.PersonSex
	}
	if person.ParentId == "" {
		person.ParentId = personTemp.ParentId
	}


	stmt, err := dbConn.Prepare("UPDATE person SET personname=?,personsex=?," +
		"personbirthday=?,personhomeplace=?,personaddress=?,personimgurl=?," +
		"parentid=? WHERE personname=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(person.PersonName,person.PersonSex,person.PersonBirthday,
		person.PersonHomeplace,person.PersonAddress,person.PersonImgurl,person.ParentId,person.PersonName)

	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func DeletePerson(id string) error{
	personTemp, err := GetPersonById(id)
	if err != nil {
		return err
	}

	if personTemp == nil {
		return errors.New("person not exists")
	}

	stmt, err := dbConn.Prepare("DELETE FROM person " +
		"WHERE personid = ? ")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {

		return err
	}


	defer stmt.Close()
	return nil
}
