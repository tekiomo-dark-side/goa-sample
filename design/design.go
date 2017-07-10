package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("page", func() {
	Title("Page")
	Description("page")
	Host("localhost:8081")
	Scheme("http")
})

var _ = Resource("pages", func() {
	Action("home", func() {
		Routing(GET("/"))
		Description("トップページ")
		Response(OK, "text/html; charset=utf-8")
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("complete", func() {
		Routing(GET("/complete"))
		Description("完了ページ")
		Response(OK, "text/html; charset=utf-8")
		Response(NotFound)
		Response(InternalServerError)
	})
})
