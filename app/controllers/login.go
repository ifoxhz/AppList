package controllers

import (
//	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"strings"
	"crypto/md5"
	"github.com/revel/revel"
//	"github.com/applist/app/routes"
)

type Login struct {
	       *revel.Controller
}

type PersonaResponse struct {
	Status   string `json:"status"`
	Email    string `json:"email"`
	Audience string `json:"audience"`
	Expires  int64  `json:"expires"`
	Issuer   string `json:"issuer"`
}

type LoginResult struct {
	StatusCode int
	Message    string
}

type LoginSession struct {
	Name    string
	Email     string
}

var LSession =make(map[string]LoginSession)

func (c Login)  Login( ) revel.Result {
	    if  val, ok := c.Session["access"]; ok  {
				  c.RenderArgs["email"] = LSession[val].Email
			}
	    return c.Render()
}

func (c Login)  Signin( ) revel.Result {
	name := c.Params.Get("UserName")
	email := c.Params.Get("UserEmail")
	if email != "ifoxhz@hotmail.com"{
		   delete(c.RenderArgs, "email")
		   return  c.Redirect(Login.Login)
	}

  session := md5.Sum([]byte(name+email))
	md := fmt.Sprintf("%x" , session)

	LSession[md] = LoginSession {Name: name, Email:email}
  //c.Session["email"] = email
	c.Session["access"] = md
	if val,ok :=c.Flash.Out["redirect"]; ok{
		   return  c.Redirect(val)
	}
	return c.Redirect("/")
}


func (c Login) Signup(assertion string) revel.Result {
	/*assertion = strings.TrimSpace(assertion)


	c.Session["email"] = p.Email
	fmt.Println("Login successful: ", p.Email)

	/*return &LoginResult{
		StatusCode: http.StatusOK,
		Message:    "Login successful.",
	}*/
	return  c.Redirect("/")

}

func (c Login) Signout() revel.Result {
	delete(c.Session, "email")
	delete(c.Session, "access")
	return c.Redirect("/")
}
