package main

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/tekiomo-dark-side/goa-sample/app"
)

// PagesController implements the pages resource.
type PagesController struct {
	*goa.Controller
}

// NewPagesController creates a pages controller.
func NewPagesController(service *goa.Service) *PagesController {
	return &PagesController{Controller: service.NewController("PagesController")}
}

// Complete runs the complete action.
func (c *PagesController) Complete(ctx *app.CompletePagesContext) error {
	// PagesController_Complete: start_implement

	// Put your logic here

	// PagesController_Complete: end_implement
	return nil
}

// Home runs the home action.
func (c *PagesController) Home(ctx *app.HomePagesContext) error {
	data := map[string]interface{}{}
	res, err := getWithTemplate(data, "home.html")
	if err != nil {
		log.Errorf("index template error: %v", err)
		return ctx.InternalServerError()
	}
	return ctx.OK(res)
}
