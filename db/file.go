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

func UpdateFile (file models.File) error {

	fileTemp, err := GetFileByUrl(file.FileUrl)
	if err != nil{
		return err
	}
	if fileTemp == nil {
		return errors.New("数据库中无记录")
	}

	if file.FileType == "" {
		file.FileType = fileTemp.FileType
	}
	if file.EventID == "" {
		file.EventID = fileTemp.EventID
	}


	query := "UPDATE file SET eventid = ?,filetype = ? WHERE fileurl = ?"
	stmt, err := dbConn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(file.EventID,file.FileType,file.FileUrl)

	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func GetFiles(eventid string) ([]models.File,error){

	rows,err := dbConn.Query("SELECT fileurl,filetype FROM file WHERE eventid = ? ORDER BY id DESC", eventid)
	if err != nil {
		return nil,err
	}

	var fileurl,filetype string
	var files []models.File

	for rows.Next() {
		err = rows.Scan(&fileurl,&filetype,)
		if fileurl == "" {
			fileurl = defaultPic
		}

		fileTemp := models.File{
			EventID:eventid,
			FileUrl:fileurl,
			FileType:filetype,
		}
		fileurl = ""
		filetype = ""

		files = append(files,fileTemp)
	}
	defer rows.Close()
	return files,err
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

	stmt, err := dbConn.Prepare("DELETE FROM file WHERE fileurl = ? ")

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