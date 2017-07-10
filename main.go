//go:generate goagen bootstrap -d github.com/tekiomo-dark-side/goa-sample/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tekiomo-dark-side/goa-sample/app"
)

func main() {
	// Create service
	service := goa.New("page")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "pages" controller
	c := NewPagesController(service)
	app.MountPagesController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
