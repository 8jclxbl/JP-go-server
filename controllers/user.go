package controllers

import (
	"github.com/astaxie/beego"
	"JP-go-server/models"
	"JP-go-server/db"
	"JP-go-server/util"
	"io/ioutil"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Reg() {
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

func (this *UserController) Update() {
	this.TplName ="test.tpl"

	var jsReq models.JsonRequest
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

func (this *UserController) Login() {
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
		if util.Cipher(user.UserPass) == userTemp.UserPass {
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

func (this *UserController) Logout() {
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


//生成要写进response的json所来源的结构体
//success：bool, 指示响应成功或失败
//msg：Message struct,回复的消息字段
func GenRespStruct(success bool,msg models.Message) *models.BaseResponse {
	var resResp models.BaseResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}
