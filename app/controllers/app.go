package controllers

import "github.com/revel/revel"
//import "strings"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//local := c.RenderArgs["controllerCurrentLocale"]
	title := revel.Message(c.Request.Locale  , "website_title")
	revel.INFO.Printf("local %s  : %s  ", revel.MessageLanguages(), title)
	return c.Render()
}
