package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	revel.WARN.Printf("app path: %s %s", revel.AppPath, revel.BasePath)
	return c.Render()
}
