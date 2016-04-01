package oslib
import (
 //"errors"
 "os"
  "crypto/md5"
  "fmt"
 	//"github.com/revel/revel"
)

func CheckFileIsExist(filename string) (bool) {
      var exist = true
      if _, err := os.Stat(filename); os.IsNotExist(err) {
          exist = false
      }
      return exist
}

func JoinQRcodeName(url string) string{
      md := fmt.Sprintf("%x" , md5.Sum( []byte(url) ))
      return  md  + ".png"
}
