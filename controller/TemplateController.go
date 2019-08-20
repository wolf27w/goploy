package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zhenorzz/goploy/core"
	"github.com/zhenorzz/goploy/model"
)

// Template struct
type Template Controller

// GetList template list
func (template Template) GetList(w http.ResponseWriter, gp *core.Goploy) {
	type RepData struct {
		Template model.Templates `json:"templateList"`
	}

	templateList, err := model.Template{}.GetList()
	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	response := core.Response{Data: RepData{Template: templateList}}
	response.JSON(w)
}

// GetOption template list
func (template Template) GetOption(w http.ResponseWriter, gp *core.Goploy) {
	type RepData struct {
		Template model.Templates `json:"templateList"`
	}

	templateList, err := model.Template{}.GetAll()
	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	response := core.Response{Data: RepData{Template: templateList}}
	response.JSON(w)
}

// Add one template
func (template Template) Add(w http.ResponseWriter, gp *core.Goploy) {
	type ReqData struct {
		Name    string `json:"name"`
		Remark  string `json:"remark"`
		Package string `json:"package"`
		Script  string `json:"scrpit"`
	}
	var reqData ReqData
	err := json.Unmarshal(gp.Body, &reqData)
	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	id, err := model.Template{
		Name:       reqData.Name,
		Remark:     reqData.Remark,
		Package:    reqData.Package,
		Script:     reqData.Script,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}.AddRow()

	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}

	for _, fileName := range strings.Split(reqData.Package, ",") {
		err := os.Rename(core.GolbalPath+"tmp/"+fileName, core.GolbalPath+"package/"+reqData.Name+":"+strconv.Itoa(int(id))+"/"+fileName)
		if err != nil {
			response := core.Response{Code: 1, Message: err.Error()}
			response.JSON(w)
			return
		}
	}

	response := core.Response{Message: "添加成功"}
	response.JSON(w)
}

// Edit one template
func (template Template) Edit(w http.ResponseWriter, gp *core.Goploy) {
	type ReqData struct {
		ID      uint32 `json:"id"`
		Name    string `json:"name"`
		Remark  string `json:"remark"`
		Package string `json:"package"`
		Script  string `json:"script"`
	}
	var reqData ReqData
	err := json.Unmarshal(gp.Body, &reqData)
	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	err = model.Template{
		ID:         reqData.ID,
		Name:       reqData.Name,
		Remark:     reqData.Remark,
		Package:    reqData.Package,
		Script:     reqData.Script,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}.EditRow()

	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	response := core.Response{Message: "修改成功"}
	response.JSON(w)
}

// Remove one Template
func (template Template) Remove(w http.ResponseWriter, gp *core.Goploy) {
	type ReqData struct {
		ID uint32 `json:"id"`
	}
	var reqData ReqData
	err := json.Unmarshal(gp.Body, &reqData)
	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	err = model.Template{
		ID:         reqData.ID,
		UpdateTime: time.Now().Unix(),
	}.Remove()

	if err != nil {
		response := core.Response{Code: 1, Message: err.Error()}
		response.JSON(w)
		return
	}
	response := core.Response{Message: "删除成功"}
	response.JSON(w)
}