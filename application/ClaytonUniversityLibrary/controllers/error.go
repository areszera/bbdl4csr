package controllers

import beego "github.com/beego/beego/v2/server/web"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Layout = "frame.gohtml"
	c.TplName = "error.gohtml"
	c.Data["statusCode"] = 401
	c.Data["info"] = "Unauthorized"
}

func (c *ErrorController) Error403() {
	c.Layout = "frame.gohtml"
	c.TplName = "error.gohtml"
	c.Data["statusCode"] = 403
	c.Data["info"] = "Forbidden"
}

func (c *ErrorController) Error404() {
	c.Layout = "frame.gohtml"
	c.TplName = "error.gohtml"
	c.Data["statusCode"] = 404
	c.Data["info"] = "Not Found"
}

func (c *ErrorController) Error500() {
	c.Layout = "frame.gohtml"
	c.TplName = "error.gohtml"
	c.Data["statusCode"] = 500
	c.Data["info"] = "Internal Server Error"
}

func (c *ErrorController) Error503() {
	c.Layout = "frame.gohtml"
	c.TplName = "error.gohtml"
	c.Data["statusCode"] = 503
	c.Data["info"] = "Service Unavailable"
}
