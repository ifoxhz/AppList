package controllers

import (
	"github.com/revel/revel"
	"github.com/applist/app/models"
	"github.com/applist/app/routes"
	//"fmt"
)



type Photo struct {
		*revel.Controller
}

func (c Photo) ShowPhoto() revel.Result {
	som := model.PhotoMod{}
	som.Init()
	list, err := som.GetAllApp()
	if err != nil {
		revel.WARN.Printf("don't find app list: %s", err)
	  return c.Redirect(App.Index)
	}
	//fmt.Println(routes.App.Index())
	c.RenderArgs["indexpage"] = routes.App.Index()
	return c.Render(list)
}
