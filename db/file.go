package db

import (
	"JP-go-server/models"
	"database/sql"
	"strconv"
	"errors"
)

func CreatFile(file models.File,eventId string) (string,error) {
	stmt, err := dbConn.Prepare("INSERT INTO file (eventid,fileurl,filetype)" +
		" VALUES (?,?,?)")
	if err != nil {
		return "",err
	}

	res, err  := stmt.Exec(eventId,file.FileUrl,file.FileType)
	if err != nil {
		return "",err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "",err
	}
	Id := strconv.Itoa(int(id))

	defer stmt.Close()
	return Id,nil
}

func GetFile(id string,isFileId bool) (*models.File,error){
	query1 := "SELECT eventid,fileurl,filetype FROM file WHERE id = ?"
	query2 := "SELECT eventid,fileurl,filetype FROM file WHERE eventid = ?"
	query := query2
	if isFileId {
		query = query1
	}

	stmt,err := dbConn.Prepare(query)
	if err != nil {
		return nil,err
	}

	var eventid,fileurl,filetype string
	err = stmt.QueryRow(id).Scan(&eventid,&fileurl,&filetype)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return  nil, nil
	}
	fileTemp := &models.File{
		EventID:eventid,
		FileUrl:fileurl,
		FileType:filetype,
	}

	defer stmt.Close()
	return fileTemp,err
}

func GetFileByUrl(url string) (*models.File,error){
	stmt,err := dbConn.Prepare("SELECT id,eventid,filetype FROM file WHERE fileurl = ?")
	if err != nil {
		return nil,err
	}

	var id,eventid,filetype string
	err = stmt.QueryRow(url).Scan(&id,&eventid,&filetype)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return  nil, nil
	}
	fileTemp := &models.File{
		EventID:eventid,
		FileUrl:url,
		FileType:filetype,
	}

	defer stmt.Close()
	return fileTemp,err
}

func DeleteFile(fileUrl string) error {
	fileTemp, err := GetFileByUrl(fileUrl)
	if err != nil{
		return err
	}
	if fileTemp == nil {
		return errors.New("file not exists")
	}

	stmt, err := dbConn.Prepare("DELETE FROM file " +
		"WHERE fileurl = ? ")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(fileUrl)
	if err != nil {

		return err
	}

	defer stmt.Close()
	return nil

}