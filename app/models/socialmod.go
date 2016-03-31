package socailorm
import (
"sort"
 "errors"
 	"github.com/revel/revel"
)

type SocialApp struct {
     Id        string
     Name string
     Url      string
     qrcode  string
     Comments   string
}

type SocialMod struct {
      AppList  map[string] SocialApp
}

func  (s  *SocialMod) Load()  map[string] string  {
  surl   := map[string] string  {
   `微信`:`https://itunes.apple.com/cn/app/wei-xin/id414478124?mt=8&v0=WWW-GCCN-ITSTOP100-FREEAPPS&l=&ign-mpt=uo%3D4`,
   `1`:`www.qq.com`,
   `2`:`www.qq.com`,
   `3`:`www.qq.com`,
   `4`:`www.qq.com`,
   `5`:`www.qq.com`,
   `6`:`www.qq.com`,
   `7`:`www.qq.com`,
   `8`:`www.qq.com`,
   `gggg`:`www.ggoo.com`,
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
          revel.WARN.Printf("items %s", v )
        s.AppList[k] = app
	}

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
         	revel.WARN.Printf("app : %s", v)
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
