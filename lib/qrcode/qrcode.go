package qrcode

import "github.com/boombuler/barcode"
import "github.com/boombuler/barcode/qr"
//import "github.com/ifoxhz/applist/lib/oslib"

import (
 "os"
 "image/png"
)
func CreateQrcode(url string, location string)  (string , error) {

    qrcode, err := qr.Encode(url, qr.L, qr. Auto)
    if (err != nil) {
        return "", err
    }
    qrcode, err = barcode.Scale(qrcode, 500, 405)
    if err != nil {
        return "" , err
    }

    f, err := os.Create(location)
    if  err != nil {
        return   "", err
    }

  defer f.Close()
   err = png.Encode(f, qrcode)
   if  err != nil  {
          //remove bad qr
          os.Remove(location)
          return "", err
    }
    return location, nil
}
