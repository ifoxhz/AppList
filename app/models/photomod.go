package socailorm
import (
"sort"
 "errors"
 //"os"
 //"crypto/md5"
 "github.com/revel/revel"
 "github.com/ifoxhz/applist/lib/oslib"
 "github.com/ifoxhz/applist/lib/qrcode"
 "github.com/ifoxhz/applsit/lib/appmodel"
)

var RedisOption =  redis.Options {
    Addr: "127.0.0.1:6379",
    Password:"",
    DB: 0,
}
var Rdb  * redis.Client = nil

type PhotoMod struct {
      AppList  map[string]  AppModel
}
const  SocialQrLocation = "public/photo/images/qrimg"



func  (s  *PhotoMod) Load()  (map[string] AppModel, err) {
  if (Rdb == nil){
      Rdb =  redis.NewClient(&RedisOption)
 }

 if err := Rdb.Incr("IOTHILL.PAGEVIEW").Err(); err != nil {
    revel.INFO.Printf("failed to conneted redis ", err)
 }
  return surl := AppModel{}
}

func  (s  *PhotoMod) Init()   {
     if  s.AppList != nil {
       return
    }
    slist := s.Load()

    s.AppList = make(map[string] SocialApp)
    for k, v := range slist {
          app := v
          qrfn := SocialQrLocation + "/" + oslib.JoinQRcodeName(v.Url)
           if !oslib.CheckFileIsExist(qrfn) {
                if _, err := qrcode.CreateQrcode(v.Url, revel.BasePath + "/" + qrfn) ; err != nil{
                    	revel.ERROR.Printf("Failed to create QRCODE  : %s %s", k,qrfn,  err)
                      continue
                }
                app.Qrcodeloc = "images/qrimg/"  + oslib.JoinQRcodeName(v.Url)
           }

        s.AppList[k] = app
	}//slist

}

func (s  *PhotoMod) GetApp(Name string)  (*SocialApp, error) {
      if  v,ok := s.AppList[Name] ; ok {
          return &v, nil
      }
      return   nil , errors.New("no data")
}

func (s  *PhotoMod) GetAllApp( )  (* []SocialApp, error) {
      if  s.AppList == nil  {
           return nil, errors.New("no data")
      }

      keys := make([]string, 0, len(s.AppList))
      for key, _ := range s.AppList {
          keys = append(keys, key)
      }
      sort.Strings(keys)

       var list []SocialApp
       for _, v := range keys {
          list =  append(list, s.AppList[v])
      }
      return &list, nil
}


func (s  *PhotoMod) GetNext() *SocialApp {
     return nil
}

func (s  *PhotoMod) GetAppById(Id uint) * SocialApp {
       if  Id >=uint(len(s.AppList)) {
         return nil
       }
      return nil
}
