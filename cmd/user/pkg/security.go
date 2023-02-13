package pkg

import (
	"crypto/md5"
	"encoding/hex"
)


// HashPassword 密码加密
func HashPassword(pwd string)(string,error){
	h := md5.New()
	_,err := h.Write([]byte(pwd))
	if err != nil {
		return "",err
	}

	return hex.EncodeToString(h.Sum(nil)),nil
}