package djc

import (
	"bytes"
	crypto "crypto/rand"
	"math/big"
)

func CreateRandomString(len int) string {
	var container string
	var str = "0123456789"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := crypto.Int(crypto.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
