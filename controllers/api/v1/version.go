package v1

import (
	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

func (c *VersionController) Get() {
	c.Data["json"] = "Version is 1."
	c.ServeJSON(true)
}
