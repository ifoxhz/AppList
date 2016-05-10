package model
import (
"sort"
 "errors"
 //"crypto/md5"
 "github.com/revel/revel"
 "github.com/applist/lib/oslib"
 "github.com/applist/lib/qrcode"
 "github.com/applist/lib/appmodel"
 "gopkg.in/redis.v3"
 "encoding/json"
 	"path/filepath"
)

var RedisOption =  redis.Options {
    Addr: "127.0.0.1:6379",
    Password:"",
    DB: 0,
}

const  APP_MODE_PHOTO_KEY = "APP:PHOTO"

var Rdb  * redis.Client = nil

type PhotoMod struct {
      AppList  map[string]   appmodel.AppModel
}
const  PhotoQrLocation = "public/photo/images/qrimg/"
const  PhotoIcon = "public/photo/images/icon/"

func  (s  *PhotoMod) Load()  (map[string]  appmodel.AppModel, error) {
  if (Rdb == nil){
      Rdb =  redis.NewClient(&RedisOption)
  }
 	result := Rdb.HGetAllMap(APP_MODE_PHOTO_KEY)
  if err := result.Err(); err != nil {
        revel.ERROR.Printf("Failed to get app from db  : %s", err)
 				return  nil,err
   }
 	rmap, err := result.Result()
 	if err != nil {
 		    revel.ERROR.Printf("Result is wrong  : %s", err)
 				return nil,err
 	}
 	//fmt.Println(result)
  var app  appmodel.AppModel
 	applist   :=make(map[string]  appmodel.AppModel)
 	for key,value := range rmap {
 		err := json.Unmarshal( []byte(value), &app)
 	  if err != nil {
 	      revel.ERROR.Printf("unmarshal error:%s, %s", err, value)
 	       return nil, err
 	  }
     applist[key] = app
 	}
  return applist,nil
}

func  (s  *PhotoMod) Init()   {
     if  s.AppList != nil {
       return
    }
    slist, _:= s.Load()

    s.AppList = make(map[string]  appmodel.AppModel)
    for k, v := range slist {
          app := v
          _, file := filepath.Split(app.Icon)
          qrfn :=  revel.BasePath + "/" + PhotoQrLocation + file
        //  iconloc :=  revel.BasePath + "/" + PhotoIcon + file
           if !oslib.CheckFileIsExist(qrfn) {
                if _, err := qrcode.CreateQrcode(app.Url, qrfn) ; err != nil{
                    	revel.ERROR.Printf("Failed to create QRCODE  : %s %s", k,qrfn,  err)
                      continue
                }
           }
       app.QrcodeLoc = "images/qrimg/" + file
       app.Icon = "images/icon/" + file
      s.AppList[k] = app
	}//slist

}

func (s  *PhotoMod) GetApp(Name string)  (* appmodel.AppModel, error) {
      if  v,ok := s.AppList[Name] ; ok {
          return &v, nil
      }
      return   nil , errors.New("no data")
}

func (s  *PhotoMod) GetAllApp( )  (* [] appmodel.AppModel, error) {
      if  s.AppList == nil  {
           return nil, errors.New("no data")
      }

      keys := make([]string, 0, len(s.AppList))
      for key, _ := range s.AppList {
          keys = append(keys, key)
      }
      sort.Strings(keys)

       var list []  appmodel.AppModel
       for _, v := range keys {
          list =  append(list, s.AppList[v])
      }
      return &list, nil
}


func (s  *PhotoMod) GetNext() *  appmodel.AppModel {
     return nil
}

func (s  *PhotoMod) GetAppById(Id uint) *  appmodel.AppModel {
       if  Id >=uint(len(s.AppList)) {
         return nil
       }
      return nil
}
