package main

//import "github.com/boombuler/barcode"
import "github.com/PuerkitoBio/goquery"
import (
	"fmt"
  "gopkg.in/redis.v3"
	//"log"
	"net/url"
	"os"
	//"image/png"
	"bytes"
	"io"
	"io/ioutil"
	HTTP "net/http"
	"path/filepath"
	"reflect"
  "encoding/json"
   "strconv"
	// "strings"
	"crypto/md5"
)

type AppModel struct {
  Id         string
	Name   string
	Url        string
	Icon      string
	QrcodeLoc string
  Tag           string
  Category   int
}


var RedisOption =  redis.Options {
    Addr: "127.0.0.1:6379",
    Password:"",
    DB: 0,
}

const CONTENT_DIV_CLASS_STRING string = "lockup product application"
const CONTENT_IMG_SRC_STRING string = "src-swap-high-dpi"

const iconlocation string ="/home/yong/go/src/github.com/applist/public/photo/images/icon/"

type Category int

const  (
    social Category = iota +100
    photo
)



func DownAvatar(tg url.URL, sl string) error {
	fmt.Println(sl)
	out, err := os.Create(sl)
	if err != nil {
    fmt.Println("---",err)
		return err
	}
	defer out.Close()
	resp, err := HTTP.Get(tg.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))
	return err
}

func ExtractIcon(ap *AppModel) error {
  isfind := false
	doc, err := goquery.NewDocument(ap.Url)
	if err != nil {
		fmt.Println(err)
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
				fmt.Printf("abs %d: %s - \n", i, url)
			} else {
				x, _ := url.Parse(ap.Url)
				url.Host = x.Host
				url.Scheme = x.Scheme
				fmt.Printf("not abs %d: %s - \n", i, url)
			}
			_, file := filepath.Split(url.Path)
			err = DownAvatar(*url, iconlocation + ap.Id + file) //down icon
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
		fmt.Printf("it is not struct or interface , is %s \n", ov.Kind())
		return
	}
	fmt.Printf("Obj: %s, Type:%s \n", ov.Type(), ov.Kind())
	if ov.Kind() == reflect.Struct {
		vref = ov
	} else {
		vref = ov.Elem()
	}
	typeOfType := vref.Type()
	for i := 0; i < vref.NumField(); i++ {
		field := vref.Field(i)
		fmt.Printf("%d. %s %s  method: %v \n", i, typeOfType.Field(i).Name, field.Type(), field.String())
	}

	fmt.Printf("method: %d  \n", typeOfType.NumMethod())
	for i := 0; i < typeOfType.NumMethod(); i++ {
		mt := typeOfType.Method(i)
		fmt.Printf("%d  method: %v \n", i, mt)
	}

}

const  APP_MODE_PHOTO_KEY = "APP:PHOTO"

func Save(app AppModel) error {
  jsob, err := json.Marshal(app)
  if err != nil {
       fmt.Println("error:", err)
       return err
  }
  rdc := redis.NewClient(&RedisOption)

  if err := rdc.HSet(APP_MODE_PHOTO_KEY, app.Id, string(jsob)).Err(); err != nil {
         panic(err)
  }
  return nil
}





func main() {

   cat ,_ := strconv.Atoi(os.Args[2])
  //fmt.Println("%s", os.Args[1])
	//ul := `https://itunes.apple.com/cn/app/id416048305?mt=8`
	app := AppModel{
		Name: os.Args[1],
		Url:  os.Args[3],
    Category: cat,
	}

  app.Id = fmt.Sprintf("%x" , md5.Sum( []byte(app.Name) ))
  fmt.Println(app, social)
	err := ExtractIcon(&app)
  if err != nil {
    panic(err)
    return
  }
  Save(app)
}
