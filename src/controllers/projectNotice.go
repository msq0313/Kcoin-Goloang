package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectNoticeController struct {
	beego.Controller
}

func (c *ProjectNoticeController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "projectNotice.html"		//该controller对应的页面
}