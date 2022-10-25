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

func (c *UserController) GetAll() {
fmt.Println(" Get all`````````````````")
}

func (c *UserController) GetOne() {
	fmt.Println("GETone")
}
