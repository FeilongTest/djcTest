package crypto

import (
	"github.com/farmerx/gorsa"
)

func Init() {
	const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoYepGcQuSHKCJK3HMQCW
iPVKJ2BGDILSFNWjD5k6Ass77+ZHmhy58zG96eD9jMwjnZwItA8jMGfGBlWTkPp6
yaKWJq0Dxqik99xBaLKnZN+Sxcfk3L8W7Mk+HoZZjurqdIjr73jDQbEzDeS3IzZG
XBm6AkuopduhMHfGAOaENJ3LjxcTN/KKNBfiIzg4CI/TX2RPTawlivlBsKXLKN8z
zziA5PQZfomld+jZX+f7Nn7ki08kqqCVINLVNJnj9JMXQI/2E6s6OVRQP1YE6wF/
qcLB0aEeDZWZyTA6h5a3h+CpCxOtmmTTX0wqH38XBwTMRHFxn9IhiIUTiTSJ053S
EQIDAQAB
-----END PUBLIC KEY-----`
	_ = gorsa.RSA.SetPublicKey(publicKey)
}

func rsaEncrypted(content []byte) (encrypt []byte, err error) {
	encrypt, err = gorsa.RSA.PubKeyENCTYPT(content)
	return
}
