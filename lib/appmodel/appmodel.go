package appmodel
import (
  //"strings"
)

type Category int
const  (
    social Category = iota +100
    photo
)

var ModeKeyMap = map[Category] string {
           social:  "APP:SOCIAL",
           photo:  "APP:PHOTO",
}
var  IconLoaction = map[Category] string {
           social:  "public/social/images/icon/",
           photo:  "public/photo/images/icon/",
}

var  QrimageLoaction = map[Category] string {
           social:  "public/social/images/qrimg/",
           photo:  "public/photo/images/qrimg/",
}

type AppModel struct {
  Id         string
	Name   string
	Url        string
	Icon      string
	QrcodeLoc string
  Tag           string
  Category   Category
}

func  GetIconLocation(cat Category) string{
        return  IconLoaction[cat]
}
