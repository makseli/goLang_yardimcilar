package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
)

func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}

func main() {

	su := "deneme0deger"

if false {
		fmt.Println("'tekbiir' değeri şifreli hali --> ", encryptPassword("tekbiir"))
}else{
	
	for i := 0; i < 10; i++ {

		su = "deger" + strconv.Itoa(i) + "sonraki"
		//fmt.Println(su)
		fmt.Println( su , " -",i,"-> " , encryptPassword(su))
	}
}


}

