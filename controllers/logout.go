package controllers

import (
	"github.com/astaxie/beego"
	"myproject/models"
	"encoding/json"
	"myproject/db"
	"io/ioutil"
)

//登出，逻辑与登陆类似
type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.TplName = "login.html"
}

func (this *LogoutController) Post() {
	var jsReq models.JsonRequest
	var msg models.Message

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
	user := jsReq.Params.User

	userName := user.UserName
	userTemp, err:= db.GetUser(userName)
	if userTemp == nil {
		msg.Desc = "user had not registed"
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	state,logOutState,err := db.LogOut(userName)
	if logOutState {
		msg.Desc = "logout successd"
	}	else {
		if err != nil {
			msg.Desc = state + err.Error()
		} else {
			msg.Desc = state
		}
	}
	resp := GenRespStruct(logOutState,msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return

	this.TplName = "login.html"
}