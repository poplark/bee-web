package v1

import (
	"encoding/json"
	"fmt"

	"bee-web/common"
	"bee-web/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NodeController struct {
	beego.Controller
}

type ListNodeReq struct {
	common.ListReq
}

type ListNodeResp struct {
	common.ListResp
}

func (c *NodeController) ListNode() {
	reqId := fmt.Sprintf("REQUESTUUID: %s", c.Ctx.Input.GetData("requestId").(string))
	clog := logs.GetLogger(reqId)

	var resp ListNodeResp
	resp.Action = "ListNodeResponse"
	resp.RetCode = 0

	offset, err := c.GetInt("Offset", 0)
	if err != nil {
		clog.Println("[ListNode] Offset error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
		c.Data["json"] = &resp
		c.ServeJSON()
		return
	}
	limit, err := c.GetInt("Limit", 10)
	if err != nil {
		clog.Println("[ListNode] Limit error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
		c.Data["json"] = &resp
		c.ServeJSON()
		return
	}

	var nodes []*model.Node
	qs := common.DB.QueryTable("node")
	if _, err := qs.OrderBy("-CreatedAt").Offset(offset).Limit(limit).All(&nodes); err != nil {
		clog.Println("[ListNode] Query error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
		c.Data["json"] = &resp
		c.ServeJSON()
		return
	}
	total, err := qs.Count()
	if err != nil {
		clog.Println("[ListNode] Count error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
		c.Data["json"] = &resp
		c.ServeJSON()
		return
	}

	resp.Offset = offset
	resp.Limit = limit
	resp.TotalCount = total
	resp.Data = &nodes

	c.Data["json"] = &resp
	c.ServeJSON()
}

type CreateNodeReq struct {
	common.BaseReq
	model.Node
}

type CreateNodeResp struct {
	common.BaseResp
}

func (c *NodeController) CreateNode() {
	reqId := fmt.Sprintf("REQUESTUUID: %s", c.Ctx.Input.GetData("requestId").(string))
	clog := logs.GetLogger(reqId)

	var req CreateNodeReq
	var resp CreateNodeResp
	resp.Action = "CreateNodeResponse"
	resp.RetCode = 0

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		clog.Println("[CreateNode] Unmarshal error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
		c.Data["json"] = &resp
		c.ServeJSON()
		return
	}

	var node = model.Node{
		Name: req.Name,
	}

	if t, err := common.DB.Insert(&node); err != nil {
		clog.Println("[CreateNode] Insert error: ", err)
		resp.RetCode = -1
		resp.Data = err.Error()
	} else {
		resp.Data = t
	}
	c.Data["json"] = &resp
	c.ServeJSON()
}
