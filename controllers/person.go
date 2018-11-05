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
		this.msg.Desc = "报文读取错误: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &this.jsReq); err != nil {
		this.msg.Desc = "json解析错误: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	action := this.jsReq.Params.Action

	if action != actionFromUrl {
		this.msg.Desc = "json中的方法和url不同名: "
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
		this.msg.Desc = "未知方法"
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
		this.msg.Desc = "未知方法"
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
		this.msg.Desc = "数据库错误: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc = "人物创建成功"
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
		this.msg.Desc ="更新失败: " + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	this.msg.Desc = "更新成功"
	resp := GenPersonResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//根据条件列表展示人物对象
func (this *PersonController) List() {
	condition := this.jsReq.Params.PersonSelect
	//一个输入的合法性检验，该条件忽略时由于数据查询后需要根据页数上限得到条目上限进行截取，如果不设置默认为0，会导致条目上线为0，而截取时条目数目小于上限，导致slice越界
	/*
	if condition.ConPageNum == 0 {
		condition.ConPageNum = 1
	}
	*/
	if condition.ConPageSize == 0 {
		condition.ConPageSize = 20
	}

	//获取满足条件的人物对象
	people, err := db.ListPerson(condition)
	if err != nil {
		this.msg.Desc ="数据库查询失败:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	if people == nil {
		this.msg.Desc ="没有符合要求的人物"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.PeopleList = people
	this.msg.Desc ="查询成功"
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
		this.msg.Desc ="删除失败:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.Desc ="删除成功"
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
		this.msg.Desc ="人物不存在"
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if err != nil {
		this.msg.Desc ="数据库查询出错:" + err.Error()
		resp := GenPersonResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.Desc ="获取成功"
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