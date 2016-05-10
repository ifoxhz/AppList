package controllers

import (
//	"golang.org/x/crypto/bcrypt"
	//"fmt"
	"github.com/revel/revel"
	"github.com/applist/app/models"
	//"github.com/ifoxhz/ifoxask/app/routes"
	//"strings"
)

type Social struct {
		*revel.Controller
}

func (c Social) ShowSocial() revel.Result {
	som := model.SocialMod{}
	som.Init()
	list, err := som.GetAllApp()
	if err != nil {
		revel.WARN.Printf("don't find app list: %s", err)
	  return c.Redirect(App.Index)
	}

	return c.Render(list)
}
