package controllers

import "github.com/revel/revel"
import  "gopkg.in/redis.v3"
//import "github.com/"
//import "strings"

var GID int = -1
var RedisOption =  redis.Options {
    Addr: "127.0.0.1:6379",
    Password:"",
    DB: 0,
}

var Rdb  * redis.Client = nil


type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
     if (Rdb == nil){
			   Rdb =  redis.NewClient(&RedisOption)
		}

   if err := Rdb.Incr("IOTHILL.PAGEVIEW").Err(); err != nil {
    revel.INFO.Printf("failed to conneted redis ", err)
   }
	//local := c.RenderArgs["controllerCurrentLocale"]
	//title := revel.Message(c.Request.Locale  , "website_title")
	//revel.INFO.Printf("local %s  : %s  ", revel.MessageLanguages(), title)
	return c.Render()
}

func (c App) WaitInfo() revel.Result {

	return c.Render()
}

func (c App) SendMessage(UserName string) revel.Result {
	//name := c.Params.Get("UserName")
	//email := c.Params.Get("UserEmail")
  c.Flash.Success("Thanks")
	//title := revel.Message(c.Request.Locale  , "website_title")
	//revel.INFO.Printf("input: %s %s ", name,email)
 // c.RenderText("Post success")
	return c.Redirect("index.html")
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
