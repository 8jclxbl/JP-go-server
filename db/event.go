package db

import (
	"JP-go-server/models"
	"JP-go-server/util"
	"database/sql"
	"github.com/kataras/iris/core/errors"
)

func CreatEvent(event models.Event) (string,error){

	stmt, err := dbConn.Prepare("INSERT INTO event " +
		"(eventid,eventtitle,eventcontent,eventtime,personid) " +
		"VALUES (?,?,?,?,?)")

	if err != nil {
		return "",err
	}
	EventId := util.GenerateId()

	var eventTime interface{}
	if event.EventTime == "" {
		eventTime = sql.NullString{}
	} else {
		eventTime = event.EventTime
	}

	_, err = stmt.Exec(EventId,event.EventTitle,event.EventContent,eventTime,event.PersonId)

	if err != nil {
		return "",err
	}
	/*
	id, err := res.LastInsertId()
	if err != nil {
		return "",err
	}
	Id := strconv.Itoa(int(id))
	*/
	defer stmt.Close()
	return EventId,nil
}

func UpdateEvent(event models.Event) error{
	eventTemp, err := GetEventById(event.EventId)
	if eventTemp == nil {
		return errors.New("时间不存在")
	}

	if err != nil {
		return err
	}


	if event.EventTitle == "" {
		event.EventTitle = eventTemp.EventTitle
	}
	if event.EventContent == "" {
		event.EventContent = eventTemp.EventContent
	}
	if event.EventTime == "" {
		event.EventTime = eventTemp.EventTime
	}
	if event.PersonId == "" {
		event.PersonId = eventTemp.PersonId
	}


	stmt, err := dbConn.Prepare("UPDATE event SET eventtitle=?,eventcontent=?," +
		"eventtime=?,personid=? WHERE eventid=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.EventTitle,event.EventContent,event.EventTime,event.PersonId,event.EventId)

	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func GetByPersonId(id string) ([]models.Event,error) {
	rows,err := dbConn.Query("SELECT eventid,eventtitle,eventcontent,eventtime FROM event WHERE personid = ? ORDER BY id DESC",id)
	if err != nil {
		return nil,err
	}
	var eventid,eventtitle,eventcontent,eventtime string
	var events []models.Event


	for rows.Next() {
		err = rows.Scan(&eventid,&eventtitle,&eventcontent,&eventtime)
		eventTemp := models.Event{
			EventId:eventid,
			EventTitle:eventtitle,
			EventContent:eventcontent,
			EventTime:eventtime,
			PersonId:id,
		}
		eventid = ""
		eventtitle = ""
		eventcontent = ""
		eventtime = ""

		events = append(events,eventTemp)
	}

	defer rows.Close()
	return events, nil
}

func GetEventById(id string) (*models.Event,error) {
	stmt,err := dbConn.Prepare("SELECT eventtitle,eventcontent," +
		"eventtime,personid FROM event WHERE eventid = ?")
	if err != nil {
		return nil,err
	}

	var eventtitle,eventcontent,eventtime,personid string
	err = stmt.QueryRow(id).Scan(&eventtitle,&eventcontent,
		&eventtime,&personid)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return  nil, nil
	}

	eventTemp := &models.Event{
		EventId:id,
		EventTitle:eventtitle,
		EventContent:eventcontent,
		EventTime:eventtime,
		PersonId:personid,
	}

	defer stmt.Close()
	return eventTemp, nil
}

func DeleteEvent(id string) error {

	stmt, err := dbConn.Prepare("DELETE FROM event WHERE eventid = ? ")

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


