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
//根据读取的字段选择相应的post方法
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
	case "list":
		this.List()
	default:
		this.msg.Desc = "Unknown method"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}
//根据读取的字段选择相应的get方法
func (this *PersonController) Get() {
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

//创建人物对象
func (this *PersonController) Create() {
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

//更新人物对像
func (this *PersonController) Update() {
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

//根据条件列表展示人物对象
func (this *PersonController) List() {
	condition := this.jsReq.Params.PersonSelect
	//一个输入的合法性检验，该条件忽略时由于数据查询后需要根据页数上限得到条目上限进行截取，如果不设置默认为0，会导致条目上线为0，而截取时条目数目小于上限，导致slice越界
	if condition.ConPageNum == 0 {
		condition.ConPageNum = 1
	}

	//获取满足条件的人物对象
	people, err := db.ListPerson(condition)
	if err != nil {
		this.msg.Desc ="db query failed:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	if people == nil {
		this.msg.Desc ="this person doesn't have any events"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.PeopleList = people
	this.msg.Desc ="query success"
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

//删除人物对象
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

//展示指定id的人物对象
func (this *PersonController) View() {
	id:= this.Ctx.Input.Param(":person_id")
	id = id[1:]
	person, err := db.GetPersonById(id)

	if person == nil {
		this.msg.Desc =" person not exists"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

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


func GenPersonResp(success bool,msg models.PersonMessage) *models.PersonResponse {
	var resResp models.PersonResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}