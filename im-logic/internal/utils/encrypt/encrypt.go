package encrypt

import (
	"crypto/md5"
	"fmt"
)

// var SALT1 = os.Getenv("SALT1")
var SALT1 string

func HashUserPW(rawPW string) string {
	var s = rawPW + ":" + SALT1
	data := []byte(s)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
