package controllers

type ErrorController struct {
    AppController
}

func (c *ErrorController) Error404() {
    c.ResponseWithError(404,"Page not found")
}

func (c *ErrorController) Error500() {
    c.ResponseWithError(500,"Something went wrong!")
}

func (c *ErrorController) Error503() {
    c.ResponseWithError(503,"Something went wrong!")
}

func (c *ErrorController) ErrorDb() {
    c.ResponseWithError(500,"Database is down!")
}
