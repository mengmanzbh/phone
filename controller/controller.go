package controller

type Controller struct {
}

type Message struct {
	Message string `json:"message" example:"message"`
}

func NewController() *Controller {
	return &Controller{}
}
