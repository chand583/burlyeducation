package v1

import (
	"burlyeducation/controllers"
	"burlyeducation/models"
	"encoding/json"
	"fmt"
	"strings"
)

type UserController struct {
	controllers.AppController
}

// // URLMapping ...
// func (c *UserController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("GetOne", c.GetOne)
// 	c.Mapping("GetAll", c.GetAll)
// 	//c.Mapping("Put", c.Put)
// 	//c.Mapping("Delete", c.Delete)
// }

func (c *UserController) Get() {

	fmt.Println("in get function---")
	responseData := make(map[string]string)
	responseData["love1"] = "love to write api"
	c.Response(responseData, "Success")
}

func (c *UserController) Post() {
	var userModel models.User
	internalApi := c.Ctx.Request.RequestURI
	if !strings.Contains(internalApi, "burlyed") {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.ResponseWithError(404, "Page not found")
		return
	}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &userModel)
	upInsrt, err := models.UserModel{}.SaveUpdateUser(&userModel, userModel.Id)
	if err != nil {
		c.ResponseWithError(1001, err.Error())
	} else {
		c.Response(upInsrt, "Record insert/update successfully")
	}

	//fmt.Println("post", err)
}

func (c *UserController) GetAll() {
	responseData := make(map[string]string)
	responseData["love"] = "love to write api"
	c.Response(responseData, "Success")
}

// @router / [getOne]
func (c *UserController) GetOne() {
	fmt.Println("in get one function---")
	responseData := make(map[string]string)
	responseData["love12"] = "love to write api one api"
	c.Response(responseData, "Success")
}
