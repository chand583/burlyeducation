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
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)

}
// @router / [get]
func (c *UserController) Get() {
	fmt.Println("in get function---")
	responseData:=make(map[string]string)
	responseData["love1"]="love to write api"
	c.Response(responseData, "Success")
}
func (c *UserController) GetAll() {

}
// @router / [get]
func (c *UserController) GetOne() {
	fmt.Println("in get one function---")
	responseData:=make(map[string]string)
	responseData["love1"]="love to write api one api"
	c.Response(responseData, "Success")
}
