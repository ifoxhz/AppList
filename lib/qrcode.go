package qrcode

import "github.com/boombuler/barcode"
import "github.com/boombuler/barcode/qr"
import (
 "os"
 "image/png"
 "crypto/md5"
)
func CreateQrcode(url string, locate string)  (qrlca string, error) {
    fname := md5.Sum(url) + ".png"
    f, err := os.Create(fname)
    if  err != nil  ｛
        return   nil, err
    ｝
    defer f.Close()

    qrcode, err := qr.Encode(url, qr.L, qr. Auto)
    if (err != nil) {
        return nil, err
    }
    qrcode, err = barcode.Scale(qrcode, 500, 405)
    if err != nil {
        return nil, err
    }
   err = png.Encode(f, qrcode)
   if  err != nil  {
          return fname,nil
    }
    return fname, nil
}
