package src

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
)


func ReadTxtFromFile(path string) string {
	bytes,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("total bytes read：",len(bytes))
	fmt.Println("string read:",string(bytes))
	return string(bytes)
}

func GetNonce() string{
	key := make([]byte, 24)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	nonce := base64.StdEncoding.EncodeToString(key)
	return nonce
}
