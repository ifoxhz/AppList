package model
import (
"sort"
 "errors"
 //"os"
 //"crypto/md5"
 "github.com/revel/revel"
 "github.com/applist/lib/oslib"
 "github.com/applist/lib/qrcode"
)

type SocialApp struct {
     Name string
     Url      string
     Icon    string
     Qrcodeloc  string
}

type SocialMod struct {
      AppList  map[string] SocialApp
}
const  SocialQrLocation = "public/social/images/qrimg"



func  (s  *SocialMod) Load()  map[string] SocialApp  {
  surl   := map[string] SocialApp  {
   `微信`: {"微信",`https://itunes.apple.com/cn/app/wei-xin/id414478124?mt=8&v0=WWW-GCCN-ITSTOP100-FREEAPPS&l=&ign-mpt=uo%3D4`,
                    "images/icon/wechat.png", ""},
   `QQ`:{"QQ",`https://itunes.apple.com/cn/app/qq/id444934666?mt=8`,
               "images/icon/qq.png", ""},
    "陌陌":{"陌陌",`https://itunes.apple.com/cn/app/mo-mo-shou-ji-bi-bei-she-jiao/id448165862?mt=8`,
                 "images/icon/momo.png", ""},
    "豆瓣":{"豆瓣",`https://itunes.apple.com/cn/app/dou-ban/id907002334?mt=8`,
                    "images/icon/douban.jpg",""},
    "微博":{"微博",`https://itunes.apple.com/cn/app/wei-bo/id350962117?mt=8`,
                      "images/icon/weibo.jpg",""},
    "人人":{"人人",`https://itunes.apple.com/cn/app/ren-ren-quan-min-zhi-bo-mei/id316709252?mt=8`,
                      "images/icon/renren.png",""},

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
