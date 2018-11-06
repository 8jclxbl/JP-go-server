package controllers

import (
	"JP-go-server/db"
	"JP-go-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type EventController struct {
	beego.Controller
	jsReq models.JsonRequest
	msg  models.EventMessage
}

//根据读到的方法名选择相应的post方法
func (this *EventController) Post() {
	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	actionFromUrl := this.Ctx.Input.Param(":action")


	//读取出错时的回复
	if err != nil {
		this.msg.Desc = "报文读取错误: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &this.jsReq); err != nil {
		this.msg.Desc = "json解析错误: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	action := this.jsReq.Params.Action

	if action != actionFromUrl {
		this.msg.Desc = "json和url中的方法不同名"
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
		this.msg.Desc = "未知方法"
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
		this.msg.Desc = "未知方法"
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
		this.msg.Desc = "数据库错误: " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc = "事件创建成功"
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
	err := db.UpdateEvent(event)
	if err != nil {
		this.msg.Desc ="更新失败 " + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	this.msg.Desc = "更新成功"
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
		this.msg.Desc ="删除失败:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.Desc ="删除成功"
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
		this.msg.Desc ="事件不存在"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if err != nil {
		this.msg.Desc ="事件获取失败:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.EventInfo = *eventTemp
	this.msg.Desc ="事件获取成功"
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
		this.msg.Desc ="数据库查询失败:" + err.Error()
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	if events == nil {
		this.msg.Desc ="该人物尚无事件或人物不存在"
		resp := GenEventResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.EventList = events
	this.msg.Desc ="查询成功"
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