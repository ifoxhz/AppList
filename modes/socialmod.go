package socailorm
import (
//	"golang.org/x/crypto/bcrypt"
	//"fmt"
	"github.com/ifoxhz/applist/app/models"
	"github.com/ifoxhz/applist/app/routes"
	"github.com/ifoxhz/applist/app/lib"
	"strings"
  "os"
)

type SocialApp struct {
     AppId        string
     AppName string
     AppUrl      string
     AppComments   string
}

type SocialMod struct {
      AppList Map[string] SocialApp
}

func (s  *SocialMod) Init(Name string)  SocialApp {
  SocailUrl := map[string] string  {
          "微信"："https://itunes.apple.com/cn/app/wei-xin/id414478124?mt=8&v0=WWW-GCCN-ITSTOP100-FREEAPPS&l=&ign-mpt=uo%3D4",
  }

     if  Applist == nil {
              for k, v := range SocailUrl {
                    SocailApp := {
                         name:    k,
                         AppUrl:  v,
              }
               AppList.append(k,v)
							


         }
    }
}

func (s  *SocialMod) GetApp(Name string)  SocialApp {
      if AppList[Name] != nil {
          return AppList[Name]
      }
      return nil
}

func (s  *SocialMod) GetAllApp( )  []SocialApp {
      if AppList[Name] == nil {
           return nil
      }
       var list []SocialApp
       for _, v = range Applist {
           list.append(v)
      }
      return list
}


func (s  *SocialMod) GetNext() SocialApp {
     return nil
}

func (s  *SocialMod) GetApp(Id uint) SocialApp {
       if  Id >= len(Applist) {
         return nil
       }
      return nil
}

type SocialOrm interface{
          GetApp(Name string) SocialMod
          GetNextApp() SocialMod
          GetApp(Id uint) SocailMod
}
