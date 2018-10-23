package controllers

import (
	"github.com/astaxie/beego"
	"JP-go-server/models"
	"io/ioutil"
	"encoding/json"
	"JP-go-server/db"
)

//用户数据更新
type UpdateController struct {
	beego.Controller
}

func (this *UpdateController) Post() {
	this.TplName ="test.tpl"

	var jsReq 	models.JsonRequest
	var msg 	models.Message

	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	if err != nil {
		msg.Desc = "body read err: " + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if err = json.Unmarshal(jsonTemp, &jsReq); err != nil {
		msg.Desc = "json parse read err: " + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//上面的错误处理与注册一致
	user := jsReq.Params.User
	//更新数据库的数据
	err = db.UpdateUser(user)
	if err != nil {
		msg.Desc ="update failed" + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	msg.Desc = "updated"
	resp := GenRespStruct(true,msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}
