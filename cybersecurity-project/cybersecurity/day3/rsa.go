package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"crypto/sha512"
	"fmt"
)

func main() {
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println("private key ", privatekey,reflect.TypeOf(privatekey))
	fmt.Println("private key2-",privatekey.D.Bytes())
	publickey := privatekey.PublicKey
	fmt.Println("Public key-",reflect.TypeOf(publickey))
	msg := "s"
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publickey, []byte(msg), nil)
	fmt.Println("encrypted text ", string(cipher))
	var ss =[]byte{73,63,243,130,225,135,228,201,112,33,174,227,32,137,103,219,16,118,26,164,179,58,103,255,54,238,217,218,202,110,52,15,50,116,231,223,26,249,83,57,82,99,122,113,235,189,107,252,164,75,106,81,93,233,34,146,111,204,212,50,214,145,19,75,178,224,7,32,174,61,55,74,98,49,75,181,199,183,103,246,12,81,83,162,66,176,136,234,85,205,52,64,73,167,216,186,130,194,62,184,179,37,67,176,140,200,55,233,222,5,254,168,112,207,141,239,67,123,108,64,47,125,36,87,48,221,169,151,136,24,71,141,186,2,155,136,163,157,209,147,168,28,70,33,227,56,248,126,253,205,97,172,174,88,55,122,65,137,5,12,142,68,216,180,214,123,240,171,29,17,38,76,160,186,58,181,221,36,111,138,217,91,226,64,155,245,233,245,232,31,110,7,154,66,50,41,38,166,153,77,113,61,97,100,122,10,62,40,160,133,203,134,155,195,35,45,250,19,7,224,232,126,0,236,6,43,216,241,139,176,255,177,0,152,10,3,2,140,253,177,254,44,201,225,210,141,227,182,234,232,94,74,74,214,237,89}
	plainText, err := privatekey.Decrypt(nil,ss , &rsa.OAEPOptions{Hash: crypto.SHA512})
	fmt.Println("Plain text ", string(plainText),err)

}
