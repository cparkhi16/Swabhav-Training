package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashForPassword(password string) string {
	p := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func ComparePasswords(hashedPwd, pwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
