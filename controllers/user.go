package controllers

import (
	"JP-go-server/db"
	"JP-go-server/models"
	"JP-go-server/util"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type UserController struct {
	beego.Controller
	jsReq models.JsonRequest
	msg  models.UserMessage
}

func (this *UserController) Post() {
	requestBody := this.Ctx.Request.Body
	jsonTemp, err := ioutil.ReadAll(requestBody)
	actionFromUrl := this.Ctx.Input.Param(":action")

	//读取出错时的回复
	if err != nil {
		this.msg.Desc = "报文读取错误: " + err.Error()
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//json解析并在有错时的回复
	if err = json.Unmarshal(jsonTemp, &this.jsReq); err != nil {
		this.msg.Desc = "json解析错误: " + err.Error()
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	action := this.jsReq.Params.Action

	if action != actionFromUrl {
		this.msg.Desc = "actions from url and json are not the same"
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	switch action {
	case "reg":
		this.Reg()
	case "login":
		this.Login()
	case "logout":
		this.Logout()
	case "update":
		this.Update()
	default:
		this.msg.Desc = "Unknown method"
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

func (this *UserController) Reg() {
	//从request中读取json数据
	user := this.jsReq.Params.User
	//检查数据库中该用户名是否已被占用
	if userTemp,_ := db.GetUser(user.UserName); userTemp != nil {
		//fmt.Println(userTemp)
		this.msg.Desc = "用户名已被占用"
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	} else {
		//未被占用将注册信息写入数据库
		userId,err := db.CreatUser(user)
		//写入时出错的回复
		if err != nil {
			this.msg.Desc = "数据库错误: " + err.Error()
			resp := GenUserResp(false,this.msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}
		//注册成功时回复数据库生成的用户id
		this.msg.Userid = userId
		resp := GenUserResp(true,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
}

func (this *UserController) Update() {
	//上面的错误处理与注册一致
	user := this.jsReq.Params.User
	//更新数据库的数据
	err := db.UpdateUser(user)
	if err != nil {
		this.msg.Desc ="更新失败" + err.Error()
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//更新成功
	this.msg.Desc = "更新成功"
	resp := GenUserResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func (this *UserController) Login() {

	//上面的错误处理同注册
	//从数据库中读入要登陆用户的信息
	user := this.jsReq.Params.User
	userTemp, err:= db.GetUser(user.UserName)

	//如果读取数据为空，说明该用户尚未注册
	if userTemp == nil {
		this.msg.Desc = "用户未注册"
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	//处理获取用户信息时的数据库读取错误
	if err != nil {
		this.msg.Desc = "db err: " + err.Error()
		resp := GenUserResp(false,this.msg)
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
				this.msg.Desc = "登陆成功"
				this.msg.Userid = userTemp.UserId
			}	else {
				//登陆时数据库登陆状态调整出错
				if err != nil {
					this.msg.Desc=state + err.Error()
				}
				this.msg.Desc=state
			}
			resp := GenUserResp(loginState,this.msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		} else {
			//密码错误
			this.msg.Desc = "密码错误"
			resp := GenUserResp(false,this.msg)
			this.Data["json"] = resp
			this.ServeJSON()
			return
		}
	}
	this.TplName = "login.html"
}

func (this *UserController) Logout() {

	user := this.jsReq.Params.User

	userName := user.UserName
	userTemp, err:= db.GetUser(userName)
	if userTemp == nil {
		this.msg.Desc = "用户未注册"
		resp := GenUserResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	state,logOutState,err := db.LogOut(userName)
	if logOutState {
		this.msg.Desc = "用户已成功下线"
	}	else {
		if err != nil {
			this.msg.Desc = state + err.Error()
		} else {
			this.msg.Desc = state
		}
	}
	resp := GenUserResp(logOutState,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return

	this.TplName = "login.html"
}


//生成要写进response的json所来源的结构体
//success：bool, 指示响应成功或失败
//msg：Message struct,回复的消息字段
func GenUserResp(success bool,msg models.UserMessage) *models.UserResponse {
	var resResp models.UserResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}
