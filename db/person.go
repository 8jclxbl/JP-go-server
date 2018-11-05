package db

import (
	"JP-go-server/models"
	"JP-go-server/util"
	"database/sql"
	"errors"
	"strconv"
)


func CreatPerson(person models.Person) (string, error){

	stmt, err := dbConn.Prepare("INSERT INTO person " +
		"(personid,personname,persondescribe,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return "",err
	}
	PersonId := util.GenerateId()

	_, err = stmt.Exec(PersonId,person.PersonName,person.PersonDescribe,person.PersonSex,person.PersonBirthday,
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
		"personname,persondescribe,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid " +
		"FROM person WHERE personid = ? ")

	if err != nil {
		return nil, err
	}
	var personname,persondescribe,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	err = stmt.QueryRow(personId).Scan(
		&personname,&persondescribe,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid,)

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
		PersonDescribe:persondescribe,
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
		"personid,personname,persondescribe,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid " +
		"FROM person WHERE personname = ? ")

	if err != nil {
		return nil, err
	}
	//fmt.Println(err.Error())
	var personid,personname,persondescribe,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	err = stmt.QueryRow(userName).Scan(
		&personid,&personname,&persondescribe,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid)

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
		PersonDescribe:persondescribe,
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
	//Change by TIAN
	//Date 2018-11-03
	//personTemp, err := GetPersonByName(person.PersonName)
	
	personTemp, err := GetPersonById(person.PersonId)
	if err != nil {
		return err
	}

	if personTemp == nil {
		return errors.New("人物不存在")
	}

/*
	if person.PersonName == "" {
		person.PersonName = personTemp.PersonName
	}
	if person.PersonSex == "" {
		person.PersonSex = personTemp.PersonSex
	}
	if person.PersonDescribe == "" {
		person.PersonDescribe = personTemp.PersonDescribe
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
	if person.ParentId == "" {
		person.ParentId = personTemp.ParentId
	}
	if person.UserId == "" {
		person.UserId = personTemp.UserId
	}
*/

	stmt, err := dbConn.Prepare("UPDATE person SET personname=?,persondescribe=?,personsex=?," +
		"personbirthday=?,personhomeplace=?,personaddress=?,personimgurl=?,parentid=?,userid=? WHERE personid=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(person.PersonName,person.PersonDescribe,person.PersonSex,person.PersonBirthday,
		person.PersonHomeplace,person.PersonAddress,person.PersonImgurl,person.ParentId,person.UserId,person.PersonId)

	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func ListPerson(personSelect models.PersonSelect) ([]models.Person, error) {
	/*
	sql := "SELECT personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid FROM person WHERE 1"
	*/
	/*
		sql := "SELECT personid,personname,personsex,personbirthday,personhomeplace," +
		"personaddress,personimgurl,parentid,userid FROM person WHERE 1 "
		fmt.Println(personSelect.ConPersonName)
	*/
	sql := "SELECT personid,personname,persondescribe,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid FROM person WHERE 1"

	step := personSelect.ConPageSize
	start := personSelect.ConPageNum * step


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
	if personSelect.ConUserId != "" {
		sql = sql + " AND userid='" + personSelect.ConUserId + "'"
		conditionCount += 1
	}

	
	if conditionCount == 0 {
		return nil,errors.New("未输入任何条件")
	}
	
	sql = sql + " order by id desc limit " + strconv.Itoa(start) + "," + strconv.Itoa(step)
	rows,err := dbConn.Query(sql)
	if err != nil {
		return nil,err
	}
	var personid,personname,persondescribe,personsex,personbirthday,personhomeplace,personaddress,personimgurl,parentid,userid string
	var people []models.Person

	for rows.Next() {
		err = rows.Scan(
			&personid,&personname,&persondescribe,&personsex,&personbirthday,&personhomeplace,&personaddress,&personimgurl,&parentid,&userid)

		personTemp := models.Person{
			PersonId:personid,
			PersonName:personname,
			PersonDescribe:persondescribe,
			PersonSex:personsex,
			PersonBirthday:personbirthday,
			PersonHomeplace:personhomeplace,
			PersonAddress:personaddress,
			PersonImgurl:personimgurl,
			ParentId:parentid,
			UserId:userid,
		}
		people = append(people,personTemp)

		personid = ""
		personname = ""
		persondescribe = ""
		personsex = ""
		personbirthday = ""
		personhomeplace = ""
		personaddress = ""
		personimgurl = ""
		parentid = ""
		userid = ""

	}


	defer rows.Close()
	return people,nil
}

func DeletePerson(id string) error{

	stmt, err := dbConn.Prepare("DELETE FROM person WHERE personid = ? ")

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
