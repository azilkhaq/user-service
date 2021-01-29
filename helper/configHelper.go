package helper

import (
	"crypto/rand"
	"fmt"
	"log"
)

func GENERATEUUID() string {
	var err error
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	result := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return result
}
