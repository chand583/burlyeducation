package v1

import (
	"burlyeducation/controllers"
	"fmt"
)

type UserController struct {
	controllers.AppController
}

// Url Mapping
func (c *UserController) URLMapping() {
	fmt.Println("mapping  "+ c.Ctx.Request.Method)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Get", c.GetAll)

}

// @router / [get]
func (c *UserController) Get() {
	
	fmt.Println("in get function---")
	responseData := make(map[string]string)
	responseData["love1"] = "love to write api"
	c.Response(responseData, "Success")
}

func (c *UserController) Post() {
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println("in post function---")
	responseData := make(map[string]string)
	responseData["post"] = "love to post	 api"
	c.Response(responseData, "Success")
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
