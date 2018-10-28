package controllers

import (
	"JP-go-server/db"
	"JP-go-server/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	beego.Controller
	jsReq models.JsonRequest
	msg  models.FileMessage
}

//读取文件的html，测试用
func (this *FileController) Get() {
	this.TplName = "upload.html"
}

//上传文件的函数，文件url为/upload/+时间戳再加上文件类型
func (this *FileController) Upload() {
	//从前端的uploadfile的input中读取
	file,fileHEader,err := this.GetFile("uploadfile")
	if err != nil {
		this.msg.Desc ="upload error:" + err.Error()
		resp := GenFileResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	//获取文件类型
	fileName := fileHEader.Filename
	fileInfo := strings.Split(fileName,".")
	fileType := fileInfo[len(fileInfo)-1]

	now := time.Now()
	timeStamp := now.Unix()
	stamp := strconv.Itoa(int(timeStamp))

	path := "./upload/" + stamp + "." + fileType
	this.SaveToFile("uploadfile",path)
	defer file.Close()

	fileTemp := models.File{
		FileUrl:path,
		FileType:fileType,
	}
	//将文件相关信息写入数据库
	_, err = db.CreatFile(fileTemp,"")
	if err != nil {
		this.msg.Desc ="db save error:" + err.Error()
		resp := GenFileResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	this.msg.FileUrl = path
	this.msg.Desc ="upload success"
	resp := GenFileResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

//删除给定url的文件
func (this *FileController) Delete() {
	url:= this.Ctx.Input.Param(":file_url")

	err := db.DeleteFile(url)
	if err != nil {
		this.msg.Desc ="delete failed: " + err.Error()
		resp := GenFileResp(false,this.msg)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	this.msg.Desc ="delete success"
	resp := GenFileResp(true,this.msg)
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func GenFileResp(success bool,msg models.FileMessage) *models.FileResponse {
	var resResp models.FileResponse
	resResp.Success = success
	resResp.Msg = msg
	return  &resResp
}