package controllers

import (
	"github.com/astaxie/beego"
	"myproject/models"
	"myproject/db"
	"encoding/json"
	"io/ioutil"
	)
//注册
type RegController struct {
	beego.Controller
}


func (this *RegController) Get() {
	this.TplName="test.tpl"
}

func (this *RegController) Post() {
	this.TplName="test.tpl"
	var jsReq models.JsonRequest
	var msg  models.Message

	//从request中读取json数据
	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	//读取出错时的回复
	if err != nil {
		msg.Desc = "body read err: " + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &jsReq); err != nil {
		msg.Desc = "json parse read err: " + err.Error()
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	user := jsReq.Params.User

	//检查数据库中该用户名是否已被占用
	if userTemp,_ := db.GetUser(user.UserName); userTemp != nil {
		//fmt.Println(userTemp)
		msg.Desc = "user exists"
		resp := GenRespStruct(false,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	} else {
		//未被占用将注册信息写入数据库
		err = db.CreatUser(user)
		//写入时出错的回复
		if err != nil {
			msg.Desc = "db err: " + err.Error()
			resp := GenRespStruct(false,msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}
		//注册成功时回复数据库生成的用户id
		userId, _ := db.GetUserID(user.UserName)
		msg.Userid = userId
		resp := GenRespStruct(true,msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

//生成要写进response的json所来源的结构体
//success：bool, 指示响应成功或失败
//msg：Message struct,回复的消息字段
func GenRespStruct(success bool,msg models.Message) *models.BaseResponse {
	var resResp models.BaseResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}
