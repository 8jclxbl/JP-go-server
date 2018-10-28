package db

import (
	"JP-go-server/models"
	"JP-go-server/util"
	"database/sql"
	"errors"
	"fmt"
)


const PageMax = 20
func CreatPerson(person models.Person) (string, error){

	stmt, err := dbConn.Prepare("INSERT INTO person " +
		"(personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid) " +
		"VALUES (?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return "",err
	}

	PersonId := util.GenerateId()

	_, err = stmt.Exec(PersonId,person.PersonName,person.PersonSex,person.PersonBirthday,
		person.PersonHomeplace,person.PersonAddress,person.PersonImgurl,
		person.ParentId,person.UserId)

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
		"FROM person WHERE personname = ? ")

	if err != nil {
		return nil, err
	}
	fmt.Println(err.Error())
	var personid,personname,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	err = stmt.QueryRow(userName).Scan(
		&personid,&personname,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid)

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
	if personTemp == nil {
		return errors.New("person not exists")
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

func ListPerson(personSelect models.PersonSelect) ([]models.Person, error) {
	sql := "SELECT personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid FROM person WHERE 1"

	//如果json没有包含任何条件，会导致输出数据库中的全部数据，个人不知具体业务逻辑是否需要这样，这里先这样处理，如果没有条件时会报错
	conditionCount := 0
	if personSelect.ConPersonSex != "" {
		sql = sql + " AND personsex='" + personSelect.ConPersonSex + "'"
		conditionCount += 1
	}
	if personSelect.ConPersonBirthday != "" {
		sql = sql + " AND personbirthday='" + personSelect.ConPersonBirthday + "'"
		conditionCount += 1
	}
	if personSelect.ConPersonHomePlace != "" {
		sql = sql + " AND personhomeplace='" + personSelect.ConPersonHomePlace + "'"
		conditionCount += 1
	}
	if personSelect.ConPersonName != "" {
		sql = sql + " AND personname='" + personSelect.ConPersonName + "'"
		conditionCount += 1
	}

	if conditionCount == 0 {
		return nil,errors.New("no conditions")
	}

	rows,err := dbConn.Query(sql)
	if err != nil {
		return nil,err
	}
	var personid,personname,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	var people []models.Person

	for rows.Next() {
		err = rows.Scan(
			&personid,&personname,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid)
		personTemp := models.Person{
			PersonId:personid,
			PersonName:personname,
			PersonSex:personsex,
			PersonBirthday:personbirthday,
			PersonHomeplace:personhomeplace,
			PersonAddress:personaddress,
			PersonImgurl:personimgurl,
			UserId:userid,
		}
		people = append(people,personTemp)
	}
	maxItems := personSelect.ConPageNum * PageMax
	if maxItems < len(people) {
		people=people[:maxItems-1]
	}

	defer rows.Close()
	return people,nil
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
