package controllers

import (
	"JP-go-server/db"
	"JP-go-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type PersonController struct {
	beego.Controller
	jsReq models.JsonRequest
	msg  models.PersonMessage
}

func (this *PersonController) Post() {
	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	actionFromUrl := this.Ctx.Input.Param(":action")


	//读取出错时的回复
	if err != nil {
		this.msg.Desc = "body read err: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &this.jsReq); err != nil {
		this.msg.Desc = "json parse read err: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	action := this.jsReq.Params.Action

	if action != actionFromUrl {
		this.msg.Desc = "actions from url and json are not the same: "
		resp := GenPersonResp(false,this.msg)
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
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

func (this *PersonController) Get() {
	this.TplName="test.tpl"
	action:= this.Ctx.Input.Param(":action")

	switch action {
	case "delete":
		this.Delete()
	case "view":
		this.View()

	default:
		this.msg.Desc = "Unknown method"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

func (this *PersonController) Create() {
	this.TplName="test.tpl"
	person := this.jsReq.Params.Person
	personID,err := db.CreatPerson(person)
	//写入时出错的回复
	if err != nil {
		this.msg.Desc = "db err: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc = "person created success"
	this.msg.PersonInfo.PersonId = personID
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return


}


func (this *PersonController) Update() {
	this.TplName="test.tpl"
	person := this.jsReq.Params.Person

	err := db.UpdatePerson(person)
	if err != nil {
		this.msg.Desc ="update failed: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	this.msg.Desc = "updated"
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func (this *PersonController) Delete() {
	id:= this.Ctx.Input.Param(":person_id")
	id = id[1:]
	err := db.DeletePerson(id)
	if err != nil {
		this.msg.Desc ="delete failed:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.Desc ="delete success"
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *PersonController) View() {
	id:= this.Ctx.Input.Param(":person_id")
	id = id[1:]
	person, err := db.GetPersonById(id)

	if err != nil {
		this.msg.Desc =" get object failed:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc ="get success"
	this.msg.PersonInfo = *person
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}
/*
func (this *PersonController) List() {

}
*/

func GenPersonResp(success bool,msg models.PersonMessage) *models.PersonResponse {
	var resResp models.PersonResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}