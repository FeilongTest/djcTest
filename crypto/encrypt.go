package crypto

import (
	"fmt"
	"strings"
)

func GetEncrypt(content string) (result string, err error) {
	aesKey := []byte("se35d32s63r7m23m")
	// AES加密
	aesEncrypted, err := aesEncrypt([]byte(content), aesKey)
	if err != nil {
		return
	}
	// RSA加密
	rsaEncrypted, err := rsaEncrypted(aesEncrypted)
	if err != nil {
		return
	}
	result = byte2Hex(rsaEncrypted)
	return
}

func byte2Hex(bArr []byte) string {
	var hexStr strings.Builder
	for _, b := range bArr {
		hexStr.WriteString(fmt.Sprintf("%02x", b))
	}
	return hexStr.String()
}
