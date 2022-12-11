package keymethod

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"os"

	"github.com/Fuses-Garage/UnityGo/util"
)

func MakeKey() {
	rsaprvkey, err := rsa.GenerateKey(rand.Reader, 2048)
	util.CheckErr(err)
	derrsaprvkey := x509.MarshalPKCS1PrivateKey(rsaprvkey)
	file, err := os.Create("derrsaprvkey.key")
	util.CheckErr(err)
	_, err = file.Write(derrsaprvkey)
	util.CheckErr(err)
	err = file.Close()
	util.CheckErr(err)
	var rsapubkey crypto.PublicKey
	rsapubkey = rsaprvkey.Public()
	if rsapubkeypoint, ok := rsapubkey.(*rsa.PublicKey); ok {
		derrsapubkey := x509.MarshalPKCS1PublicKey(rsapubkeypoint)
		file, err = os.Create("derrsapubkey.key")
		util.CheckErr(err)
		_, err = file.Write(derrsapubkey)
		util.CheckErr(err)
		err = file.Close()
		util.CheckErr(err)
	}
}
