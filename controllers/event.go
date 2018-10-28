package controllers

import (
	"JP-go-server/db"
	"JP-go-server/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type EventController struct {
	beego.Controller
	jsReq models.JsonRequest
	msg  models.EventMessage
}

//根据读到的方法名选择相应的破石头方法
func (this *EventController) Post() {
	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	actionFromUrl := this.Ctx.Input.Param(":action")


	//读取出错时的回复
	if err != nil {
		this.msg.Desc = "body read err: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &this.jsReq); err != nil {
		this.msg.Desc = "json parse read err: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	fmt.Println(this.jsReq)

	action := this.jsReq.Params.Action

	if action != actionFromUrl {
		this.msg.Desc = "actions from url and json are not the same "
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	switch action {
	case "create":
		this.Create()
	case "update":
		this.Update()
	default:
		this.msg.Desc = "Unknown method"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

//根据读到的方法名选择相应的get方法
func (this *EventController) Get() {
	action:= this.Ctx.Input.Param(":action")

	switch action {
	case "delete":
		this.Delete()
	case "view":
		this.View()
	case "list":
		this.List()
	default:
		this.msg.Desc = "Unknown method"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

//创建事件
func (this *EventController) Create() {
	Event := this.jsReq.Params.Event
	eventID,err := db.CreatEvent(Event)
	//写入时出错的回复
	if err != nil {
		this.msg.Desc = "db err: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc = "event created success"
	this.msg.EventInfo.EventId = eventID
	resp := GenEventResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//更新事件
func (this *EventController) Update() {
	//上面的错误处理与注册一致
	event := this.jsReq.Params.Event
	//更新数据库的数据
	fmt.Println(event.PersonId)
	err := db.UpdateEvent(event)
	if err != nil {
		this.msg.Desc ="update failed " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	this.msg.Desc = "updated"
	resp := GenEventResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//删除事件
func (this *EventController) Delete() {
	id:= this.Ctx.Input.Param(":event_id")
	id = id[1:]
	err := db.DeleteEvent(id)
	if err != nil {
		this.msg.Desc ="delete failed:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.Desc ="delete success"
	resp := GenEventResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//展示指定id的事件
func (this *EventController) View() {
	id:= this.Ctx.Input.Param(":event_id")
	id = id[1:]
	eventTemp,err := db.GetEventById(id)

	if eventTemp == nil {
		this.msg.Desc ="event not exists"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if err != nil {
		this.msg.Desc ="delete failed:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.EventInfo = *eventTemp
	this.msg.Desc ="get success"
	resp := GenEventResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//展示某个人的所有事件
func (this *EventController) List() {
	//这里虽然写的时eventid，但实际上时正则表达式匹配到的personid
	id:= this.Ctx.Input.Param(":event_id")
	id = id[1:]
	events,err := db.GetByPersonId(id)

	if err != nil {
		this.msg.Desc ="db query failed:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	if events == nil {
		this.msg.Desc ="this person doesn't have any events"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.EventList = events
	this.msg.Desc ="query success"
	resp := GenEventResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func GenEventResp(success bool,msg models.EventMessage) *models.EventResponse {
	var resResp models.EventResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}