package keymethod

import (
	"crypto/rsa"
	"crypto/x509"
	"os"

	"github.com/Fuses-Garage/UnityGo/util"
)

func GetPubKey() *rsa.PublicKey {
	bytes, err := os.ReadFile("derrsapubkey.key")
	util.CheckErr(err)
	pubk, err := x509.ParsePKCS1PublicKey(bytes)
	util.CheckErr(err)
	return pubk
}

func getPrvKey() *rsa.PrivateKey {
	bytes, err := os.ReadFile("derrsaprvkey.key")
	util.CheckErr(err)
	prvk, err := x509.ParsePKCS1PrivateKey(bytes)
	util.CheckErr(err)
	return prvk
}
