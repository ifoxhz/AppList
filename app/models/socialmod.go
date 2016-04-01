package socailorm
import (
"sort"
 "errors"
 //"os"
 //"crypto/md5"
 "github.com/revel/revel"
 "github.com/ifoxhz/applist/lib/oslib"
 "github.com/ifoxhz/applist/lib/qrcode"
)

type SocialApp struct {
     Id        string
     Name string
     Url      string
     Qrcodeloc  string
     Comments   string
}

type SocialMod struct {
      AppList  map[string] SocialApp
}
const  SocialQrLocation = "public/social/images/qrimg"



func  (s  *SocialMod) Load()  map[string] string  {
  surl   := map[string] string  {
   `微信`:`https://itunes.apple.com/cn/app/wei-xin/id414478124?mt=8&v0=WWW-GCCN-ITSTOP100-FREEAPPS&l=&ign-mpt=uo%3D4`,
   `QQ`:`www.qq.com`,
 }
 return surl
}

func  (s  *SocialMod) Init()   {
     if  s.AppList != nil {
       return
    }
    slist := s.Load()

    s.AppList = make(map[string] SocialApp)
    for k, v := range slist {
          app := SocialApp  {
                           Name:  k,
                           Url:  v,
                      }
           qrfn := SocialQrLocation + "/" + oslib.JoinQRcodeName(v)
           revel.INFO.Printf("%s",qrfn)
           if !oslib.CheckFileIsExist(qrfn) {
                if _, err := qrcode.CreateQrcode(v, revel.BasePath + "/" + qrfn) ; err != nil{
                    	revel.ERROR.Printf("Failed to create QRCODE  : %s %s", k,qrfn,  err)
                      continue
                }
                app.Qrcodeloc = "images/qrimg/"  + oslib.JoinQRcodeName(v)
           }

        s.AppList[k] = app
	}//slist

}

func (s  *SocialMod) GetApp(Name string)  (*SocialApp, error) {
      if  v,ok := s.AppList[Name] ; ok {
          return &v, nil
      }
      return   nil , errors.New("no data")
}

func (s  *SocialMod) GetAllApp( )  (* []SocialApp, error) {
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
         	revel.INFO.Printf("app : %s", v)
          list =  append(list, s.AppList[v])
      }
      return &list, nil
}


func (s  *SocialMod) GetNext() *SocialApp {
     return nil
}

func (s  *SocialMod) GetAppById(Id uint) * SocialApp {
       if  Id >=uint(len(s.AppList)) {
         return nil
       }
      return nil
}
