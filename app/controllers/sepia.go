package controllers

//import "github.com/boombuler/barcode"
import "github.com/PuerkitoBio/goquery"

import (
	"fmt"
  "gopkg.in/redis.v3"
	"net/url"
	"os"
	"bytes"
	"io"
	"io/ioutil"
	HTTP "net/http"
	"path/filepath"
	"reflect"
  "encoding/json"
	//"github.com/applist/app/models"
	AM "github.com/applist/lib/appmodel"
	"github.com/revel/revel"
	"github.com/applist/app/routes"
	"strconv"
	// "strings"
	"crypto/md5"
)

const CONTENT_DIV_CLASS_STRING string = "lockup product application"
const CONTENT_IMG_SRC_STRING string = "src-swap-high-dpi"

type Sepia struct {
	       *revel.Controller
}

func (c Sepia) ShowSepia() revel.Result {
	      c.RenderArgs["indexpage"] = routes.App.Index()
	      return c.Render()
}

func (c Sepia) SubmitApp() revel.Result {
	name := c.Params.Get("app_name")
	url := c.Params.Get("app_url")
	cat := c.Params.Get("app_cat")
	catid,_ := strconv.Atoi(cat)
  c.RenderArgs["indexpage"] = routes.App.Index()
	revel.ERROR.Printf("temp: %s", routes.Sepia.ShowSepia())


		if  name == "" || url == "" {
      revel.ERROR.Printf("name or url is wrong")
			c.Flash.Error("提交的错误的url")
			return c.Redirect(Sepia.ShowSepia)
	}

	val,ok := AM.ModeKeyMap[AM.Category(catid)];
	if  !ok{
      revel.ERROR.Printf("get wrong category: %s", cat)
			c.Flash.Error("提交的数据有误")
				return c.Redirect(Sepia.ShowSepia)
	}
	revel.INFO.Printf("submit: %s %s %s", name,url,val)
	var app AM.AppModel
	app.Id = fmt.Sprintf("%x" , md5.Sum( []byte(name) ))
	app.Name = name
	app.Url = url
	app.Category = AM.Category(catid)

  err := ExtractIcon(&app)
	if err != nil{
		c.Flash.Error("创建App信息失败")
		return c.Redirect(Sepia.ShowSepia)
	}

  err = Save(app)
	if err != nil{
		c.Flash.Error("创建App信息失败")
	}else{
		  c.Flash.Success("submit is success!")
	}
	return c.Redirect(Sepia.ShowSepia)
}

func DownAvatar(tg url.URL, sl string) error {
	revel.INFO.Println(sl)
	out, err := os.Create(sl)
	if err != nil {
    fmt.Println("---",err)
		return err
	}
	defer out.Close()
	resp, err := HTTP.Get(tg.String())
	if err != nil {
		revel.INFO.Println(err)
		return err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))
	return err
}

func ExtractIcon(ap * AM.AppModel) error {
  isfind := false
	doc, err := goquery.NewDocument(ap.Url)
	if err != nil {
		revel.ERROR.Printf("%s", err)
		return err
	}
	doc.Find("div").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if attr, is := s.Attr("class"); is == false || attr != CONTENT_DIV_CLASS_STRING {
			//fmt.Printf("class %d: %s-%s\n", i, attr, parent.Text())
			return true
		}

		s.Find("img").EachWithBreak(func(i int, s *goquery.Selection) bool {
			src, is := s.Attr(CONTENT_IMG_SRC_STRING)
			if !is {
				return true
			}
			url, _ := url.Parse(src)
			if url.IsAbs() {
				revel.INFO.Printf("abs url: %s - \n",  url)
			} else {
				x, _ := url.Parse(ap.Url)
				url.Host = x.Host
				url.Scheme = x.Scheme
					revel.INFO.Printf("not abs url: %s - \n",  url)
			}
		  var	iconlocation = AM.GetIconLocation(ap.Category)
			_, file := filepath.Split(url.Path)
			err = DownAvatar(*url,  revel.BasePath + "/" +iconlocation + ap.Id + file) //down icon
      if err != nil {
				return false
			}
		  ap.Icon =  iconlocation + ap.Id + file  //keep icon location
      isfind = true
			return false
		})
    return !isfind
	}) //each over
	if isfind {
       return  nil
  }else {
      return err
  }
}

func debugObj(obj interface{}) {
	ov := reflect.ValueOf(obj)
	var vref reflect.Value
	if ov.Kind() != reflect.Struct && ov.Kind() != reflect.Interface && ov.Kind() != reflect.Ptr {
		revel.INFO.Printf("it is not struct or interface , is %s \n", ov.Kind())
		return
	}
	revel.INFO.Printf("Obj: %s, Type:%s \n", ov.Type(), ov.Kind())
	if ov.Kind() == reflect.Struct {
		vref = ov
	} else {
		vref = ov.Elem()
	}
	typeOfType := vref.Type()
	for i := 0; i < vref.NumField(); i++ {
		field := vref.Field(i)
		revel.INFO.Printf("%d. %s %s  method: %v \n", i, typeOfType.Field(i).Name, field.Type(), field.String())
	}

	revel.INFO.Printf("method: %d  \n", typeOfType.NumMethod())
	for i := 0; i < typeOfType.NumMethod(); i++ {
		mt := typeOfType.Method(i)
		revel.INFO.Printf("%d  method: %v \n", i, mt)
	}

}

func Save(app AM.AppModel) error {
  jsob, err := json.Marshal(app)
  if err != nil {
       revel.ERROR.Printf("error:", err)
       return err
  }
  rdc := redis.NewClient(&RedisOption)

  if err := rdc.HSet(AM.ModeKeyMap[app.Category], app.Id, string(jsob)).Err(); err != nil {
          revel.ERROR.Print("error:", err)
					return err
  }
  return nil
}
