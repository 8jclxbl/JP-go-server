package controllers

import (
	"github.com/astaxie/beego"
	"myproject/db"
		"io/ioutil"
	"encoding/json"
	"myproject/models"
	)

//登陆
type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
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
	//上面的错误处理同注册
	//从数据库中读入要登陆用户的信息
	user := jsReq.Params.User
	userTemp, err:= db.GetUser(user.UserName)

	//如果读取数据为空，说明该用户尚未注册
	if userTemp == nil {
		msg.Desc = "user had not registed"
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//处理获取用户信息时的数据库读取错误
	if err != nil {
		msg.Desc = "db err: " + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return

	} else {
		//密码验证
		if user.UserPass == userTemp.UserPass {
			//登陆状态调整
			state,loginState,err := db.LogIn(user.UserName)
			if loginState {
				//成功登陆
				msg.Desc = "sign in success"
			}	else {
				//登陆时数据库登陆状态调整出错
				if err != nil {
					msg.Desc=state + err.Error()
				}
				msg.Desc=state
			}
			resp := GenRespStruct(loginState,msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		} else {
			//密码错误
			msg.Desc = "password unmatched"
			resp := GenRespStruct(false,msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}
	}
	this.TplName = "login.html"
}
