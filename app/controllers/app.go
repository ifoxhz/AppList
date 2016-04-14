package controllers

import "github.com/revel/revel"
//import "encoding/json"
//import "strings"

var GID int = 1
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//local := c.RenderArgs["controllerCurrentLocale"]
	//title := revel.Message(c.Request.Locale  , "website_title")
	//revel.INFO.Printf("local %s  : %s  ", revel.MessageLanguages(), title)
	return c.Render()
}

func (c App) GetDog() revel.Result {

	type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}
GID ++
GID = GID % 1000
group := ColorGroup{
	ID:     GID,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}
//  lang, _ := json.Marshal(group)
//c.Response.Status = 301
//c.Response.ContentType = "application/json"
	return c.RenderJsonP(c.Params.Get("callback"), group)
}
