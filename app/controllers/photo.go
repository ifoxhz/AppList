package controllers

import (
	"github.com/revel/revel"
	"github.com/ifoxhz/applist/app/models"
	//"github.com/ifoxhz/ifoxask/app/routes"
	//"strings"
)



type Photo struct {
		*revel.Controller
}

func (c Social) ShowPhoto() revel.Result {
	som := socailorm.SocialMod{}
	som.Init()
	list, err := som.GetAllApp()
	if err != nil {
		revel.WARN.Printf("don't find app list: %s", err)
	  return c.Redirect(App.Index)
	}

	return c.Render(list)
}
